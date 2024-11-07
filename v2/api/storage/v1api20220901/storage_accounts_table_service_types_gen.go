// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

import (
	"fmt"
	arm "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901/arm"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901/storage"
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
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2022-09-01/table.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default
type StorageAccountsTableService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountsTableService_Spec   `json:"spec,omitempty"`
	Status            StorageAccountsTableService_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsTableService{}

// GetConditions returns the conditions of the resource
func (service *StorageAccountsTableService) GetConditions() conditions.Conditions {
	return service.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (service *StorageAccountsTableService) SetConditions(conditions conditions.Conditions) {
	service.Status.Conditions = conditions
}

var _ conversion.Convertible = &StorageAccountsTableService{}

// ConvertFrom populates our StorageAccountsTableService from the provided hub StorageAccountsTableService
func (service *StorageAccountsTableService) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source storage.StorageAccountsTableService

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = service.AssignProperties_From_StorageAccountsTableService(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to service")
	}

	return nil
}

// ConvertTo populates the provided hub StorageAccountsTableService from our StorageAccountsTableService
func (service *StorageAccountsTableService) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination storage.StorageAccountsTableService
	err := service.AssignProperties_To_StorageAccountsTableService(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from service")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-storage-azure-com-v1api20220901-storageaccountstableservice,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountstableservices,verbs=create;update,versions=v1api20220901,name=default.v1api20220901.storageaccountstableservices.storage.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &StorageAccountsTableService{}

// Default applies defaults to the StorageAccountsTableService resource
func (service *StorageAccountsTableService) Default() {
	service.defaultImpl()
	var temp any = service
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the StorageAccountsTableService resource
func (service *StorageAccountsTableService) defaultImpl() {}

var _ configmaps.Exporter = &StorageAccountsTableService{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (service *StorageAccountsTableService) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if service.Spec.OperatorSpec == nil {
		return nil
	}
	return service.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &StorageAccountsTableService{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (service *StorageAccountsTableService) SecretDestinationExpressions() []*core.DestinationExpression {
	if service.Spec.OperatorSpec == nil {
		return nil
	}
	return service.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &StorageAccountsTableService{}

// AzureName returns the Azure name of the resource (always "default")
func (service *StorageAccountsTableService) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-09-01"
func (service StorageAccountsTableService) GetAPIVersion() string {
	return "2022-09-01"
}

// GetResourceScope returns the scope of the resource
func (service *StorageAccountsTableService) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (service *StorageAccountsTableService) GetSpec() genruntime.ConvertibleSpec {
	return &service.Spec
}

// GetStatus returns the status of this resource
func (service *StorageAccountsTableService) GetStatus() genruntime.ConvertibleStatus {
	return &service.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (service *StorageAccountsTableService) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/tableServices"
func (service *StorageAccountsTableService) GetType() string {
	return "Microsoft.Storage/storageAccounts/tableServices"
}

// NewEmptyStatus returns a new empty (blank) status
func (service *StorageAccountsTableService) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageAccountsTableService_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (service *StorageAccountsTableService) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(service.Spec)
	return service.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (service *StorageAccountsTableService) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*StorageAccountsTableService_STATUS); ok {
		service.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageAccountsTableService_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	service.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-storage-azure-com-v1api20220901-storageaccountstableservice,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountstableservices,verbs=create;update,versions=v1api20220901,name=validate.v1api20220901.storageaccountstableservices.storage.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &StorageAccountsTableService{}

// ValidateCreate validates the creation of the resource
func (service *StorageAccountsTableService) ValidateCreate() (admission.Warnings, error) {
	validations := service.createValidations()
	var temp any = service
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (service *StorageAccountsTableService) ValidateDelete() (admission.Warnings, error) {
	validations := service.deleteValidations()
	var temp any = service
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (service *StorageAccountsTableService) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := service.updateValidations()
	var temp any = service
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (service *StorageAccountsTableService) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){service.validateResourceReferences, service.validateOwnerReference, service.validateSecretDestinations, service.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (service *StorageAccountsTableService) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (service *StorageAccountsTableService) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return service.validateResourceReferences()
		},
		service.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return service.validateOwnerReference()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return service.validateSecretDestinations()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return service.validateConfigMapDestinations()
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (service *StorageAccountsTableService) validateConfigMapDestinations() (admission.Warnings, error) {
	if service.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(service, nil, service.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (service *StorageAccountsTableService) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(service)
}

// validateResourceReferences validates all resource references
func (service *StorageAccountsTableService) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&service.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (service *StorageAccountsTableService) validateSecretDestinations() (admission.Warnings, error) {
	if service.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(service, nil, service.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (service *StorageAccountsTableService) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*StorageAccountsTableService)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, service)
}

// AssignProperties_From_StorageAccountsTableService populates our StorageAccountsTableService from the provided source StorageAccountsTableService
func (service *StorageAccountsTableService) AssignProperties_From_StorageAccountsTableService(source *storage.StorageAccountsTableService) error {

	// ObjectMeta
	service.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec StorageAccountsTableService_Spec
	err := spec.AssignProperties_From_StorageAccountsTableService_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_StorageAccountsTableService_Spec() to populate field Spec")
	}
	service.Spec = spec

	// Status
	var status StorageAccountsTableService_STATUS
	err = status.AssignProperties_From_StorageAccountsTableService_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_StorageAccountsTableService_STATUS() to populate field Status")
	}
	service.Status = status

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsTableService populates the provided destination StorageAccountsTableService from our StorageAccountsTableService
func (service *StorageAccountsTableService) AssignProperties_To_StorageAccountsTableService(destination *storage.StorageAccountsTableService) error {

	// ObjectMeta
	destination.ObjectMeta = *service.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.StorageAccountsTableService_Spec
	err := service.Spec.AssignProperties_To_StorageAccountsTableService_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_StorageAccountsTableService_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.StorageAccountsTableService_STATUS
	err = service.Status.AssignProperties_To_StorageAccountsTableService_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_StorageAccountsTableService_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (service *StorageAccountsTableService) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: service.Spec.OriginalVersion(),
		Kind:    "StorageAccountsTableService",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2022-09-01/table.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default
type StorageAccountsTableServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsTableService `json:"items"`
}

type StorageAccountsTableService_Spec struct {
	// Cors: Specifies CORS rules for the Table service. You can include up to five CorsRule elements in the request. If no
	// CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the
	// Table service.
	Cors *CorsRules `json:"cors,omitempty"`

	// OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
	// passed directly to Azure
	OperatorSpec *StorageAccountsTableServiceOperatorSpec `json:"operatorSpec,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a storage.azure.com/StorageAccount resource
	Owner *genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner,omitempty" kind:"StorageAccount"`
}

var _ genruntime.ARMTransformer = &StorageAccountsTableService_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (service *StorageAccountsTableService_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if service == nil {
		return nil, nil
	}
	result := &arm.StorageAccountsTableService_Spec{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if service.Cors != nil {
		result.Properties = &arm.StorageAccounts_TableService_Properties_Spec{}
	}
	if service.Cors != nil {
		cors_ARM, err := (*service.Cors).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		cors := *cors_ARM.(*arm.CorsRules)
		result.Properties.Cors = &cors
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (service *StorageAccountsTableService_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.StorageAccountsTableService_Spec{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (service *StorageAccountsTableService_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.StorageAccountsTableService_Spec)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.StorageAccountsTableService_Spec, got %T", armInput)
	}

	// Set property "Cors":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Cors != nil {
			var cors1 CorsRules
			err := cors1.PopulateFromARM(owner, *typedInput.Properties.Cors)
			if err != nil {
				return err
			}
			cors := cors1
			service.Cors = &cors
		}
	}

	// no assignment for property "OperatorSpec"

	// Set property "Owner":
	service.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &StorageAccountsTableService_Spec{}

// ConvertSpecFrom populates our StorageAccountsTableService_Spec from the provided source
func (service *StorageAccountsTableService_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.StorageAccountsTableService_Spec)
	if ok {
		// Populate our instance from source
		return service.AssignProperties_From_StorageAccountsTableService_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.StorageAccountsTableService_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = service.AssignProperties_From_StorageAccountsTableService_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our StorageAccountsTableService_Spec
func (service *StorageAccountsTableService_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.StorageAccountsTableService_Spec)
	if ok {
		// Populate destination from our instance
		return service.AssignProperties_To_StorageAccountsTableService_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.StorageAccountsTableService_Spec{}
	err := service.AssignProperties_To_StorageAccountsTableService_Spec(dst)
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

// AssignProperties_From_StorageAccountsTableService_Spec populates our StorageAccountsTableService_Spec from the provided source StorageAccountsTableService_Spec
func (service *StorageAccountsTableService_Spec) AssignProperties_From_StorageAccountsTableService_Spec(source *storage.StorageAccountsTableService_Spec) error {

	// Cors
	if source.Cors != nil {
		var cor CorsRules
		err := cor.AssignProperties_From_CorsRules(source.Cors)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_CorsRules() to populate field Cors")
		}
		service.Cors = &cor
	} else {
		service.Cors = nil
	}

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec StorageAccountsTableServiceOperatorSpec
		err := operatorSpec.AssignProperties_From_StorageAccountsTableServiceOperatorSpec(source.OperatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_StorageAccountsTableServiceOperatorSpec() to populate field OperatorSpec")
		}
		service.OperatorSpec = &operatorSpec
	} else {
		service.OperatorSpec = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		service.Owner = &owner
	} else {
		service.Owner = nil
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsTableService_Spec populates the provided destination StorageAccountsTableService_Spec from our StorageAccountsTableService_Spec
func (service *StorageAccountsTableService_Spec) AssignProperties_To_StorageAccountsTableService_Spec(destination *storage.StorageAccountsTableService_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Cors
	if service.Cors != nil {
		var cor storage.CorsRules
		err := service.Cors.AssignProperties_To_CorsRules(&cor)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_CorsRules() to populate field Cors")
		}
		destination.Cors = &cor
	} else {
		destination.Cors = nil
	}

	// OperatorSpec
	if service.OperatorSpec != nil {
		var operatorSpec storage.StorageAccountsTableServiceOperatorSpec
		err := service.OperatorSpec.AssignProperties_To_StorageAccountsTableServiceOperatorSpec(&operatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_StorageAccountsTableServiceOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = service.OriginalVersion()

	// Owner
	if service.Owner != nil {
		owner := service.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
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
func (service *StorageAccountsTableService_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type StorageAccountsTableService_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Cors: Specifies CORS rules for the Table service. You can include up to five CorsRule elements in the request. If no
	// CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the
	// Table service.
	Cors *CorsRules_STATUS `json:"cors,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccountsTableService_STATUS{}

// ConvertStatusFrom populates our StorageAccountsTableService_STATUS from the provided source
func (service *StorageAccountsTableService_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.StorageAccountsTableService_STATUS)
	if ok {
		// Populate our instance from source
		return service.AssignProperties_From_StorageAccountsTableService_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.StorageAccountsTableService_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = service.AssignProperties_From_StorageAccountsTableService_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our StorageAccountsTableService_STATUS
func (service *StorageAccountsTableService_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.StorageAccountsTableService_STATUS)
	if ok {
		// Populate destination from our instance
		return service.AssignProperties_To_StorageAccountsTableService_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.StorageAccountsTableService_STATUS{}
	err := service.AssignProperties_To_StorageAccountsTableService_STATUS(dst)
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

var _ genruntime.FromARMConverter = &StorageAccountsTableService_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (service *StorageAccountsTableService_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.StorageAccountsTableService_STATUS{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (service *StorageAccountsTableService_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.StorageAccountsTableService_STATUS)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.StorageAccountsTableService_STATUS, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Cors":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Cors != nil {
			var cors1 CorsRules_STATUS
			err := cors1.PopulateFromARM(owner, *typedInput.Properties.Cors)
			if err != nil {
				return err
			}
			cors := cors1
			service.Cors = &cors
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		service.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		service.Name = &name
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		service.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_StorageAccountsTableService_STATUS populates our StorageAccountsTableService_STATUS from the provided source StorageAccountsTableService_STATUS
func (service *StorageAccountsTableService_STATUS) AssignProperties_From_StorageAccountsTableService_STATUS(source *storage.StorageAccountsTableService_STATUS) error {

	// Conditions
	service.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Cors
	if source.Cors != nil {
		var cor CorsRules_STATUS
		err := cor.AssignProperties_From_CorsRules_STATUS(source.Cors)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_CorsRules_STATUS() to populate field Cors")
		}
		service.Cors = &cor
	} else {
		service.Cors = nil
	}

	// Id
	service.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	service.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	service.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsTableService_STATUS populates the provided destination StorageAccountsTableService_STATUS from our StorageAccountsTableService_STATUS
func (service *StorageAccountsTableService_STATUS) AssignProperties_To_StorageAccountsTableService_STATUS(destination *storage.StorageAccountsTableService_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(service.Conditions)

	// Cors
	if service.Cors != nil {
		var cor storage.CorsRules_STATUS
		err := service.Cors.AssignProperties_To_CorsRules_STATUS(&cor)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_CorsRules_STATUS() to populate field Cors")
		}
		destination.Cors = &cor
	} else {
		destination.Cors = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(service.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(service.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(service.Type)

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
type StorageAccountsTableServiceOperatorSpec struct {
	// ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`

	// SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).
	SecretExpressions []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_StorageAccountsTableServiceOperatorSpec populates our StorageAccountsTableServiceOperatorSpec from the provided source StorageAccountsTableServiceOperatorSpec
func (operator *StorageAccountsTableServiceOperatorSpec) AssignProperties_From_StorageAccountsTableServiceOperatorSpec(source *storage.StorageAccountsTableServiceOperatorSpec) error {

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

// AssignProperties_To_StorageAccountsTableServiceOperatorSpec populates the provided destination StorageAccountsTableServiceOperatorSpec from our StorageAccountsTableServiceOperatorSpec
func (operator *StorageAccountsTableServiceOperatorSpec) AssignProperties_To_StorageAccountsTableServiceOperatorSpec(destination *storage.StorageAccountsTableServiceOperatorSpec) error {
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
	SchemeBuilder.Register(&StorageAccountsTableService{}, &StorageAccountsTableServiceList{})
}
