package resources

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"

	integreatlyv1alpha1 "github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"

	productsConfig "github.com/integr8ly/integreatly-operator/pkg/config"
	"github.com/integr8ly/integreatly-operator/pkg/resources/marketplace"
	oauthv1 "github.com/openshift/api/oauth/v1"

	operatorsv1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
	"github.com/operator-framework/operator-lifecycle-manager/pkg/lib/ownerutil"

	corev1 "k8s.io/api/core/v1"
	k8serr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	OwnerLabelKey = integreatlyv1alpha1.SchemeGroupVersion.Group + "/installation-uid"
)

// This is the base reconciler that all the other reconcilers extend. It handles things like namespace creation, subscription creation etc

type Reconciler struct {
	mpm marketplace.MarketplaceInterface
}

func NewReconciler(mpm marketplace.MarketplaceInterface) *Reconciler {
	return &Reconciler{
		mpm: mpm,
	}
}

func (r *Reconciler) ReconcileOauthClient(ctx context.Context, inst *integreatlyv1alpha1.Installation, client *oauthv1.OAuthClient, apiClient k8sclient.Client) (integreatlyv1alpha1.StatusPhase, error) {
	if err := apiClient.Get(ctx, k8sclient.ObjectKey{Name: client.Name}, client); err != nil {
		if k8serr.IsNotFound(err) {
			PrepareObject(client, inst)
			if err := apiClient.Create(ctx, client); err != nil {
				return integreatlyv1alpha1.PhaseFailed, fmt.Errorf("failed to create oauth client: %s. %w", client.Name, err)
			}
			return integreatlyv1alpha1.PhaseCompleted, nil
		}
		return integreatlyv1alpha1.PhaseFailed, fmt.Errorf("failed to get oauth client: %s. %w", client.Name, err)
	}
	PrepareObject(client, inst)
	if err := apiClient.Update(ctx, client); err != nil {
		return integreatlyv1alpha1.PhaseFailed, fmt.Errorf("failed to update oauth client: %s. %w", client.Name, err)
	}
	return integreatlyv1alpha1.PhaseCompleted, nil
}

// GetNS gets the specified corev1.Namespace from the k8s API server
func GetNS(ctx context.Context, namespace string, client k8sclient.Client) (*corev1.Namespace, error) {
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}
	err := client.Get(ctx, k8sclient.ObjectKey{Name: ns.Name}, ns)
	if err == nil {
		// workaround for https://github.com/kubernetes/client-go/issues/541
		ns.TypeMeta = metav1.TypeMeta{Kind: "Namespace", APIVersion: corev1.SchemeGroupVersion.Version}
	}
	return ns, err
}

func (r *Reconciler) ReconcileNamespace(ctx context.Context, namespace string, inst *integreatlyv1alpha1.Installation, client k8sclient.Client) (integreatlyv1alpha1.StatusPhase, error) {
	ns, err := GetNS(ctx, namespace, client)
	if err != nil {
		if !k8serr.IsNotFound(err) {
			return integreatlyv1alpha1.PhaseFailed, fmt.Errorf("could not retrieve namespace: %s. %w", ns.Name, err)
		}
		PrepareObject(ns, inst)
		if err = client.Create(ctx, ns); err != nil {
			return integreatlyv1alpha1.PhaseFailed, fmt.Errorf("could not create namespace: %s. %w", ns.Name, err)
		}
		return integreatlyv1alpha1.PhaseCompleted, nil
	}
	// ns exists so check it is our namespace
	if !IsOwnedBy(ns, inst) && ns.Status.Phase != corev1.NamespaceTerminating {
		return integreatlyv1alpha1.PhaseFailed, fmt.Errorf("existing namespace found with name %v but it is not owned by the integreatly installation and it isn't being deleted", ns.Name)
	}
	if ns.Status.Phase == corev1.NamespaceTerminating {
		logrus.Debugf("namespace %s is terminating, maintaining phase to try again on next reconcile", namespace)
		return integreatlyv1alpha1.PhaseInProgress, nil
	}
	PrepareObject(ns, inst)
	if err := client.Update(ctx, ns); err != nil {
		return integreatlyv1alpha1.PhaseFailed, fmt.Errorf("failed to update the ns definition: %w", err)
	}
	if ns.Status.Phase != corev1.NamespaceActive {
		return integreatlyv1alpha1.PhaseInProgress, nil
	}
	return integreatlyv1alpha1.PhaseCompleted, nil
}

type finalizerFunc func() (integreatlyv1alpha1.StatusPhase, error)

func (r *Reconciler) ReconcileFinalizer(ctx context.Context, client k8sclient.Client, inst *integreatlyv1alpha1.Installation, productName string, finalFunc finalizerFunc) (integreatlyv1alpha1.StatusPhase, error) {
	finalizer := "finalizer." + productName + ".integreatly.org"
	// Add finalizer if not there
	err := AddFinalizer(ctx, inst, client, finalizer)
	if err != nil {
		logrus.Error(fmt.Sprintf("Error adding finalizer %s to installation", finalizer), err)
		return integreatlyv1alpha1.PhaseFailed, err
	}

	// Run finalization logic. If it fails, don't remove the finalizer
	// so that we can retry during the next reconciliation
	if inst.GetDeletionTimestamp() != nil {
		if contains(inst.GetFinalizers(), finalizer) {
			phase, err := finalFunc()
			if err != nil || phase != integreatlyv1alpha1.PhaseCompleted {
				return phase, err
			}

			// Remove the finalizer to allow for deletion of the installation cr
			logrus.Infof("Removing finalizer: %s", finalizer)
			err = RemoveProductFinalizer(ctx, inst, client, productName)
			if err != nil {
				return integreatlyv1alpha1.PhaseFailed, err
			}
		}
		// Don't continue reconciling the product
		return integreatlyv1alpha1.PhaseNone, nil
	}
	return integreatlyv1alpha1.PhaseCompleted, nil
}

func (r *Reconciler) ReconcilePullSecret(ctx context.Context, namespace, secretName string, inst *integreatlyv1alpha1.Installation, client k8sclient.Client) (integreatlyv1alpha1.StatusPhase, error) {
	pullSecretName := DefaultOriginPullSecretName
	if secretName != "" {
		pullSecretName = secretName
	}

	err := CopyDefaultPullSecretToNameSpace(namespace, pullSecretName, inst, client, ctx)
	if err != nil {
		return integreatlyv1alpha1.PhaseFailed, fmt.Errorf("error creating/updating secret '%s' in namespace: '%s': %w", pullSecretName, namespace, err)
	}

	return integreatlyv1alpha1.PhaseCompleted, nil
}

func (r *Reconciler) ReconcileSubscription(ctx context.Context, owner ownerutil.Owner, t marketplace.Target, targetNS string, client k8sclient.Client) (integreatlyv1alpha1.StatusPhase, error) {
	logrus.Infof("reconciling subscription %s from channel %s in namespace: %s", t.Pkg, "integreatly", t.Namespace)
	err := r.mpm.InstallOperator(ctx, client, owner, t, []string{targetNS}, operatorsv1alpha1.ApprovalManual)

	if err != nil && !k8serr.IsAlreadyExists(err) {
		return integreatlyv1alpha1.PhaseFailed, fmt.Errorf("could not create subscription in namespace: %s: %w", t.Namespace, err)
	}
	ips, _, err := r.mpm.GetSubscriptionInstallPlans(ctx, client, t.Pkg, t.Namespace)
	if err != nil {
		// this could be the install plan or subscription so need to check if sub nil or not TODO refactor
		if k8serr.IsNotFound(err) || k8serr.IsNotFound(errors.Unwrap(err)) {
			return integreatlyv1alpha1.PhaseAwaitingOperator, nil
		}
		return integreatlyv1alpha1.PhaseFailed, fmt.Errorf("could not retrieve installplan and subscription in namespace: %s: %w", t.Namespace, err)
	}

	if len(ips.Items) == 0 {
		return integreatlyv1alpha1.PhaseInProgress, nil
	}

	for _, ip := range ips.Items {
		err = upgradeApproval(ctx, client, &ip)
		if err != nil {
			return integreatlyv1alpha1.PhaseFailed, fmt.Errorf("error approving installplan for %v: %w", t.Pkg, err)
		}

		//if it's approved but not complete, then it's in progress
		if ip.Status.Phase != operatorsv1alpha1.InstallPlanPhaseComplete && ip.Spec.Approved {
			logrus.Infof("%s install plan is not complete yet ", t.Pkg)
			return integreatlyv1alpha1.PhaseInProgress, nil
			//if it's not approved by now, then it will not be approved by this version of the integreatly-operator
		} else if !ip.Spec.Approved {
			logrus.Infof("%s has an upgrade installplan above the maximum allowed version", t.Pkg)
		}
	}

	return integreatlyv1alpha1.PhaseCompleted, nil
}

func PrepareObject(ns metav1.Object, install *integreatlyv1alpha1.Installation) {
	labels := ns.GetLabels()
	if labels == nil {
		labels = map[string]string{}
	}
	labels["integreatly"] = "true"
	labels["monitoring-key"] = "middleware"
	labels[OwnerLabelKey] = string(install.GetUID())
	monitoringConfig := productsConfig.NewMonitoring(productsConfig.ProductConfig{})
	labels[monitoringConfig.GetLabelSelectorKey()] = monitoringConfig.GetLabelSelector()
	ns.SetLabels(labels)
}

func IsOwnedBy(o metav1.Object, owner *integreatlyv1alpha1.Installation) bool {
	// TODO change logic to check for our finalizer?
	for k, v := range o.GetLabels() {
		if k == OwnerLabelKey && v == string(owner.UID) {
			return true
		}
	}
	return false
}
