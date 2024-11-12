// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=loadbalancersinboundnatrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={loadbalancersinboundnatrules/status,loadbalancersinboundnatrules/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20240301.LoadBalancersInboundNatRule
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2024-03-01/loadBalancer.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}
type LoadBalancersInboundNatRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancersInboundNatRule_Spec   `json:"spec,omitempty"`
	Status            LoadBalancersInboundNatRule_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &LoadBalancersInboundNatRule{}

// GetConditions returns the conditions of the resource
func (rule *LoadBalancersInboundNatRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *LoadBalancersInboundNatRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ configmaps.Exporter = &LoadBalancersInboundNatRule{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (rule *LoadBalancersInboundNatRule) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if rule.Spec.OperatorSpec == nil {
		return nil
	}
	return rule.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &LoadBalancersInboundNatRule{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (rule *LoadBalancersInboundNatRule) SecretDestinationExpressions() []*core.DestinationExpression {
	if rule.Spec.OperatorSpec == nil {
		return nil
	}
	return rule.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &LoadBalancersInboundNatRule{}

// AzureName returns the Azure name of the resource
func (rule *LoadBalancersInboundNatRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-03-01"
func (rule LoadBalancersInboundNatRule) GetAPIVersion() string {
	return "2024-03-01"
}

// GetResourceScope returns the scope of the resource
func (rule *LoadBalancersInboundNatRule) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (rule *LoadBalancersInboundNatRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *LoadBalancersInboundNatRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (rule *LoadBalancersInboundNatRule) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/loadBalancers/inboundNatRules"
func (rule *LoadBalancersInboundNatRule) GetType() string {
	return "Microsoft.Network/loadBalancers/inboundNatRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *LoadBalancersInboundNatRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &LoadBalancersInboundNatRule_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (rule *LoadBalancersInboundNatRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return rule.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (rule *LoadBalancersInboundNatRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*LoadBalancersInboundNatRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st LoadBalancersInboundNatRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// Hub marks that this LoadBalancersInboundNatRule is the hub type for conversion
func (rule *LoadBalancersInboundNatRule) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *LoadBalancersInboundNatRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion,
		Kind:    "LoadBalancersInboundNatRule",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20240301.LoadBalancersInboundNatRule
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2024-03-01/loadBalancer.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}
type LoadBalancersInboundNatRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancersInboundNatRule `json:"items"`
}

// Storage version of v1api20240301.LoadBalancersInboundNatRule_Spec
type LoadBalancersInboundNatRule_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName               string                                   `json:"azureName,omitempty"`
	BackendAddressPool      *SubResource                             `json:"backendAddressPool,omitempty"`
	BackendPort             *int                                     `json:"backendPort,omitempty"`
	EnableFloatingIP        *bool                                    `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                                    `json:"enableTcpReset,omitempty"`
	FrontendIPConfiguration *SubResource                             `json:"frontendIPConfiguration,omitempty"`
	FrontendPort            *int                                     `json:"frontendPort,omitempty"`
	FrontendPortRangeEnd    *int                                     `json:"frontendPortRangeEnd,omitempty"`
	FrontendPortRangeStart  *int                                     `json:"frontendPortRangeStart,omitempty"`
	IdleTimeoutInMinutes    *int                                     `json:"idleTimeoutInMinutes,omitempty"`
	OperatorSpec            *LoadBalancersInboundNatRuleOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion         string                                   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/LoadBalancer resource
	Owner       *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"LoadBalancer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Protocol    *string                            `json:"protocol,omitempty"`
}

var _ genruntime.ConvertibleSpec = &LoadBalancersInboundNatRule_Spec{}

// ConvertSpecFrom populates our LoadBalancersInboundNatRule_Spec from the provided source
func (rule *LoadBalancersInboundNatRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == rule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(rule)
}

// ConvertSpecTo populates the provided destination from our LoadBalancersInboundNatRule_Spec
func (rule *LoadBalancersInboundNatRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == rule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(rule)
}

// Storage version of v1api20240301.LoadBalancersInboundNatRule_STATUS
type LoadBalancersInboundNatRule_STATUS struct {
	BackendAddressPool      *SubResource_STATUS                                                                      `json:"backendAddressPool,omitempty"`
	BackendIPConfiguration  *NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded `json:"backendIPConfiguration,omitempty"`
	BackendPort             *int                                                                                     `json:"backendPort,omitempty"`
	Conditions              []conditions.Condition                                                                   `json:"conditions,omitempty"`
	EnableFloatingIP        *bool                                                                                    `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                                                                                    `json:"enableTcpReset,omitempty"`
	Etag                    *string                                                                                  `json:"etag,omitempty"`
	FrontendIPConfiguration *SubResource_STATUS                                                                      `json:"frontendIPConfiguration,omitempty"`
	FrontendPort            *int                                                                                     `json:"frontendPort,omitempty"`
	FrontendPortRangeEnd    *int                                                                                     `json:"frontendPortRangeEnd,omitempty"`
	FrontendPortRangeStart  *int                                                                                     `json:"frontendPortRangeStart,omitempty"`
	Id                      *string                                                                                  `json:"id,omitempty"`
	IdleTimeoutInMinutes    *int                                                                                     `json:"idleTimeoutInMinutes,omitempty"`
	Name                    *string                                                                                  `json:"name,omitempty"`
	PropertyBag             genruntime.PropertyBag                                                                   `json:"$propertyBag,omitempty"`
	Protocol                *string                                                                                  `json:"protocol,omitempty"`
	ProvisioningState       *string                                                                                  `json:"provisioningState,omitempty"`
	Type                    *string                                                                                  `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &LoadBalancersInboundNatRule_STATUS{}

// ConvertStatusFrom populates our LoadBalancersInboundNatRule_STATUS from the provided source
func (rule *LoadBalancersInboundNatRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == rule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(rule)
}

// ConvertStatusTo populates the provided destination from our LoadBalancersInboundNatRule_STATUS
func (rule *LoadBalancersInboundNatRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == rule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(rule)
}

// Storage version of v1api20240301.LoadBalancersInboundNatRuleOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type LoadBalancersInboundNatRuleOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20240301.NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded
// IPConfiguration in a network interface.
type NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&LoadBalancersInboundNatRule{}, &LoadBalancersInboundNatRuleList{})
}
