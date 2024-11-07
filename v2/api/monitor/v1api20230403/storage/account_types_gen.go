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

// +kubebuilder:rbac:groups=monitor.azure.com,resources=accounts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=monitor.azure.com,resources={accounts/status,accounts/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230403.Account
// Generator information:
// - Generated from: /monitor/resource-manager/Microsoft.Monitor/stable/2023-04-03/monitoringAccounts_API.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Monitor/accounts/{azureMonitorWorkspaceName}
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Account_Spec   `json:"spec,omitempty"`
	Status            Account_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Account{}

// GetConditions returns the conditions of the resource
func (account *Account) GetConditions() conditions.Conditions {
	return account.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (account *Account) SetConditions(conditions conditions.Conditions) {
	account.Status.Conditions = conditions
}

var _ configmaps.Exporter = &Account{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (account *Account) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if account.Spec.OperatorSpec == nil {
		return nil
	}
	return account.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &Account{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (account *Account) SecretDestinationExpressions() []*core.DestinationExpression {
	if account.Spec.OperatorSpec == nil {
		return nil
	}
	return account.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &Account{}

// AzureName returns the Azure name of the resource
func (account *Account) AzureName() string {
	return account.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-04-03"
func (account Account) GetAPIVersion() string {
	return "2023-04-03"
}

// GetResourceScope returns the scope of the resource
func (account *Account) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (account *Account) GetSpec() genruntime.ConvertibleSpec {
	return &account.Spec
}

// GetStatus returns the status of this resource
func (account *Account) GetStatus() genruntime.ConvertibleStatus {
	return &account.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (account *Account) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Monitor/accounts"
func (account *Account) GetType() string {
	return "Microsoft.Monitor/accounts"
}

// NewEmptyStatus returns a new empty (blank) status
func (account *Account) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Account_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (account *Account) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(account.Spec)
	return account.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (account *Account) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Account_STATUS); ok {
		account.Status = *st
		return nil
	}

	// Convert status to required version
	var st Account_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	account.Status = st
	return nil
}

// Hub marks that this Account is the hub type for conversion
func (account *Account) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (account *Account) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: account.Spec.OriginalVersion,
		Kind:    "Account",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230403.Account
// Generator information:
// - Generated from: /monitor/resource-manager/Microsoft.Monitor/stable/2023-04-03/monitoringAccounts_API.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Monitor/accounts/{azureMonitorWorkspaceName}
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Account `json:"items"`
}

// Storage version of v1api20230403.Account_Spec
type Account_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string               `json:"azureName,omitempty"`
	Location        *string              `json:"location,omitempty"`
	OperatorSpec    *AccountOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner               *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag         genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                            `json:"publicNetworkAccess,omitempty"`
	Tags                map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Account_Spec{}

// ConvertSpecFrom populates our Account_Spec from the provided source
func (account *Account_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(account)
}

// ConvertSpecTo populates the provided destination from our Account_Spec
func (account *Account_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(account)
}

// Storage version of v1api20230403.Account_STATUS
type Account_STATUS struct {
	AccountId                  *string                            `json:"accountId,omitempty"`
	Conditions                 []conditions.Condition             `json:"conditions,omitempty"`
	DefaultIngestionSettings   *IngestionSettings_STATUS          `json:"defaultIngestionSettings,omitempty"`
	Etag                       *string                            `json:"etag,omitempty"`
	Id                         *string                            `json:"id,omitempty"`
	Location                   *string                            `json:"location,omitempty"`
	Metrics                    *Metrics_STATUS                    `json:"metrics,omitempty"`
	Name                       *string                            `json:"name,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                            `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *string                            `json:"publicNetworkAccess,omitempty"`
	SystemData                 *SystemData_STATUS                 `json:"systemData,omitempty"`
	Tags                       map[string]string                  `json:"tags,omitempty"`
	Type                       *string                            `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Account_STATUS{}

// ConvertStatusFrom populates our Account_STATUS from the provided source
func (account *Account_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(account)
}

// ConvertStatusTo populates the provided destination from our Account_STATUS
func (account *Account_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(account)
}

// Storage version of v1api20230403.APIVersion
// +kubebuilder:validation:Enum={"2023-04-03"}
type APIVersion string

const APIVersion_Value = APIVersion("2023-04-03")

// Storage version of v1api20230403.AccountOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type AccountOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20230403.IngestionSettings_STATUS
// Settings for data ingestion
type IngestionSettings_STATUS struct {
	DataCollectionEndpointResourceId *string                `json:"dataCollectionEndpointResourceId,omitempty"`
	DataCollectionRuleResourceId     *string                `json:"dataCollectionRuleResourceId,omitempty"`
	PropertyBag                      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230403.Metrics_STATUS
// Properties related to the metrics container in the Azure Monitor Workspace
type Metrics_STATUS struct {
	InternalId              *string                `json:"internalId,omitempty"`
	PrometheusQueryEndpoint *string                `json:"prometheusQueryEndpoint,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230403.PrivateEndpointConnection_STATUS
// The private endpoint connection resource.
type PrivateEndpointConnection_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230403.SystemData_STATUS
// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Account{}, &AccountList{})
}
