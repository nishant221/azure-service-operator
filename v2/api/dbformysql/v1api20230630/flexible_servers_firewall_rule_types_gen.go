// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230630

import (
	"fmt"
	arm "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20230630/arm"
	storage "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20230630/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /mysql/resource-manager/Microsoft.DBforMySQL/Firewall/stable/2023-06-30/FirewallRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/firewallRules/{firewallRuleName}
type FlexibleServersFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServersFirewallRule_Spec   `json:"spec,omitempty"`
	Status            FlexibleServersFirewallRule_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServersFirewallRule{}

// GetConditions returns the conditions of the resource
func (rule *FlexibleServersFirewallRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *FlexibleServersFirewallRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ conversion.Convertible = &FlexibleServersFirewallRule{}

// ConvertFrom populates our FlexibleServersFirewallRule from the provided hub FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.FlexibleServersFirewallRule)
	if !ok {
		return fmt.Errorf("expected dbformysql/v1api20230630/storage/FlexibleServersFirewallRule but received %T instead", hub)
	}

	return rule.AssignProperties_From_FlexibleServersFirewallRule(source)
}

// ConvertTo populates the provided hub FlexibleServersFirewallRule from our FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.FlexibleServersFirewallRule)
	if !ok {
		return fmt.Errorf("expected dbformysql/v1api20230630/storage/FlexibleServersFirewallRule but received %T instead", hub)
	}

	return rule.AssignProperties_To_FlexibleServersFirewallRule(destination)
}

// +kubebuilder:webhook:path=/mutate-dbformysql-azure-com-v1api20230630-flexibleserversfirewallrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformysql.azure.com,resources=flexibleserversfirewallrules,verbs=create;update,versions=v1api20230630,name=default.v1api20230630.flexibleserversfirewallrules.dbformysql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &FlexibleServersFirewallRule{}

// Default applies defaults to the FlexibleServersFirewallRule resource
func (rule *FlexibleServersFirewallRule) Default() {
	rule.defaultImpl()
	var temp any = rule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (rule *FlexibleServersFirewallRule) defaultAzureName() {
	if rule.Spec.AzureName == "" {
		rule.Spec.AzureName = rule.Name
	}
}

// defaultImpl applies the code generated defaults to the FlexibleServersFirewallRule resource
func (rule *FlexibleServersFirewallRule) defaultImpl() { rule.defaultAzureName() }

var _ configmaps.Exporter = &FlexibleServersFirewallRule{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (rule *FlexibleServersFirewallRule) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if rule.Spec.OperatorSpec == nil {
		return nil
	}
	return rule.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &FlexibleServersFirewallRule{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (rule *FlexibleServersFirewallRule) SecretDestinationExpressions() []*core.DestinationExpression {
	if rule.Spec.OperatorSpec == nil {
		return nil
	}
	return rule.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.ImportableResource = &FlexibleServersFirewallRule{}

// InitializeSpec initializes the spec for this resource from the given status
func (rule *FlexibleServersFirewallRule) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*FlexibleServersFirewallRule_STATUS); ok {
		return rule.Spec.Initialize_From_FlexibleServersFirewallRule_STATUS(s)
	}

	return fmt.Errorf("expected Status of type FlexibleServersFirewallRule_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &FlexibleServersFirewallRule{}

// AzureName returns the Azure name of the resource
func (rule *FlexibleServersFirewallRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-06-30"
func (rule FlexibleServersFirewallRule) GetAPIVersion() string {
	return "2023-06-30"
}

// GetResourceScope returns the scope of the resource
func (rule *FlexibleServersFirewallRule) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (rule *FlexibleServersFirewallRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *FlexibleServersFirewallRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (rule *FlexibleServersFirewallRule) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/firewallRules"
func (rule *FlexibleServersFirewallRule) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/firewallRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *FlexibleServersFirewallRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FlexibleServersFirewallRule_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (rule *FlexibleServersFirewallRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return rule.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (rule *FlexibleServersFirewallRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FlexibleServersFirewallRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st FlexibleServersFirewallRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-dbformysql-azure-com-v1api20230630-flexibleserversfirewallrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformysql.azure.com,resources=flexibleserversfirewallrules,verbs=create;update,versions=v1api20230630,name=validate.v1api20230630.flexibleserversfirewallrules.dbformysql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &FlexibleServersFirewallRule{}

// ValidateCreate validates the creation of the resource
func (rule *FlexibleServersFirewallRule) ValidateCreate() (admission.Warnings, error) {
	validations := rule.createValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (rule *FlexibleServersFirewallRule) ValidateDelete() (admission.Warnings, error) {
	validations := rule.deleteValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (rule *FlexibleServersFirewallRule) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := rule.updateValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (rule *FlexibleServersFirewallRule) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){rule.validateResourceReferences, rule.validateOwnerReference, rule.validateSecretDestinations, rule.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (rule *FlexibleServersFirewallRule) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (rule *FlexibleServersFirewallRule) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return rule.validateResourceReferences()
		},
		rule.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return rule.validateOwnerReference()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return rule.validateSecretDestinations()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return rule.validateConfigMapDestinations()
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (rule *FlexibleServersFirewallRule) validateConfigMapDestinations() (admission.Warnings, error) {
	if rule.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(rule, nil, rule.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (rule *FlexibleServersFirewallRule) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(rule)
}

// validateResourceReferences validates all resource references
func (rule *FlexibleServersFirewallRule) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (rule *FlexibleServersFirewallRule) validateSecretDestinations() (admission.Warnings, error) {
	if rule.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(rule, nil, rule.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (rule *FlexibleServersFirewallRule) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*FlexibleServersFirewallRule)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, rule)
}

// AssignProperties_From_FlexibleServersFirewallRule populates our FlexibleServersFirewallRule from the provided source FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) AssignProperties_From_FlexibleServersFirewallRule(source *storage.FlexibleServersFirewallRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServersFirewallRule_Spec
	err := spec.AssignProperties_From_FlexibleServersFirewallRule_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_FlexibleServersFirewallRule_Spec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status FlexibleServersFirewallRule_STATUS
	err = status.AssignProperties_From_FlexibleServersFirewallRule_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_FlexibleServersFirewallRule_STATUS() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersFirewallRule populates the provided destination FlexibleServersFirewallRule from our FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) AssignProperties_To_FlexibleServersFirewallRule(destination *storage.FlexibleServersFirewallRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.FlexibleServersFirewallRule_Spec
	err := rule.Spec.AssignProperties_To_FlexibleServersFirewallRule_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_FlexibleServersFirewallRule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.FlexibleServersFirewallRule_STATUS
	err = rule.Status.AssignProperties_To_FlexibleServersFirewallRule_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_FlexibleServersFirewallRule_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *FlexibleServersFirewallRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion(),
		Kind:    "FlexibleServersFirewallRule",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /mysql/resource-manager/Microsoft.DBforMySQL/Firewall/stable/2023-06-30/FirewallRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/firewallRules/{firewallRuleName}
type FlexibleServersFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersFirewallRule `json:"items"`
}

type FlexibleServersFirewallRule_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	// EndIpAddress: The end IP address of the server firewall rule. Must be IPv4 format.
	EndIpAddress *string `json:"endIpAddress,omitempty"`

	// OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
	// passed directly to Azure
	OperatorSpec *FlexibleServersFirewallRuleOperatorSpec `json:"operatorSpec,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbformysql.azure.com/FlexibleServer resource
	Owner *genruntime.KnownResourceReference `group:"dbformysql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	// StartIpAddress: The start IP address of the server firewall rule. Must be IPv4 format.
	StartIpAddress *string `json:"startIpAddress,omitempty"`
}

var _ genruntime.ARMTransformer = &FlexibleServersFirewallRule_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (rule *FlexibleServersFirewallRule_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if rule == nil {
		return nil, nil
	}
	result := &arm.FlexibleServersFirewallRule_Spec{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if rule.EndIpAddress != nil || rule.StartIpAddress != nil {
		result.Properties = &arm.FirewallRuleProperties{}
	}
	if rule.EndIpAddress != nil {
		endIpAddress := *rule.EndIpAddress
		result.Properties.EndIpAddress = &endIpAddress
	}
	if rule.StartIpAddress != nil {
		startIpAddress := *rule.StartIpAddress
		result.Properties.StartIpAddress = &startIpAddress
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *FlexibleServersFirewallRule_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.FlexibleServersFirewallRule_Spec{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *FlexibleServersFirewallRule_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.FlexibleServersFirewallRule_Spec)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.FlexibleServersFirewallRule_Spec, got %T", armInput)
	}

	// Set property "AzureName":
	rule.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "EndIpAddress":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EndIpAddress != nil {
			endIpAddress := *typedInput.Properties.EndIpAddress
			rule.EndIpAddress = &endIpAddress
		}
	}

	// no assignment for property "OperatorSpec"

	// Set property "Owner":
	rule.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "StartIpAddress":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StartIpAddress != nil {
			startIpAddress := *typedInput.Properties.StartIpAddress
			rule.StartIpAddress = &startIpAddress
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &FlexibleServersFirewallRule_Spec{}

// ConvertSpecFrom populates our FlexibleServersFirewallRule_Spec from the provided source
func (rule *FlexibleServersFirewallRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.FlexibleServersFirewallRule_Spec)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_FlexibleServersFirewallRule_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.FlexibleServersFirewallRule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_FlexibleServersFirewallRule_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServersFirewallRule_Spec
func (rule *FlexibleServersFirewallRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.FlexibleServersFirewallRule_Spec)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_FlexibleServersFirewallRule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.FlexibleServersFirewallRule_Spec{}
	err := rule.AssignProperties_To_FlexibleServersFirewallRule_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_FlexibleServersFirewallRule_Spec populates our FlexibleServersFirewallRule_Spec from the provided source FlexibleServersFirewallRule_Spec
func (rule *FlexibleServersFirewallRule_Spec) AssignProperties_From_FlexibleServersFirewallRule_Spec(source *storage.FlexibleServersFirewallRule_Spec) error {

	// AzureName
	rule.AzureName = source.AzureName

	// EndIpAddress
	if source.EndIpAddress != nil {
		endIpAddress := *source.EndIpAddress
		rule.EndIpAddress = &endIpAddress
	} else {
		rule.EndIpAddress = nil
	}

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec FlexibleServersFirewallRuleOperatorSpec
		err := operatorSpec.AssignProperties_From_FlexibleServersFirewallRuleOperatorSpec(source.OperatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_FlexibleServersFirewallRuleOperatorSpec() to populate field OperatorSpec")
		}
		rule.OperatorSpec = &operatorSpec
	} else {
		rule.OperatorSpec = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rule.Owner = &owner
	} else {
		rule.Owner = nil
	}

	// StartIpAddress
	if source.StartIpAddress != nil {
		startIpAddress := *source.StartIpAddress
		rule.StartIpAddress = &startIpAddress
	} else {
		rule.StartIpAddress = nil
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersFirewallRule_Spec populates the provided destination FlexibleServersFirewallRule_Spec from our FlexibleServersFirewallRule_Spec
func (rule *FlexibleServersFirewallRule_Spec) AssignProperties_To_FlexibleServersFirewallRule_Spec(destination *storage.FlexibleServersFirewallRule_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = rule.AzureName

	// EndIpAddress
	if rule.EndIpAddress != nil {
		endIpAddress := *rule.EndIpAddress
		destination.EndIpAddress = &endIpAddress
	} else {
		destination.EndIpAddress = nil
	}

	// OperatorSpec
	if rule.OperatorSpec != nil {
		var operatorSpec storage.FlexibleServersFirewallRuleOperatorSpec
		err := rule.OperatorSpec.AssignProperties_To_FlexibleServersFirewallRuleOperatorSpec(&operatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_FlexibleServersFirewallRuleOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = rule.OriginalVersion()

	// Owner
	if rule.Owner != nil {
		owner := rule.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// StartIpAddress
	if rule.StartIpAddress != nil {
		startIpAddress := *rule.StartIpAddress
		destination.StartIpAddress = &startIpAddress
	} else {
		destination.StartIpAddress = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_FlexibleServersFirewallRule_STATUS populates our FlexibleServersFirewallRule_Spec from the provided source FlexibleServersFirewallRule_STATUS
func (rule *FlexibleServersFirewallRule_Spec) Initialize_From_FlexibleServersFirewallRule_STATUS(source *FlexibleServersFirewallRule_STATUS) error {

	// EndIpAddress
	if source.EndIpAddress != nil {
		endIpAddress := *source.EndIpAddress
		rule.EndIpAddress = &endIpAddress
	} else {
		rule.EndIpAddress = nil
	}

	// StartIpAddress
	if source.StartIpAddress != nil {
		startIpAddress := *source.StartIpAddress
		rule.StartIpAddress = &startIpAddress
	} else {
		rule.StartIpAddress = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (rule *FlexibleServersFirewallRule_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (rule *FlexibleServersFirewallRule_Spec) SetAzureName(azureName string) {
	rule.AzureName = azureName
}

type FlexibleServersFirewallRule_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// EndIpAddress: The end IP address of the server firewall rule. Must be IPv4 format.
	EndIpAddress *string `json:"endIpAddress,omitempty"`

	// Id: Fully qualified resource ID for the resource. E.g.
	// "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// StartIpAddress: The start IP address of the server firewall rule. Must be IPv4 format.
	StartIpAddress *string `json:"startIpAddress,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FlexibleServersFirewallRule_STATUS{}

// ConvertStatusFrom populates our FlexibleServersFirewallRule_STATUS from the provided source
func (rule *FlexibleServersFirewallRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.FlexibleServersFirewallRule_STATUS)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_FlexibleServersFirewallRule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.FlexibleServersFirewallRule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_FlexibleServersFirewallRule_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our FlexibleServersFirewallRule_STATUS
func (rule *FlexibleServersFirewallRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.FlexibleServersFirewallRule_STATUS)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_FlexibleServersFirewallRule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.FlexibleServersFirewallRule_STATUS{}
	err := rule.AssignProperties_To_FlexibleServersFirewallRule_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &FlexibleServersFirewallRule_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *FlexibleServersFirewallRule_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.FlexibleServersFirewallRule_STATUS{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *FlexibleServersFirewallRule_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.FlexibleServersFirewallRule_STATUS)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.FlexibleServersFirewallRule_STATUS, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "EndIpAddress":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EndIpAddress != nil {
			endIpAddress := *typedInput.Properties.EndIpAddress
			rule.EndIpAddress = &endIpAddress
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		rule.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		rule.Name = &name
	}

	// Set property "StartIpAddress":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StartIpAddress != nil {
			startIpAddress := *typedInput.Properties.StartIpAddress
			rule.StartIpAddress = &startIpAddress
		}
	}

	// Set property "SystemData":
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		rule.SystemData = &systemData
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		rule.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_FlexibleServersFirewallRule_STATUS populates our FlexibleServersFirewallRule_STATUS from the provided source FlexibleServersFirewallRule_STATUS
func (rule *FlexibleServersFirewallRule_STATUS) AssignProperties_From_FlexibleServersFirewallRule_STATUS(source *storage.FlexibleServersFirewallRule_STATUS) error {

	// Conditions
	rule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// EndIpAddress
	rule.EndIpAddress = genruntime.ClonePointerToString(source.EndIpAddress)

	// Id
	rule.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	rule.Name = genruntime.ClonePointerToString(source.Name)

	// StartIpAddress
	rule.StartIpAddress = genruntime.ClonePointerToString(source.StartIpAddress)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		rule.SystemData = &systemDatum
	} else {
		rule.SystemData = nil
	}

	// Type
	rule.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersFirewallRule_STATUS populates the provided destination FlexibleServersFirewallRule_STATUS from our FlexibleServersFirewallRule_STATUS
func (rule *FlexibleServersFirewallRule_STATUS) AssignProperties_To_FlexibleServersFirewallRule_STATUS(destination *storage.FlexibleServersFirewallRule_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(rule.Conditions)

	// EndIpAddress
	destination.EndIpAddress = genruntime.ClonePointerToString(rule.EndIpAddress)

	// Id
	destination.Id = genruntime.ClonePointerToString(rule.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(rule.Name)

	// StartIpAddress
	destination.StartIpAddress = genruntime.ClonePointerToString(rule.StartIpAddress)

	// SystemData
	if rule.SystemData != nil {
		var systemDatum storage.SystemData_STATUS
		err := rule.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(rule.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type FlexibleServersFirewallRuleOperatorSpec struct {
	// ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`

	// SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).
	SecretExpressions []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_FlexibleServersFirewallRuleOperatorSpec populates our FlexibleServersFirewallRuleOperatorSpec from the provided source FlexibleServersFirewallRuleOperatorSpec
func (operator *FlexibleServersFirewallRuleOperatorSpec) AssignProperties_From_FlexibleServersFirewallRuleOperatorSpec(source *storage.FlexibleServersFirewallRuleOperatorSpec) error {

	// ConfigMapExpressions
	if source.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(source.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range source.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		operator.ConfigMapExpressions = configMapExpressionList
	} else {
		operator.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if source.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(source.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range source.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		operator.SecretExpressions = secretExpressionList
	} else {
		operator.SecretExpressions = nil
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersFirewallRuleOperatorSpec populates the provided destination FlexibleServersFirewallRuleOperatorSpec from our FlexibleServersFirewallRuleOperatorSpec
func (operator *FlexibleServersFirewallRuleOperatorSpec) AssignProperties_To_FlexibleServersFirewallRuleOperatorSpec(destination *storage.FlexibleServersFirewallRuleOperatorSpec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// ConfigMapExpressions
	if operator.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(operator.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range operator.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		destination.ConfigMapExpressions = configMapExpressionList
	} else {
		destination.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if operator.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(operator.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range operator.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		destination.SecretExpressions = secretExpressionList
	} else {
		destination.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&FlexibleServersFirewallRule{}, &FlexibleServersFirewallRuleList{})
}
