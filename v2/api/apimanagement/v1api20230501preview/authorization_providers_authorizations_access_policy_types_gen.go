// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501preview

import (
	"fmt"
	arm "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20230501preview/arm"
	storage "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20230501preview/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
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
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimauthorizationproviders.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationProviders/{authorizationProviderId}/authorizations/{authorizationId}/accessPolicies/{authorizationAccessPolicyId}
type AuthorizationProvidersAuthorizationsAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AuthorizationProvidersAuthorizationsAccessPolicy_Spec   `json:"spec,omitempty"`
	Status            AuthorizationProvidersAuthorizationsAccessPolicy_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &AuthorizationProvidersAuthorizationsAccessPolicy{}

// GetConditions returns the conditions of the resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) GetConditions() conditions.Conditions {
	return policy.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) SetConditions(conditions conditions.Conditions) {
	policy.Status.Conditions = conditions
}

var _ conversion.Convertible = &AuthorizationProvidersAuthorizationsAccessPolicy{}

// ConvertFrom populates our AuthorizationProvidersAuthorizationsAccessPolicy from the provided hub AuthorizationProvidersAuthorizationsAccessPolicy
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source storage.AuthorizationProvidersAuthorizationsAccessPolicy

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = policy.AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to policy")
	}

	return nil
}

// ConvertTo populates the provided hub AuthorizationProvidersAuthorizationsAccessPolicy from our AuthorizationProvidersAuthorizationsAccessPolicy
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination storage.AuthorizationProvidersAuthorizationsAccessPolicy
	err := policy.AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from policy")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-apimanagement-azure-com-v1api20230501preview-authorizationprovidersauthorizationsaccesspolicy,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=authorizationprovidersauthorizationsaccesspolicies,verbs=create;update,versions=v1api20230501preview,name=default.v1api20230501preview.authorizationprovidersauthorizationsaccesspolicies.apimanagement.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &AuthorizationProvidersAuthorizationsAccessPolicy{}

// Default applies defaults to the AuthorizationProvidersAuthorizationsAccessPolicy resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) Default() {
	policy.defaultImpl()
	var temp any = policy
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) defaultAzureName() {
	if policy.Spec.AzureName == "" {
		policy.Spec.AzureName = policy.Name
	}
}

// defaultImpl applies the code generated defaults to the AuthorizationProvidersAuthorizationsAccessPolicy resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) defaultImpl() {
	policy.defaultAzureName()
}

var _ genruntime.KubernetesResource = &AuthorizationProvidersAuthorizationsAccessPolicy{}

// AzureName returns the Azure name of the resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) AzureName() string {
	return policy.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01-preview"
func (policy AuthorizationProvidersAuthorizationsAccessPolicy) GetAPIVersion() string {
	return "2023-05-01-preview"
}

// GetResourceScope returns the scope of the resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) GetSpec() genruntime.ConvertibleSpec {
	return &policy.Spec
}

// GetStatus returns the status of this resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) GetStatus() genruntime.ConvertibleStatus {
	return &policy.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/authorizationProviders/authorizations/accessPolicies"
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) GetType() string {
	return "Microsoft.ApiManagement/service/authorizationProviders/authorizations/accessPolicies"
}

// NewEmptyStatus returns a new empty (blank) status
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &AuthorizationProvidersAuthorizationsAccessPolicy_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(policy.Spec)
	return policy.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*AuthorizationProvidersAuthorizationsAccessPolicy_STATUS); ok {
		policy.Status = *st
		return nil
	}

	// Convert status to required version
	var st AuthorizationProvidersAuthorizationsAccessPolicy_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-apimanagement-azure-com-v1api20230501preview-authorizationprovidersauthorizationsaccesspolicy,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=authorizationprovidersauthorizationsaccesspolicies,verbs=create;update,versions=v1api20230501preview,name=validate.v1api20230501preview.authorizationprovidersauthorizationsaccesspolicies.apimanagement.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &AuthorizationProvidersAuthorizationsAccessPolicy{}

// ValidateCreate validates the creation of the resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) ValidateCreate() (admission.Warnings, error) {
	validations := policy.createValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) ValidateDelete() (admission.Warnings, error) {
	validations := policy.deleteValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := policy.updateValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){policy.validateResourceReferences, policy.validateOwnerReference, policy.validateOptionalConfigMapReferences}
}

// deleteValidations validates the deletion of the resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return policy.validateResourceReferences()
		},
		policy.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return policy.validateOwnerReference()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return policy.validateOptionalConfigMapReferences()
		},
	}
}

// validateOptionalConfigMapReferences validates all optional configmap reference pairs to ensure that at most 1 is set
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) validateOptionalConfigMapReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindOptionalConfigMapReferences(&policy.Spec)
	if err != nil {
		return nil, err
	}
	return configmaps.ValidateOptionalReferences(refs)
}

// validateOwnerReference validates the owner field
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(policy)
}

// validateResourceReferences validates all resource references
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&policy.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*AuthorizationProvidersAuthorizationsAccessPolicy)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, policy)
}

// AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy populates our AuthorizationProvidersAuthorizationsAccessPolicy from the provided source AuthorizationProvidersAuthorizationsAccessPolicy
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy(source *storage.AuthorizationProvidersAuthorizationsAccessPolicy) error {

	// ObjectMeta
	policy.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec AuthorizationProvidersAuthorizationsAccessPolicy_Spec
	err := spec.AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy_Spec() to populate field Spec")
	}
	policy.Spec = spec

	// Status
	var status AuthorizationProvidersAuthorizationsAccessPolicy_STATUS
	err = status.AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy_STATUS() to populate field Status")
	}
	policy.Status = status

	// No error
	return nil
}

// AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy populates the provided destination AuthorizationProvidersAuthorizationsAccessPolicy from our AuthorizationProvidersAuthorizationsAccessPolicy
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy(destination *storage.AuthorizationProvidersAuthorizationsAccessPolicy) error {

	// ObjectMeta
	destination.ObjectMeta = *policy.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.AuthorizationProvidersAuthorizationsAccessPolicy_Spec
	err := policy.Spec.AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.AuthorizationProvidersAuthorizationsAccessPolicy_STATUS
	err = policy.Status.AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion(),
		Kind:    "AuthorizationProvidersAuthorizationsAccessPolicy",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimauthorizationproviders.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationProviders/{authorizationProviderId}/authorizations/{authorizationId}/accessPolicies/{authorizationAccessPolicyId}
type AuthorizationProvidersAuthorizationsAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthorizationProvidersAuthorizationsAccessPolicy `json:"items"`
}

type AuthorizationProvidersAuthorizationsAccessPolicy_Spec struct {
	// AppIds: The allowed Azure Active Directory Application IDs
	AppIds []string `json:"appIds,omitempty"`

	// +kubebuilder:validation:MaxLength=256
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern="^[^*#&+:<>?]+$"
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// ObjectId: The Object Id
	ObjectId *string `json:"objectId,omitempty" optionalConfigMapPair:"ObjectId"`

	// ObjectIdFromConfig: The Object Id
	ObjectIdFromConfig *genruntime.ConfigMapReference `json:"objectIdFromConfig,omitempty" optionalConfigMapPair:"ObjectId"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/AuthorizationProvidersAuthorization resource
	Owner *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"AuthorizationProvidersAuthorization"`

	// TenantId: The Tenant Id
	TenantId *string `json:"tenantId,omitempty" optionalConfigMapPair:"TenantId"`

	// TenantIdFromConfig: The Tenant Id
	TenantIdFromConfig *genruntime.ConfigMapReference `json:"tenantIdFromConfig,omitempty" optionalConfigMapPair:"TenantId"`
}

var _ genruntime.ARMTransformer = &AuthorizationProvidersAuthorizationsAccessPolicy_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if policy == nil {
		return nil, nil
	}
	result := &arm.AuthorizationProvidersAuthorizationsAccessPolicy_Spec{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if policy.AppIds != nil ||
		policy.ObjectId != nil ||
		policy.ObjectIdFromConfig != nil ||
		policy.TenantId != nil ||
		policy.TenantIdFromConfig != nil {
		result.Properties = &arm.AuthorizationAccessPolicyContractProperties{}
	}
	for _, item := range policy.AppIds {
		result.Properties.AppIds = append(result.Properties.AppIds, item)
	}
	if policy.ObjectId != nil {
		objectId := *policy.ObjectId
		result.Properties.ObjectId = &objectId
	}
	if policy.ObjectIdFromConfig != nil {
		objectIdValue, err := resolved.ResolvedConfigMaps.Lookup(*policy.ObjectIdFromConfig)
		if err != nil {
			return nil, errors.Wrap(err, "looking up configmap for property ObjectId")
		}
		objectId := objectIdValue
		result.Properties.ObjectId = &objectId
	}
	if policy.TenantId != nil {
		tenantId := *policy.TenantId
		result.Properties.TenantId = &tenantId
	}
	if policy.TenantIdFromConfig != nil {
		tenantIdValue, err := resolved.ResolvedConfigMaps.Lookup(*policy.TenantIdFromConfig)
		if err != nil {
			return nil, errors.Wrap(err, "looking up configmap for property TenantId")
		}
		tenantId := tenantIdValue
		result.Properties.TenantId = &tenantId
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.AuthorizationProvidersAuthorizationsAccessPolicy_Spec{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.AuthorizationProvidersAuthorizationsAccessPolicy_Spec)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.AuthorizationProvidersAuthorizationsAccessPolicy_Spec, got %T", armInput)
	}

	// Set property "AppIds":
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.AppIds {
			policy.AppIds = append(policy.AppIds, item)
		}
	}

	// Set property "AzureName":
	policy.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "ObjectId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ObjectId != nil {
			objectId := *typedInput.Properties.ObjectId
			policy.ObjectId = &objectId
		}
	}

	// no assignment for property "ObjectIdFromConfig"

	// Set property "Owner":
	policy.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "TenantId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.TenantId != nil {
			tenantId := *typedInput.Properties.TenantId
			policy.TenantId = &tenantId
		}
	}

	// no assignment for property "TenantIdFromConfig"

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &AuthorizationProvidersAuthorizationsAccessPolicy_Spec{}

// ConvertSpecFrom populates our AuthorizationProvidersAuthorizationsAccessPolicy_Spec from the provided source
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.AuthorizationProvidersAuthorizationsAccessPolicy_Spec)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.AuthorizationProvidersAuthorizationsAccessPolicy_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our AuthorizationProvidersAuthorizationsAccessPolicy_Spec
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.AuthorizationProvidersAuthorizationsAccessPolicy_Spec)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.AuthorizationProvidersAuthorizationsAccessPolicy_Spec{}
	err := policy.AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy_Spec(dst)
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

// AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy_Spec populates our AuthorizationProvidersAuthorizationsAccessPolicy_Spec from the provided source AuthorizationProvidersAuthorizationsAccessPolicy_Spec
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_Spec) AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy_Spec(source *storage.AuthorizationProvidersAuthorizationsAccessPolicy_Spec) error {

	// AppIds
	policy.AppIds = genruntime.CloneSliceOfString(source.AppIds)

	// AzureName
	policy.AzureName = source.AzureName

	// ObjectId
	policy.ObjectId = genruntime.ClonePointerToString(source.ObjectId)

	// ObjectIdFromConfig
	if source.ObjectIdFromConfig != nil {
		objectIdFromConfig := source.ObjectIdFromConfig.Copy()
		policy.ObjectIdFromConfig = &objectIdFromConfig
	} else {
		policy.ObjectIdFromConfig = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		policy.Owner = &owner
	} else {
		policy.Owner = nil
	}

	// TenantId
	policy.TenantId = genruntime.ClonePointerToString(source.TenantId)

	// TenantIdFromConfig
	if source.TenantIdFromConfig != nil {
		tenantIdFromConfig := source.TenantIdFromConfig.Copy()
		policy.TenantIdFromConfig = &tenantIdFromConfig
	} else {
		policy.TenantIdFromConfig = nil
	}

	// No error
	return nil
}

// AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy_Spec populates the provided destination AuthorizationProvidersAuthorizationsAccessPolicy_Spec from our AuthorizationProvidersAuthorizationsAccessPolicy_Spec
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_Spec) AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy_Spec(destination *storage.AuthorizationProvidersAuthorizationsAccessPolicy_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AppIds
	destination.AppIds = genruntime.CloneSliceOfString(policy.AppIds)

	// AzureName
	destination.AzureName = policy.AzureName

	// ObjectId
	destination.ObjectId = genruntime.ClonePointerToString(policy.ObjectId)

	// ObjectIdFromConfig
	if policy.ObjectIdFromConfig != nil {
		objectIdFromConfig := policy.ObjectIdFromConfig.Copy()
		destination.ObjectIdFromConfig = &objectIdFromConfig
	} else {
		destination.ObjectIdFromConfig = nil
	}

	// OriginalVersion
	destination.OriginalVersion = policy.OriginalVersion()

	// Owner
	if policy.Owner != nil {
		owner := policy.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// TenantId
	destination.TenantId = genruntime.ClonePointerToString(policy.TenantId)

	// TenantIdFromConfig
	if policy.TenantIdFromConfig != nil {
		tenantIdFromConfig := policy.TenantIdFromConfig.Copy()
		destination.TenantIdFromConfig = &tenantIdFromConfig
	} else {
		destination.TenantIdFromConfig = nil
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

// OriginalVersion returns the original API version used to create the resource.
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_Spec) SetAzureName(azureName string) {
	policy.AzureName = azureName
}

type AuthorizationProvidersAuthorizationsAccessPolicy_STATUS struct {
	// AppIds: The allowed Azure Active Directory Application IDs
	AppIds []string `json:"appIds,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// ObjectId: The Object Id
	ObjectId *string `json:"objectId,omitempty"`

	// TenantId: The Tenant Id
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &AuthorizationProvidersAuthorizationsAccessPolicy_STATUS{}

// ConvertStatusFrom populates our AuthorizationProvidersAuthorizationsAccessPolicy_STATUS from the provided source
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.AuthorizationProvidersAuthorizationsAccessPolicy_STATUS)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.AuthorizationProvidersAuthorizationsAccessPolicy_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our AuthorizationProvidersAuthorizationsAccessPolicy_STATUS
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.AuthorizationProvidersAuthorizationsAccessPolicy_STATUS)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.AuthorizationProvidersAuthorizationsAccessPolicy_STATUS{}
	err := policy.AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy_STATUS(dst)
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

var _ genruntime.FromARMConverter = &AuthorizationProvidersAuthorizationsAccessPolicy_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.AuthorizationProvidersAuthorizationsAccessPolicy_STATUS{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.AuthorizationProvidersAuthorizationsAccessPolicy_STATUS)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.AuthorizationProvidersAuthorizationsAccessPolicy_STATUS, got %T", armInput)
	}

	// Set property "AppIds":
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.AppIds {
			policy.AppIds = append(policy.AppIds, item)
		}
	}

	// no assignment for property "Conditions"

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		policy.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		policy.Name = &name
	}

	// Set property "ObjectId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ObjectId != nil {
			objectId := *typedInput.Properties.ObjectId
			policy.ObjectId = &objectId
		}
	}

	// Set property "TenantId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.TenantId != nil {
			tenantId := *typedInput.Properties.TenantId
			policy.TenantId = &tenantId
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		policy.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy_STATUS populates our AuthorizationProvidersAuthorizationsAccessPolicy_STATUS from the provided source AuthorizationProvidersAuthorizationsAccessPolicy_STATUS
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_STATUS) AssignProperties_From_AuthorizationProvidersAuthorizationsAccessPolicy_STATUS(source *storage.AuthorizationProvidersAuthorizationsAccessPolicy_STATUS) error {

	// AppIds
	policy.AppIds = genruntime.CloneSliceOfString(source.AppIds)

	// Conditions
	policy.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	policy.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	policy.Name = genruntime.ClonePointerToString(source.Name)

	// ObjectId
	policy.ObjectId = genruntime.ClonePointerToString(source.ObjectId)

	// TenantId
	policy.TenantId = genruntime.ClonePointerToString(source.TenantId)

	// Type
	policy.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy_STATUS populates the provided destination AuthorizationProvidersAuthorizationsAccessPolicy_STATUS from our AuthorizationProvidersAuthorizationsAccessPolicy_STATUS
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_STATUS) AssignProperties_To_AuthorizationProvidersAuthorizationsAccessPolicy_STATUS(destination *storage.AuthorizationProvidersAuthorizationsAccessPolicy_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AppIds
	destination.AppIds = genruntime.CloneSliceOfString(policy.AppIds)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(policy.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(policy.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(policy.Name)

	// ObjectId
	destination.ObjectId = genruntime.ClonePointerToString(policy.ObjectId)

	// TenantId
	destination.TenantId = genruntime.ClonePointerToString(policy.TenantId)

	// Type
	destination.Type = genruntime.ClonePointerToString(policy.Type)

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
	SchemeBuilder.Register(&AuthorizationProvidersAuthorizationsAccessPolicy{}, &AuthorizationProvidersAuthorizationsAccessPolicyList{})
}
