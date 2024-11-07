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

// +kubebuilder:rbac:groups=apimanagement.azure.com,resources=policies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apimanagement.azure.com,resources={policies/status,policies/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220801.Policy
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimpolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Policy_Spec   `json:"spec,omitempty"`
	Status            Policy_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Policy{}

// GetConditions returns the conditions of the resource
func (policy *Policy) GetConditions() conditions.Conditions {
	return policy.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (policy *Policy) SetConditions(conditions conditions.Conditions) {
	policy.Status.Conditions = conditions
}

var _ configmaps.Exporter = &Policy{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (policy *Policy) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if policy.Spec.OperatorSpec == nil {
		return nil
	}
	return policy.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &Policy{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (policy *Policy) SecretDestinationExpressions() []*core.DestinationExpression {
	if policy.Spec.OperatorSpec == nil {
		return nil
	}
	return policy.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &Policy{}

// AzureName returns the Azure name of the resource (always "policy")
func (policy *Policy) AzureName() string {
	return "policy"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (policy Policy) GetAPIVersion() string {
	return "2022-08-01"
}

// GetResourceScope returns the scope of the resource
func (policy *Policy) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (policy *Policy) GetSpec() genruntime.ConvertibleSpec {
	return &policy.Spec
}

// GetStatus returns the status of this resource
func (policy *Policy) GetStatus() genruntime.ConvertibleStatus {
	return &policy.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (policy *Policy) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationHead,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/policies"
func (policy *Policy) GetType() string {
	return "Microsoft.ApiManagement/service/policies"
}

// NewEmptyStatus returns a new empty (blank) status
func (policy *Policy) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Policy_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (policy *Policy) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(policy.Spec)
	return policy.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (policy *Policy) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Policy_STATUS); ok {
		policy.Status = *st
		return nil
	}

	// Convert status to required version
	var st Policy_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// Hub marks that this Policy is the hub type for conversion
func (policy *Policy) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *Policy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion,
		Kind:    "Policy",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220801.Policy
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimpolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Storage version of v1api20220801.Policy_Spec
type Policy_Spec struct {
	Format          *string             `json:"format,omitempty"`
	OperatorSpec    *PolicyOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string              `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Service resource
	Owner       *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Service"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Value       *string                            `json:"value,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Policy_Spec{}

// ConvertSpecFrom populates our Policy_Spec from the provided source
func (policy *Policy_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(policy)
}

// ConvertSpecTo populates the provided destination from our Policy_Spec
func (policy *Policy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(policy)
}

// Storage version of v1api20220801.Policy_STATUS
type Policy_STATUS struct {
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Format      *string                `json:"format,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Policy_STATUS{}

// ConvertStatusFrom populates our Policy_STATUS from the provided source
func (policy *Policy_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(policy)
}

// ConvertStatusTo populates the provided destination from our Policy_STATUS
func (policy *Policy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(policy)
}

// Storage version of v1api20220801.PolicyOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type PolicyOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
