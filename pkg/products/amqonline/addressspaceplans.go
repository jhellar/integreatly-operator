package amqonline

import (
	"github.com/integr8ly/integreatly-operator/pkg/apis/enmasse/v1beta2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetDefaultAddressSpacePlans(ns string) []*v1beta2.AddressSpacePlan {
	return []*v1beta2.AddressSpacePlan{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "brokered-single-broker",
				Namespace: ns,
			},
			Spec: v1beta2.AddressSpacePlanSpec{
				DisplayName:      "Single Broker",
				DisplayOrder:     0,
				InfraConfigRef:   "default",
				ShortDescription: "Single Broker instance",
				LongDescription:  "Single Broker plan where you can create an infinite number of queues until the system falls over.",
				AddressSpaceType: "brokered",
				ResourceLimits: v1beta2.AddressSpacePlanResourceLimits{
					Broker: 1.9,
				},
				AddressPlans: []string{
					"brokered-queue",
					"brokered-topic",
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "standard-medium",
				Namespace: ns,
			},
			Spec: v1beta2.AddressSpacePlanSpec{
				DisplayName:      "Medium",
				DisplayOrder:     0,
				InfraConfigRef:   "default",
				ShortDescription: "Messaging infrastructure based on Apache Qpid Dispatch Router and Apache ActiveMQ Artemis.",
				LongDescription:  "Messaging infrastructure based on Apache Qpid Dispatch Router and Apache ActiveMQ Artemis. This plan allows up to 3 routers and 3 broker in total, and is suitable for applications using small address plans and few addresses.",
				AddressSpaceType: "standard",
				ResourceLimits: v1beta2.AddressSpacePlanResourceLimits{
					Broker:    3.0,
					Router:    3.0,
					Aggregate: 6.0,
				},
				AddressPlans: []string{
					"standard-small-anycast",
					"standard-medium-anycast",
					"standard-large-anycast",
					"standard-small-multicast",
					"standard-medium-multicast",
					"standard-large-multicast",
					"standard-small-queue",
					"standard-medium-queue",
					"standard-large-queue",
					"standard-xlarge-queue",
					"standard-small-topic",
					"standard-medium-topic",
					"standard-large-topic",
					"standard-xlarge-topic",
					"standard-small-subscription",
					"standard-medium-subscription",
					"standard-large-subscription",
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "standard-small",
				Namespace: ns,
			},
			Spec: v1beta2.AddressSpacePlanSpec{
				DisplayName:      "Small",
				DisplayOrder:     0,
				InfraConfigRef:   "default-minimal",
				ShortDescription: "Messaging infrastructure based on Apache Qpid Dispatch Router and Apache ActiveMQ Artemis.",
				LongDescription:  "Messaging infrastructure based on Apache Qpid Dispatch Router and Apache ActiveMQ Artemis. This plan allows up to 1 router and 1 broker in total, and is suitable for small applications using small address plans and few addresses.",
				AddressSpaceType: "standard",
				ResourceLimits: v1beta2.AddressSpacePlanResourceLimits{
					Broker:    1.0,
					Router:    1.0,
					Aggregate: 2.0,
				},
				AddressPlans: []string{
					"standard-small-anycast",
					"standard-medium-anycast",
					"standard-large-anycast",
					"standard-small-multicast",
					"standard-medium-multicast",
					"standard-large-multicast",
					"standard-small-queue",
					"standard-medium-queue",
					"standard-large-queue",
					"standard-small-topic",
					"standard-medium-topic",
					"standard-large-topic",
					"standard-small-subscription",
					"standard-medium-subscription",
					"standard-large-subscription",
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "standard-unlimited-with-mqtt",
				Namespace: ns,
			},
			Spec: v1beta2.AddressSpacePlanSpec{
				DisplayName:      "Unlimited With MQTT",
				DisplayOrder:     0,
				InfraConfigRef:   "default-with-mqtt",
				ShortDescription: "Messaging infrastructure based on Apache Qpid Dispatch Router and Apache ActiveMQ Artemis and MQTT support.",
				LongDescription:  "Messaging infrastructure based on Apache Qpid Dispatch Router and Apache ActiveMQ Artemis and MQTT support. This plan allows an unlimited number of routers (minimum 2) and brokers, and is suitable for applications where you do not want to impose any restrictions.",
				AddressSpaceType: "standard",
				ResourceLimits: v1beta2.AddressSpacePlanResourceLimits{
					Broker:    10000.0,
					Router:    10000.0,
					Aggregate: 10000.0,
				},
				AddressPlans: []string{
					"standard-small-anycast",
					"standard-medium-anycast",
					"standard-large-anycast",
					"standard-small-multicast",
					"standard-medium-multicast",
					"standard-large-multicast",
					"standard-small-queue",
					"standard-medium-queue",
					"standard-large-queue",
					"standard-xlarge-queue",
					"standard-small-topic",
					"standard-medium-topic",
					"standard-large-topic",
					"standard-xlarge-topic",
					"standard-small-subscription",
					"standard-medium-subscription",
					"standard-large-subscription",
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "standard-unlimited",
				Namespace: ns,
			},
			Spec: v1beta2.AddressSpacePlanSpec{
				DisplayName:      "Unlimited",
				DisplayOrder:     0,
				InfraConfigRef:   "default",
				ShortDescription: "Messaging infrastructure based on Apache Qpid Dispatch Router and Apache ActiveMQ Artemis and MQTT support.",
				LongDescription:  "Messaging infrastructure based on Apache Qpid Dispatch Router and Apache ActiveMQ Artemis. This plan allows an unlimited number of routers and brokers, and is suitable for applications where you do not want to impose any restrictions.",
				AddressSpaceType: "standard",
				ResourceLimits: v1beta2.AddressSpacePlanResourceLimits{
					Broker:    10000.0,
					Router:    10000.0,
					Aggregate: 10000.0,
				},
				AddressPlans: []string{
					"standard-small-anycast",
					"standard-medium-anycast",
					"standard-large-anycast",
					"standard-small-multicast",
					"standard-medium-multicast",
					"standard-large-multicast",
					"standard-small-queue",
					"standard-medium-queue",
					"standard-large-queue",
					"standard-xlarge-queue",
					"standard-small-topic",
					"standard-medium-topic",
					"standard-large-topic",
					"standard-xlarge-topic",
					"standard-small-subscription",
					"standard-medium-subscription",
					"standard-large-subscription",
				},
			},
		},
	}
}
