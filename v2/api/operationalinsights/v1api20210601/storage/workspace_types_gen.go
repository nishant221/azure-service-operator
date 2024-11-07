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

// +kubebuilder:rbac:groups=operationalinsights.azure.com,resources=workspaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operationalinsights.azure.com,resources={workspaces/status,workspaces/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20210601.Workspace
// Generator information:
// - Generated from: /operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/Workspaces.json
// - ARM URI: /subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Workspace_Spec   `json:"spec,omitempty"`
	Status            Workspace_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Workspace{}

// GetConditions returns the conditions of the resource
func (workspace *Workspace) GetConditions() conditions.Conditions {
	return workspace.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (workspace *Workspace) SetConditions(conditions conditions.Conditions) {
	workspace.Status.Conditions = conditions
}

var _ configmaps.Exporter = &Workspace{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (workspace *Workspace) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if workspace.Spec.OperatorSpec == nil {
		return nil
	}
	return workspace.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &Workspace{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (workspace *Workspace) SecretDestinationExpressions() []*core.DestinationExpression {
	if workspace.Spec.OperatorSpec == nil {
		return nil
	}
	return workspace.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &Workspace{}

// AzureName returns the Azure name of the resource
func (workspace *Workspace) AzureName() string {
	return workspace.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (workspace Workspace) GetAPIVersion() string {
	return "2021-06-01"
}

// GetResourceScope returns the scope of the resource
func (workspace *Workspace) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (workspace *Workspace) GetSpec() genruntime.ConvertibleSpec {
	return &workspace.Spec
}

// GetStatus returns the status of this resource
func (workspace *Workspace) GetStatus() genruntime.ConvertibleStatus {
	return &workspace.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (workspace *Workspace) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.OperationalInsights/workspaces"
func (workspace *Workspace) GetType() string {
	return "Microsoft.OperationalInsights/workspaces"
}

// NewEmptyStatus returns a new empty (blank) status
func (workspace *Workspace) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Workspace_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (workspace *Workspace) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(workspace.Spec)
	return workspace.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (workspace *Workspace) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Workspace_STATUS); ok {
		workspace.Status = *st
		return nil
	}

	// Convert status to required version
	var st Workspace_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	workspace.Status = st
	return nil
}

// Hub marks that this Workspace is the hub type for conversion
func (workspace *Workspace) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (workspace *Workspace) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: workspace.Spec.OriginalVersion,
		Kind:    "Workspace",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210601.Workspace
// Generator information:
// - Generated from: /operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/Workspaces.json
// - ARM URI: /subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

// Storage version of v1api20210601.APIVersion
// +kubebuilder:validation:Enum={"2021-06-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-06-01")

// Storage version of v1api20210601.Workspace_Spec
type Workspace_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName        string                 `json:"azureName,omitempty"`
	Etag             *string                `json:"etag,omitempty"`
	Features         *WorkspaceFeatures     `json:"features,omitempty"`
	ForceCmkForQuery *bool                  `json:"forceCmkForQuery,omitempty"`
	Location         *string                `json:"location,omitempty"`
	OperatorSpec     *WorkspaceOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion  string                 `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner                           *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag                     genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ProvisioningState               *string                            `json:"provisioningState,omitempty"`
	PublicNetworkAccessForIngestion *string                            `json:"publicNetworkAccessForIngestion,omitempty"`
	PublicNetworkAccessForQuery     *string                            `json:"publicNetworkAccessForQuery,omitempty"`
	RetentionInDays                 *int                               `json:"retentionInDays,omitempty"`
	Sku                             *WorkspaceSku                      `json:"sku,omitempty"`
	Tags                            map[string]string                  `json:"tags,omitempty"`
	WorkspaceCapping                *WorkspaceCapping                  `json:"workspaceCapping,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Workspace_Spec{}

// ConvertSpecFrom populates our Workspace_Spec from the provided source
func (workspace *Workspace_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == workspace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(workspace)
}

// ConvertSpecTo populates the provided destination from our Workspace_Spec
func (workspace *Workspace_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == workspace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(workspace)
}

// Storage version of v1api20210601.Workspace_STATUS
// The top level Workspace resource container.
type Workspace_STATUS struct {
	Conditions                      []conditions.Condition             `json:"conditions,omitempty"`
	CreatedDate                     *string                            `json:"createdDate,omitempty"`
	CustomerId                      *string                            `json:"customerId,omitempty"`
	Etag                            *string                            `json:"etag,omitempty"`
	Features                        *WorkspaceFeatures_STATUS          `json:"features,omitempty"`
	ForceCmkForQuery                *bool                              `json:"forceCmkForQuery,omitempty"`
	Id                              *string                            `json:"id,omitempty"`
	Location                        *string                            `json:"location,omitempty"`
	ModifiedDate                    *string                            `json:"modifiedDate,omitempty"`
	Name                            *string                            `json:"name,omitempty"`
	PrivateLinkScopedResources      []PrivateLinkScopedResource_STATUS `json:"privateLinkScopedResources,omitempty"`
	PropertyBag                     genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ProvisioningState               *string                            `json:"provisioningState,omitempty"`
	PublicNetworkAccessForIngestion *string                            `json:"publicNetworkAccessForIngestion,omitempty"`
	PublicNetworkAccessForQuery     *string                            `json:"publicNetworkAccessForQuery,omitempty"`
	RetentionInDays                 *int                               `json:"retentionInDays,omitempty"`
	Sku                             *WorkspaceSku_STATUS               `json:"sku,omitempty"`
	Tags                            map[string]string                  `json:"tags,omitempty"`
	Type                            *string                            `json:"type,omitempty"`
	WorkspaceCapping                *WorkspaceCapping_STATUS           `json:"workspaceCapping,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Workspace_STATUS{}

// ConvertStatusFrom populates our Workspace_STATUS from the provided source
func (workspace *Workspace_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == workspace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(workspace)
}

// ConvertStatusTo populates the provided destination from our Workspace_STATUS
func (workspace *Workspace_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == workspace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(workspace)
}

// Storage version of v1api20210601.PrivateLinkScopedResource_STATUS
// The private link scope resource reference.
type PrivateLinkScopedResource_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceId  *string                `json:"resourceId,omitempty"`
	ScopeId     *string                `json:"scopeId,omitempty"`
}

// Storage version of v1api20210601.WorkspaceCapping
// The daily volume cap for ingestion.
type WorkspaceCapping struct {
	DailyQuotaGb *float64               `json:"dailyQuotaGb,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.WorkspaceCapping_STATUS
// The daily volume cap for ingestion.
type WorkspaceCapping_STATUS struct {
	DailyQuotaGb        *float64               `json:"dailyQuotaGb,omitempty"`
	DataIngestionStatus *string                `json:"dataIngestionStatus,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	QuotaNextResetTime  *string                `json:"quotaNextResetTime,omitempty"`
}

// Storage version of v1api20210601.WorkspaceFeatures
// Workspace features.
type WorkspaceFeatures struct {
	// ClusterResourceReference: Dedicated LA cluster resourceId that is linked to the workspaces.
	ClusterResourceReference                    *genruntime.ResourceReference `armReference:"ClusterResourceId" json:"clusterResourceReference,omitempty"`
	DisableLocalAuth                            *bool                         `json:"disableLocalAuth,omitempty"`
	EnableDataExport                            *bool                         `json:"enableDataExport,omitempty"`
	EnableLogAccessUsingOnlyResourcePermissions *bool                         `json:"enableLogAccessUsingOnlyResourcePermissions,omitempty"`
	ImmediatePurgeDataOn30Days                  *bool                         `json:"immediatePurgeDataOn30Days,omitempty"`
	PropertyBag                                 genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.WorkspaceFeatures_STATUS
// Workspace features.
type WorkspaceFeatures_STATUS struct {
	ClusterResourceId                           *string                `json:"clusterResourceId,omitempty"`
	DisableLocalAuth                            *bool                  `json:"disableLocalAuth,omitempty"`
	EnableDataExport                            *bool                  `json:"enableDataExport,omitempty"`
	EnableLogAccessUsingOnlyResourcePermissions *bool                  `json:"enableLogAccessUsingOnlyResourcePermissions,omitempty"`
	ImmediatePurgeDataOn30Days                  *bool                  `json:"immediatePurgeDataOn30Days,omitempty"`
	PropertyBag                                 genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.WorkspaceOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type WorkspaceOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20210601.WorkspaceSku
// The SKU (tier) of a workspace.
type WorkspaceSku struct {
	CapacityReservationLevel *int                   `json:"capacityReservationLevel,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.WorkspaceSku_STATUS
// The SKU (tier) of a workspace.
type WorkspaceSku_STATUS struct {
	CapacityReservationLevel *int                   `json:"capacityReservationLevel,omitempty"`
	LastSkuUpdate            *string                `json:"lastSkuUpdate,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
