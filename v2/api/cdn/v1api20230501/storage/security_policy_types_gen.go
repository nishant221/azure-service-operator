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
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=cdn.azure.com,resources=securitypolicies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cdn.azure.com,resources={securitypolicies/status,securitypolicies/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230501.SecurityPolicy
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/afdx.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/securityPolicies/{securityPolicyName}
type SecurityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityPolicy_Spec   `json:"spec,omitempty"`
	Status            SecurityPolicy_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SecurityPolicy{}

// GetConditions returns the conditions of the resource
func (policy *SecurityPolicy) GetConditions() conditions.Conditions {
	return policy.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (policy *SecurityPolicy) SetConditions(conditions conditions.Conditions) {
	policy.Status.Conditions = conditions
}

var _ configmaps.Exporter = &SecurityPolicy{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (policy *SecurityPolicy) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if policy.Spec.OperatorSpec == nil {
		return nil
	}
	return policy.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &SecurityPolicy{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (policy *SecurityPolicy) SecretDestinationExpressions() []*core.DestinationExpression {
	if policy.Spec.OperatorSpec == nil {
		return nil
	}
	return policy.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &SecurityPolicy{}

// AzureName returns the Azure name of the resource
func (policy *SecurityPolicy) AzureName() string {
	return policy.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (policy SecurityPolicy) GetAPIVersion() string {
	return "2023-05-01"
}

// GetResourceScope returns the scope of the resource
func (policy *SecurityPolicy) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (policy *SecurityPolicy) GetSpec() genruntime.ConvertibleSpec {
	return &policy.Spec
}

// GetStatus returns the status of this resource
func (policy *SecurityPolicy) GetStatus() genruntime.ConvertibleStatus {
	return &policy.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (policy *SecurityPolicy) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles/securityPolicies"
func (policy *SecurityPolicy) GetType() string {
	return "Microsoft.Cdn/profiles/securityPolicies"
}

// NewEmptyStatus returns a new empty (blank) status
func (policy *SecurityPolicy) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SecurityPolicy_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (policy *SecurityPolicy) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(policy.Spec)
	return policy.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (policy *SecurityPolicy) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SecurityPolicy_STATUS); ok {
		policy.Status = *st
		return nil
	}

	// Convert status to required version
	var st SecurityPolicy_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// Hub marks that this SecurityPolicy is the hub type for conversion
func (policy *SecurityPolicy) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *SecurityPolicy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion,
		Kind:    "SecurityPolicy",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230501.SecurityPolicy
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/afdx.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/securityPolicies/{securityPolicyName}
type SecurityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityPolicy `json:"items"`
}

// Storage version of v1api20230501.SecurityPolicy_Spec
type SecurityPolicy_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                      `json:"azureName,omitempty"`
	OperatorSpec    *SecurityPolicyOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                      `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cdn.azure.com/Profile resource
	Owner       *genruntime.KnownResourceReference  `group:"cdn.azure.com" json:"owner,omitempty" kind:"Profile"`
	Parameters  *SecurityPolicyPropertiesParameters `json:"parameters,omitempty"`
	PropertyBag genruntime.PropertyBag              `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &SecurityPolicy_Spec{}

// ConvertSpecFrom populates our SecurityPolicy_Spec from the provided source
func (policy *SecurityPolicy_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == policy {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(policy)
}

// ConvertSpecTo populates the provided destination from our SecurityPolicy_Spec
func (policy *SecurityPolicy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == policy {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(policy)
}

// Storage version of v1api20230501.SecurityPolicy_STATUS
type SecurityPolicy_STATUS struct {
	Conditions        []conditions.Condition                     `json:"conditions,omitempty"`
	DeploymentStatus  *string                                    `json:"deploymentStatus,omitempty"`
	Id                *string                                    `json:"id,omitempty"`
	Name              *string                                    `json:"name,omitempty"`
	Parameters        *SecurityPolicyPropertiesParameters_STATUS `json:"parameters,omitempty"`
	ProfileName       *string                                    `json:"profileName,omitempty"`
	PropertyBag       genruntime.PropertyBag                     `json:"$propertyBag,omitempty"`
	ProvisioningState *string                                    `json:"provisioningState,omitempty"`
	SystemData        *SystemData_STATUS                         `json:"systemData,omitempty"`
	Type              *string                                    `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SecurityPolicy_STATUS{}

// ConvertStatusFrom populates our SecurityPolicy_STATUS from the provided source
func (policy *SecurityPolicy_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == policy {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(policy)
}

// ConvertStatusTo populates the provided destination from our SecurityPolicy_STATUS
func (policy *SecurityPolicy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == policy {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(policy)
}

// Storage version of v1api20230501.SecurityPolicyOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type SecurityPolicyOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20230501.SecurityPolicyPropertiesParameters
type SecurityPolicyPropertiesParameters struct {
	PropertyBag            genruntime.PropertyBag                          `json:"$propertyBag,omitempty"`
	WebApplicationFirewall *SecurityPolicyWebApplicationFirewallParameters `json:"webApplicationFirewall,omitempty"`
}

// Storage version of v1api20230501.SecurityPolicyPropertiesParameters_STATUS
type SecurityPolicyPropertiesParameters_STATUS struct {
	PropertyBag            genruntime.PropertyBag                                 `json:"$propertyBag,omitempty"`
	WebApplicationFirewall *SecurityPolicyWebApplicationFirewallParameters_STATUS `json:"webApplicationFirewall,omitempty"`
}

// Storage version of v1api20230501.SecurityPolicyWebApplicationFirewallParameters
type SecurityPolicyWebApplicationFirewallParameters struct {
	Associations []SecurityPolicyWebApplicationFirewallAssociation `json:"associations,omitempty"`
	PropertyBag  genruntime.PropertyBag                            `json:"$propertyBag,omitempty"`
	Type         *string                                           `json:"type,omitempty"`
	WafPolicy    *ResourceReference                                `json:"wafPolicy,omitempty"`
}

// Storage version of v1api20230501.SecurityPolicyWebApplicationFirewallParameters_STATUS
type SecurityPolicyWebApplicationFirewallParameters_STATUS struct {
	Associations []SecurityPolicyWebApplicationFirewallAssociation_STATUS `json:"associations,omitempty"`
	PropertyBag  genruntime.PropertyBag                                   `json:"$propertyBag,omitempty"`
	Type         *string                                                  `json:"type,omitempty"`
	WafPolicy    *ResourceReference_STATUS                                `json:"wafPolicy,omitempty"`
}

// Storage version of v1api20230501.SecurityPolicyWebApplicationFirewallAssociation
// settings for security policy patterns to match
type SecurityPolicyWebApplicationFirewallAssociation struct {
	Domains         []ActivatedResourceReference `json:"domains,omitempty"`
	PatternsToMatch []string                     `json:"patternsToMatch,omitempty"`
	PropertyBag     genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230501.SecurityPolicyWebApplicationFirewallAssociation_STATUS
// settings for security policy patterns to match
type SecurityPolicyWebApplicationFirewallAssociation_STATUS struct {
	Domains         []ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded `json:"domains,omitempty"`
	PatternsToMatch []string                                                                        `json:"patternsToMatch,omitempty"`
	PropertyBag     genruntime.PropertyBag                                                          `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230501.ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded
// Reference to another resource along with its state.
type ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SecurityPolicy{}, &SecurityPolicyList{})
}
