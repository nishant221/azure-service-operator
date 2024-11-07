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

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=sqldatabasecontainerthroughputsettings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={sqldatabasecontainerthroughputsettings/status,sqldatabasecontainerthroughputsettings/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20231115.SqlDatabaseContainerThroughputSetting
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-11-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/throughputSettings/default
type SqlDatabaseContainerThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlDatabaseContainerThroughputSetting_Spec   `json:"spec,omitempty"`
	Status            SqlDatabaseContainerThroughputSetting_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerThroughputSetting{}

// GetConditions returns the conditions of the resource
func (setting *SqlDatabaseContainerThroughputSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *SqlDatabaseContainerThroughputSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ configmaps.Exporter = &SqlDatabaseContainerThroughputSetting{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (setting *SqlDatabaseContainerThroughputSetting) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if setting.Spec.OperatorSpec == nil {
		return nil
	}
	return setting.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &SqlDatabaseContainerThroughputSetting{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (setting *SqlDatabaseContainerThroughputSetting) SecretDestinationExpressions() []*core.DestinationExpression {
	if setting.Spec.OperatorSpec == nil {
		return nil
	}
	return setting.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainerThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (setting *SqlDatabaseContainerThroughputSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-11-15"
func (setting SqlDatabaseContainerThroughputSetting) GetAPIVersion() string {
	return "2023-11-15"
}

// GetResourceScope returns the scope of the resource
func (setting *SqlDatabaseContainerThroughputSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *SqlDatabaseContainerThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *SqlDatabaseContainerThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (setting *SqlDatabaseContainerThroughputSetting) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/throughputSettings"
func (setting *SqlDatabaseContainerThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *SqlDatabaseContainerThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlDatabaseContainerThroughputSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (setting *SqlDatabaseContainerThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return setting.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (setting *SqlDatabaseContainerThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlDatabaseContainerThroughputSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlDatabaseContainerThroughputSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// Hub marks that this SqlDatabaseContainerThroughputSetting is the hub type for conversion
func (setting *SqlDatabaseContainerThroughputSetting) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *SqlDatabaseContainerThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion,
		Kind:    "SqlDatabaseContainerThroughputSetting",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20231115.SqlDatabaseContainerThroughputSetting
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-11-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/throughputSettings/default
type SqlDatabaseContainerThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerThroughputSetting `json:"items"`
}

// Storage version of v1api20231115.SqlDatabaseContainerThroughputSetting_Spec
type SqlDatabaseContainerThroughputSetting_Spec struct {
	Location        *string                                            `json:"location,omitempty"`
	OperatorSpec    *SqlDatabaseContainerThroughputSettingOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                             `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabaseContainer resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabaseContainer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsResource        `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &SqlDatabaseContainerThroughputSetting_Spec{}

// ConvertSpecFrom populates our SqlDatabaseContainerThroughputSetting_Spec from the provided source
func (setting *SqlDatabaseContainerThroughputSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(setting)
}

// ConvertSpecTo populates the provided destination from our SqlDatabaseContainerThroughputSetting_Spec
func (setting *SqlDatabaseContainerThroughputSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(setting)
}

// Storage version of v1api20231115.SqlDatabaseContainerThroughputSetting_STATUS
type SqlDatabaseContainerThroughputSetting_STATUS struct {
	Conditions  []conditions.Condition                           `json:"conditions,omitempty"`
	Id          *string                                          `json:"id,omitempty"`
	Location    *string                                          `json:"location,omitempty"`
	Name        *string                                          `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                           `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                                `json:"tags,omitempty"`
	Type        *string                                          `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlDatabaseContainerThroughputSetting_STATUS{}

// ConvertStatusFrom populates our SqlDatabaseContainerThroughputSetting_STATUS from the provided source
func (setting *SqlDatabaseContainerThroughputSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(setting)
}

// ConvertStatusTo populates the provided destination from our SqlDatabaseContainerThroughputSetting_STATUS
func (setting *SqlDatabaseContainerThroughputSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(setting)
}

// Storage version of v1api20231115.SqlDatabaseContainerThroughputSettingOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type SqlDatabaseContainerThroughputSettingOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerThroughputSetting{}, &SqlDatabaseContainerThroughputSettingList{})
}
