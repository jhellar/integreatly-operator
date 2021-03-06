package config

import (
	crov1alpha1 "github.com/integr8ly/cloud-resource-operator/pkg/apis/integreatly/v1alpha1"
	integreatlyv1alpha1 "github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type CloudResources struct {
	Config ProductConfig
}

func NewCloudResources(config ProductConfig) *CloudResources {
	return &CloudResources{Config: config}
}

func (c *CloudResources) GetWatchableCRDs() []runtime.Object {
	return []runtime.Object{
		&crov1alpha1.Postgres{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Postgres",
				APIVersion: crov1alpha1.SchemeGroupVersion.String(),
			},
		},
		&crov1alpha1.Redis{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Redis",
				APIVersion: crov1alpha1.SchemeGroupVersion.String(),
			},
		},
		&crov1alpha1.BlobStorage{
			TypeMeta: metav1.TypeMeta{
				Kind:       "BlobStorage",
				APIVersion: crov1alpha1.SchemeGroupVersion.String(),
			},
		},
		&crov1alpha1.SMTPCredentialSet{
			TypeMeta: metav1.TypeMeta{
				Kind:       "SMTPCredentialSet",
				APIVersion: crov1alpha1.SchemeGroupVersion.String(),
			},
		},
	}
}

func (c *CloudResources) GetHost() string {
	return c.Config["HOST"]
}

func (c *CloudResources) SetHost(newHost string) {
	c.Config["HOST"] = newHost
}

func (c *CloudResources) GetNamespace() string {
	return c.Config["NAMESPACE"]
}

func (c *CloudResources) SetNamespace(newNamespace string) {
	c.Config["NAMESPACE"] = newNamespace
}

func (c *CloudResources) Read() ProductConfig {
	return c.Config
}

func (c *CloudResources) GetProductName() integreatlyv1alpha1.ProductName {
	return integreatlyv1alpha1.ProductCloudResources
}

func (c *CloudResources) GetProductVersion() integreatlyv1alpha1.ProductVersion {
	return integreatlyv1alpha1.VersionCloudResources
}

func (c *CloudResources) GetOperatorVersion() integreatlyv1alpha1.OperatorVersion {
	return integreatlyv1alpha1.OperatorVersionCloudResources
}
