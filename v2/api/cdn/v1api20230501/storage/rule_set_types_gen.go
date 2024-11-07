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

// +kubebuilder:rbac:groups=cdn.azure.com,resources=rulesets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cdn.azure.com,resources={rulesets/status,rulesets/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230501.RuleSet
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/afdx.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/ruleSets/{ruleSetName}
type RuleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuleSet_Spec   `json:"spec,omitempty"`
	Status            RuleSet_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RuleSet{}

// GetConditions returns the conditions of the resource
func (ruleSet *RuleSet) GetConditions() conditions.Conditions {
	return ruleSet.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (ruleSet *RuleSet) SetConditions(conditions conditions.Conditions) {
	ruleSet.Status.Conditions = conditions
}

var _ configmaps.Exporter = &RuleSet{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (ruleSet *RuleSet) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if ruleSet.Spec.OperatorSpec == nil {
		return nil
	}
	return ruleSet.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &RuleSet{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (ruleSet *RuleSet) SecretDestinationExpressions() []*core.DestinationExpression {
	if ruleSet.Spec.OperatorSpec == nil {
		return nil
	}
	return ruleSet.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &RuleSet{}

// AzureName returns the Azure name of the resource
func (ruleSet *RuleSet) AzureName() string {
	return ruleSet.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (ruleSet RuleSet) GetAPIVersion() string {
	return "2023-05-01"
}

// GetResourceScope returns the scope of the resource
func (ruleSet *RuleSet) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (ruleSet *RuleSet) GetSpec() genruntime.ConvertibleSpec {
	return &ruleSet.Spec
}

// GetStatus returns the status of this resource
func (ruleSet *RuleSet) GetStatus() genruntime.ConvertibleStatus {
	return &ruleSet.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (ruleSet *RuleSet) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles/ruleSets"
func (ruleSet *RuleSet) GetType() string {
	return "Microsoft.Cdn/profiles/ruleSets"
}

// NewEmptyStatus returns a new empty (blank) status
func (ruleSet *RuleSet) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RuleSet_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (ruleSet *RuleSet) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(ruleSet.Spec)
	return ruleSet.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (ruleSet *RuleSet) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RuleSet_STATUS); ok {
		ruleSet.Status = *st
		return nil
	}

	// Convert status to required version
	var st RuleSet_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	ruleSet.Status = st
	return nil
}

// Hub marks that this RuleSet is the hub type for conversion
func (ruleSet *RuleSet) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (ruleSet *RuleSet) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: ruleSet.Spec.OriginalVersion,
		Kind:    "RuleSet",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230501.RuleSet
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/afdx.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/ruleSets/{ruleSetName}
type RuleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RuleSet `json:"items"`
}

// Storage version of v1api20230501.RuleSet_Spec
type RuleSet_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string               `json:"azureName,omitempty"`
	OperatorSpec    *RuleSetOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cdn.azure.com/Profile resource
	Owner       *genruntime.KnownResourceReference `group:"cdn.azure.com" json:"owner,omitempty" kind:"Profile"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RuleSet_Spec{}

// ConvertSpecFrom populates our RuleSet_Spec from the provided source
func (ruleSet *RuleSet_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == ruleSet {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(ruleSet)
}

// ConvertSpecTo populates the provided destination from our RuleSet_Spec
func (ruleSet *RuleSet_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == ruleSet {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(ruleSet)
}

// Storage version of v1api20230501.RuleSet_STATUS
type RuleSet_STATUS struct {
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	DeploymentStatus  *string                `json:"deploymentStatus,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	ProfileName       *string                `json:"profileName,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	SystemData        *SystemData_STATUS     `json:"systemData,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RuleSet_STATUS{}

// ConvertStatusFrom populates our RuleSet_STATUS from the provided source
func (ruleSet *RuleSet_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == ruleSet {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(ruleSet)
}

// ConvertStatusTo populates the provided destination from our RuleSet_STATUS
func (ruleSet *RuleSet_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == ruleSet {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(ruleSet)
}

// Storage version of v1api20230501.RuleSetOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type RuleSetOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&RuleSet{}, &RuleSetList{})
}
