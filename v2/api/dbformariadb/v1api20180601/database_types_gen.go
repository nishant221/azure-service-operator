// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180601

import (
	"fmt"
	arm "github.com/Azure/azure-service-operator/v2/api/dbformariadb/v1api20180601/arm"
	storage "github.com/Azure/azure-service-operator/v2/api/dbformariadb/v1api20180601/storage"
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
// - Generated from: /mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/mariadb.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/databases/{databaseName}
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Database_Spec   `json:"spec,omitempty"`
	Status            Database_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Database{}

// GetConditions returns the conditions of the resource
func (database *Database) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *Database) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ conversion.Convertible = &Database{}

// ConvertFrom populates our Database from the provided hub Database
func (database *Database) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.Database)
	if !ok {
		return fmt.Errorf("expected dbformariadb/v1api20180601/storage/Database but received %T instead", hub)
	}

	return database.AssignProperties_From_Database(source)
}

// ConvertTo populates the provided hub Database from our Database
func (database *Database) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.Database)
	if !ok {
		return fmt.Errorf("expected dbformariadb/v1api20180601/storage/Database but received %T instead", hub)
	}

	return database.AssignProperties_To_Database(destination)
}

// +kubebuilder:webhook:path=/mutate-dbformariadb-azure-com-v1api20180601-database,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformariadb.azure.com,resources=databases,verbs=create;update,versions=v1api20180601,name=default.v1api20180601.databases.dbformariadb.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &Database{}

// Default applies defaults to the Database resource
func (database *Database) Default() {
	database.defaultImpl()
	var temp any = database
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (database *Database) defaultAzureName() {
	if database.Spec.AzureName == "" {
		database.Spec.AzureName = database.Name
	}
}

// defaultImpl applies the code generated defaults to the Database resource
func (database *Database) defaultImpl() { database.defaultAzureName() }

var _ configmaps.Exporter = &Database{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (database *Database) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if database.Spec.OperatorSpec == nil {
		return nil
	}
	return database.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &Database{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (database *Database) SecretDestinationExpressions() []*core.DestinationExpression {
	if database.Spec.OperatorSpec == nil {
		return nil
	}
	return database.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.ImportableResource = &Database{}

// InitializeSpec initializes the spec for this resource from the given status
func (database *Database) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Database_STATUS); ok {
		return database.Spec.Initialize_From_Database_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Database_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &Database{}

// AzureName returns the Azure name of the resource
func (database *Database) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-06-01"
func (database Database) GetAPIVersion() string {
	return "2018-06-01"
}

// GetResourceScope returns the scope of the resource
func (database *Database) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (database *Database) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *Database) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (database *Database) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMariaDB/servers/databases"
func (database *Database) GetType() string {
	return "Microsoft.DBforMariaDB/servers/databases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *Database) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Database_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (database *Database) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return database.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (database *Database) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Database_STATUS); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st Database_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-dbformariadb-azure-com-v1api20180601-database,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformariadb.azure.com,resources=databases,verbs=create;update,versions=v1api20180601,name=validate.v1api20180601.databases.dbformariadb.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &Database{}

// ValidateCreate validates the creation of the resource
func (database *Database) ValidateCreate() (admission.Warnings, error) {
	validations := database.createValidations()
	var temp any = database
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (database *Database) ValidateDelete() (admission.Warnings, error) {
	validations := database.deleteValidations()
	var temp any = database
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (database *Database) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := database.updateValidations()
	var temp any = database
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (database *Database) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){database.validateResourceReferences, database.validateOwnerReference, database.validateSecretDestinations, database.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (database *Database) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (database *Database) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return database.validateResourceReferences()
		},
		database.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return database.validateOwnerReference()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return database.validateSecretDestinations()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return database.validateConfigMapDestinations()
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (database *Database) validateConfigMapDestinations() (admission.Warnings, error) {
	if database.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(database, nil, database.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (database *Database) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(database)
}

// validateResourceReferences validates all resource references
func (database *Database) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&database.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (database *Database) validateSecretDestinations() (admission.Warnings, error) {
	if database.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(database, nil, database.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (database *Database) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*Database)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, database)
}

// AssignProperties_From_Database populates our Database from the provided source Database
func (database *Database) AssignProperties_From_Database(source *storage.Database) error {

	// ObjectMeta
	database.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Database_Spec
	err := spec.AssignProperties_From_Database_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Database_Spec() to populate field Spec")
	}
	database.Spec = spec

	// Status
	var status Database_STATUS
	err = status.AssignProperties_From_Database_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Database_STATUS() to populate field Status")
	}
	database.Status = status

	// No error
	return nil
}

// AssignProperties_To_Database populates the provided destination Database from our Database
func (database *Database) AssignProperties_To_Database(destination *storage.Database) error {

	// ObjectMeta
	destination.ObjectMeta = *database.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.Database_Spec
	err := database.Spec.AssignProperties_To_Database_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Database_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.Database_STATUS
	err = database.Status.AssignProperties_To_Database_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Database_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *Database) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion(),
		Kind:    "Database",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/mariadb.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/databases/{databaseName}
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

type Database_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// Charset: The charset of the database.
	Charset *string `json:"charset,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`

	// OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
	// passed directly to Azure
	OperatorSpec *DatabaseOperatorSpec `json:"operatorSpec,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbformariadb.azure.com/Server resource
	Owner *genruntime.KnownResourceReference `group:"dbformariadb.azure.com" json:"owner,omitempty" kind:"Server"`
}

var _ genruntime.ARMTransformer = &Database_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (database *Database_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if database == nil {
		return nil, nil
	}
	result := &arm.Database_Spec{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if database.Charset != nil || database.Collation != nil {
		result.Properties = &arm.DatabaseProperties{}
	}
	if database.Charset != nil {
		charset := *database.Charset
		result.Properties.Charset = &charset
	}
	if database.Collation != nil {
		collation := *database.Collation
		result.Properties.Collation = &collation
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (database *Database_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.Database_Spec{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (database *Database_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.Database_Spec)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.Database_Spec, got %T", armInput)
	}

	// Set property "AzureName":
	database.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "Charset":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Charset != nil {
			charset := *typedInput.Properties.Charset
			database.Charset = &charset
		}
	}

	// Set property "Collation":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Collation != nil {
			collation := *typedInput.Properties.Collation
			database.Collation = &collation
		}
	}

	// no assignment for property "OperatorSpec"

	// Set property "Owner":
	database.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Database_Spec{}

// ConvertSpecFrom populates our Database_Spec from the provided source
func (database *Database_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.Database_Spec)
	if ok {
		// Populate our instance from source
		return database.AssignProperties_From_Database_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.Database_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = database.AssignProperties_From_Database_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Database_Spec
func (database *Database_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.Database_Spec)
	if ok {
		// Populate destination from our instance
		return database.AssignProperties_To_Database_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Database_Spec{}
	err := database.AssignProperties_To_Database_Spec(dst)
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

// AssignProperties_From_Database_Spec populates our Database_Spec from the provided source Database_Spec
func (database *Database_Spec) AssignProperties_From_Database_Spec(source *storage.Database_Spec) error {

	// AzureName
	database.AzureName = source.AzureName

	// Charset
	database.Charset = genruntime.ClonePointerToString(source.Charset)

	// Collation
	database.Collation = genruntime.ClonePointerToString(source.Collation)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec DatabaseOperatorSpec
		err := operatorSpec.AssignProperties_From_DatabaseOperatorSpec(source.OperatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_DatabaseOperatorSpec() to populate field OperatorSpec")
		}
		database.OperatorSpec = &operatorSpec
	} else {
		database.OperatorSpec = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		database.Owner = &owner
	} else {
		database.Owner = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Database_Spec populates the provided destination Database_Spec from our Database_Spec
func (database *Database_Spec) AssignProperties_To_Database_Spec(destination *storage.Database_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = database.AzureName

	// Charset
	destination.Charset = genruntime.ClonePointerToString(database.Charset)

	// Collation
	destination.Collation = genruntime.ClonePointerToString(database.Collation)

	// OperatorSpec
	if database.OperatorSpec != nil {
		var operatorSpec storage.DatabaseOperatorSpec
		err := database.OperatorSpec.AssignProperties_To_DatabaseOperatorSpec(&operatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_DatabaseOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = database.OriginalVersion()

	// Owner
	if database.Owner != nil {
		owner := database.Owner.Copy()
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

// Initialize_From_Database_STATUS populates our Database_Spec from the provided source Database_STATUS
func (database *Database_Spec) Initialize_From_Database_STATUS(source *Database_STATUS) error {

	// Charset
	database.Charset = genruntime.ClonePointerToString(source.Charset)

	// Collation
	database.Collation = genruntime.ClonePointerToString(source.Collation)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (database *Database_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (database *Database_Spec) SetAzureName(azureName string) { database.AzureName = azureName }

type Database_STATUS struct {
	// Charset: The charset of the database.
	Charset *string `json:"charset,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Database_STATUS{}

// ConvertStatusFrom populates our Database_STATUS from the provided source
func (database *Database_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.Database_STATUS)
	if ok {
		// Populate our instance from source
		return database.AssignProperties_From_Database_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.Database_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = database.AssignProperties_From_Database_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Database_STATUS
func (database *Database_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.Database_STATUS)
	if ok {
		// Populate destination from our instance
		return database.AssignProperties_To_Database_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Database_STATUS{}
	err := database.AssignProperties_To_Database_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Database_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (database *Database_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.Database_STATUS{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (database *Database_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.Database_STATUS)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.Database_STATUS, got %T", armInput)
	}

	// Set property "Charset":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Charset != nil {
			charset := *typedInput.Properties.Charset
			database.Charset = &charset
		}
	}

	// Set property "Collation":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Collation != nil {
			collation := *typedInput.Properties.Collation
			database.Collation = &collation
		}
	}

	// no assignment for property "Conditions"

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		database.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		database.Name = &name
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		database.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Database_STATUS populates our Database_STATUS from the provided source Database_STATUS
func (database *Database_STATUS) AssignProperties_From_Database_STATUS(source *storage.Database_STATUS) error {

	// Charset
	database.Charset = genruntime.ClonePointerToString(source.Charset)

	// Collation
	database.Collation = genruntime.ClonePointerToString(source.Collation)

	// Conditions
	database.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	database.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	database.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	database.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Database_STATUS populates the provided destination Database_STATUS from our Database_STATUS
func (database *Database_STATUS) AssignProperties_To_Database_STATUS(destination *storage.Database_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Charset
	destination.Charset = genruntime.ClonePointerToString(database.Charset)

	// Collation
	destination.Collation = genruntime.ClonePointerToString(database.Collation)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(database.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(database.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(database.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(database.Type)

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
type DatabaseOperatorSpec struct {
	// ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`

	// SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).
	SecretExpressions []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_DatabaseOperatorSpec populates our DatabaseOperatorSpec from the provided source DatabaseOperatorSpec
func (operator *DatabaseOperatorSpec) AssignProperties_From_DatabaseOperatorSpec(source *storage.DatabaseOperatorSpec) error {

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

// AssignProperties_To_DatabaseOperatorSpec populates the provided destination DatabaseOperatorSpec from our DatabaseOperatorSpec
func (operator *DatabaseOperatorSpec) AssignProperties_To_DatabaseOperatorSpec(destination *storage.DatabaseOperatorSpec) error {
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
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}
