// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=mongodbdatabases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={mongodbdatabases/status,mongodbdatabases/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210515.MongodbDatabase
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases
type MongodbDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsMongodbDatabases_Spec `json:"spec,omitempty"`
	Status            MongoDBDatabaseGetResults_Status      `json:"status,omitempty"`
}

var _ conditions.Conditioner = &MongodbDatabase{}

// GetConditions returns the conditions of the resource
func (database *MongodbDatabase) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *MongodbDatabase) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &MongodbDatabase{}

// AzureName returns the Azure name of the resource
func (database *MongodbDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (database MongodbDatabase) GetAPIVersion() string {
	return "2021-05-15"
}

// GetResourceKind returns the kind of the resource
func (database *MongodbDatabase) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (database *MongodbDatabase) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *MongodbDatabase) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases"
func (database *MongodbDatabase) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *MongodbDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &MongoDBDatabaseGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (database *MongodbDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  database.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (database *MongodbDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*MongoDBDatabaseGetResults_Status); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st MongoDBDatabaseGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// Hub marks that this MongodbDatabase is the hub type for conversion
func (database *MongodbDatabase) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *MongodbDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion,
		Kind:    "MongodbDatabase",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210515.MongodbDatabase
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases
type MongodbDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabase `json:"items"`
}

//Storage version of v1alpha1api20210515.DatabaseAccountsMongodbDatabases_Spec
type DatabaseAccountsMongodbDatabases_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName       string               `json:"azureName"`
	Location        *string              `json:"location,omitempty"`
	Options         *CreateUpdateOptions `json:"options,omitempty"`
	OriginalVersion string               `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner" kind:"DatabaseAccount"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Resource    *MongoDBDatabaseResource          `json:"resource,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsMongodbDatabases_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsMongodbDatabases_Spec from the provided source
func (databases *DatabaseAccountsMongodbDatabases_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == databases {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(databases)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsMongodbDatabases_Spec
func (databases *DatabaseAccountsMongodbDatabases_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == databases {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(databases)
}

//Storage version of v1alpha1api20210515.MongoDBDatabaseGetResults_Status
type MongoDBDatabaseGetResults_Status struct {
	Conditions  []conditions.Condition                        `json:"conditions,omitempty"`
	Id          *string                                       `json:"id,omitempty"`
	Location    *string                                       `json:"location,omitempty"`
	Name        *string                                       `json:"name,omitempty"`
	Options     *OptionsResource_Status                       `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag                        `json:"$propertyBag,omitempty"`
	Resource    *MongoDBDatabaseGetProperties_Status_Resource `json:"resource,omitempty"`
	Tags        map[string]string                             `json:"tags,omitempty"`
	Type        *string                                       `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &MongoDBDatabaseGetResults_Status{}

// ConvertStatusFrom populates our MongoDBDatabaseGetResults_Status from the provided source
func (results *MongoDBDatabaseGetResults_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == results {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(results)
}

// ConvertStatusTo populates the provided destination from our MongoDBDatabaseGetResults_Status
func (results *MongoDBDatabaseGetResults_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == results {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(results)
}

//Storage version of v1alpha1api20210515.CreateUpdateOptions
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CreateUpdateOptions
type CreateUpdateOptions struct {
	AutoscaleSettings *AutoscaleSettings     `json:"autoscaleSettings,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Throughput        *int                   `json:"throughput,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoDBDatabaseGetProperties_Status_Resource
type MongoDBDatabaseGetProperties_Status_Resource struct {
	Etag        *string                `json:"_etag,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid         *string                `json:"_rid,omitempty"`
	Ts          *float64               `json:"_ts,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoDBDatabaseResource
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBDatabaseResource
type MongoDBDatabaseResource struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.OptionsResource_Status
type OptionsResource_Status struct {
	AutoscaleSettings *AutoscaleSettings_Status `json:"autoscaleSettings,omitempty"`
	PropertyBag       genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	Throughput        *int                      `json:"throughput,omitempty"`
}

//Storage version of v1alpha1api20210515.AutoscaleSettings
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoscaleSettings
type AutoscaleSettings struct {
	MaxThroughput *int                   `json:"maxThroughput,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.AutoscaleSettings_Status
type AutoscaleSettings_Status struct {
	MaxThroughput *int                   `json:"maxThroughput,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&MongodbDatabase{}, &MongodbDatabaseList{})
}
