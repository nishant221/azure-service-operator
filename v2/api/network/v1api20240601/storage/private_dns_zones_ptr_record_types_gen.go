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

// +kubebuilder:rbac:groups=network.azure.com,resources=privatednszonesptrrecords,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={privatednszonesptrrecords/status,privatednszonesptrrecords/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20240601.PrivateDnsZonesPTRRecord
// Generator information:
// - Generated from: /privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/privatedns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/PTR/{relativeRecordSetName}
type PrivateDnsZonesPTRRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDnsZonesPTRRecord_Spec   `json:"spec,omitempty"`
	Status            PrivateDnsZonesPTRRecord_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &PrivateDnsZonesPTRRecord{}

// GetConditions returns the conditions of the resource
func (record *PrivateDnsZonesPTRRecord) GetConditions() conditions.Conditions {
	return record.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (record *PrivateDnsZonesPTRRecord) SetConditions(conditions conditions.Conditions) {
	record.Status.Conditions = conditions
}

var _ configmaps.Exporter = &PrivateDnsZonesPTRRecord{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (record *PrivateDnsZonesPTRRecord) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if record.Spec.OperatorSpec == nil {
		return nil
	}
	return record.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &PrivateDnsZonesPTRRecord{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (record *PrivateDnsZonesPTRRecord) SecretDestinationExpressions() []*core.DestinationExpression {
	if record.Spec.OperatorSpec == nil {
		return nil
	}
	return record.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &PrivateDnsZonesPTRRecord{}

// AzureName returns the Azure name of the resource
func (record *PrivateDnsZonesPTRRecord) AzureName() string {
	return record.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-06-01"
func (record PrivateDnsZonesPTRRecord) GetAPIVersion() string {
	return "2024-06-01"
}

// GetResourceScope returns the scope of the resource
func (record *PrivateDnsZonesPTRRecord) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (record *PrivateDnsZonesPTRRecord) GetSpec() genruntime.ConvertibleSpec {
	return &record.Spec
}

// GetStatus returns the status of this resource
func (record *PrivateDnsZonesPTRRecord) GetStatus() genruntime.ConvertibleStatus {
	return &record.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (record *PrivateDnsZonesPTRRecord) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/privateDnsZones/PTR"
func (record *PrivateDnsZonesPTRRecord) GetType() string {
	return "Microsoft.Network/privateDnsZones/PTR"
}

// NewEmptyStatus returns a new empty (blank) status
func (record *PrivateDnsZonesPTRRecord) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &PrivateDnsZonesPTRRecord_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (record *PrivateDnsZonesPTRRecord) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(record.Spec)
	return record.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (record *PrivateDnsZonesPTRRecord) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*PrivateDnsZonesPTRRecord_STATUS); ok {
		record.Status = *st
		return nil
	}

	// Convert status to required version
	var st PrivateDnsZonesPTRRecord_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	record.Status = st
	return nil
}

// Hub marks that this PrivateDnsZonesPTRRecord is the hub type for conversion
func (record *PrivateDnsZonesPTRRecord) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (record *PrivateDnsZonesPTRRecord) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: record.Spec.OriginalVersion,
		Kind:    "PrivateDnsZonesPTRRecord",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20240601.PrivateDnsZonesPTRRecord
// Generator information:
// - Generated from: /privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/privatedns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/PTR/{relativeRecordSetName}
type PrivateDnsZonesPTRRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDnsZonesPTRRecord `json:"items"`
}

// Storage version of v1api20240601.PrivateDnsZonesPTRRecord_Spec
type PrivateDnsZonesPTRRecord_Spec struct {
	ARecords    []ARecord    `json:"aRecords,omitempty"`
	AaaaRecords []AaaaRecord `json:"aaaaRecords,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                                `json:"azureName,omitempty"`
	CnameRecord     *CnameRecord                          `json:"cnameRecord,omitempty"`
	Etag            *string                               `json:"etag,omitempty"`
	Metadata        map[string]string                     `json:"metadata,omitempty"`
	MxRecords       []MxRecord                            `json:"mxRecords,omitempty"`
	OperatorSpec    *PrivateDnsZonesPTRRecordOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/PrivateDnsZone resource
	Owner       *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"PrivateDnsZone"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PtrRecords  []PtrRecord                        `json:"ptrRecords,omitempty"`
	SoaRecord   *SoaRecord                         `json:"soaRecord,omitempty"`
	SrvRecords  []SrvRecord                        `json:"srvRecords,omitempty"`
	Ttl         *int                               `json:"ttl,omitempty"`
	TxtRecords  []TxtRecord                        `json:"txtRecords,omitempty"`
}

var _ genruntime.ConvertibleSpec = &PrivateDnsZonesPTRRecord_Spec{}

// ConvertSpecFrom populates our PrivateDnsZonesPTRRecord_Spec from the provided source
func (record *PrivateDnsZonesPTRRecord_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == record {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(record)
}

// ConvertSpecTo populates the provided destination from our PrivateDnsZonesPTRRecord_Spec
func (record *PrivateDnsZonesPTRRecord_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == record {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(record)
}

// Storage version of v1api20240601.PrivateDnsZonesPTRRecord_STATUS
type PrivateDnsZonesPTRRecord_STATUS struct {
	ARecords         []ARecord_STATUS       `json:"aRecords,omitempty"`
	AaaaRecords      []AaaaRecord_STATUS    `json:"aaaaRecords,omitempty"`
	CnameRecord      *CnameRecord_STATUS    `json:"cnameRecord,omitempty"`
	Conditions       []conditions.Condition `json:"conditions,omitempty"`
	Etag             *string                `json:"etag,omitempty"`
	Fqdn             *string                `json:"fqdn,omitempty"`
	Id               *string                `json:"id,omitempty"`
	IsAutoRegistered *bool                  `json:"isAutoRegistered,omitempty"`
	Metadata         map[string]string      `json:"metadata,omitempty"`
	MxRecords        []MxRecord_STATUS      `json:"mxRecords,omitempty"`
	Name             *string                `json:"name,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PtrRecords       []PtrRecord_STATUS     `json:"ptrRecords,omitempty"`
	SoaRecord        *SoaRecord_STATUS      `json:"soaRecord,omitempty"`
	SrvRecords       []SrvRecord_STATUS     `json:"srvRecords,omitempty"`
	Ttl              *int                   `json:"ttl,omitempty"`
	TxtRecords       []TxtRecord_STATUS     `json:"txtRecords,omitempty"`
	Type             *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &PrivateDnsZonesPTRRecord_STATUS{}

// ConvertStatusFrom populates our PrivateDnsZonesPTRRecord_STATUS from the provided source
func (record *PrivateDnsZonesPTRRecord_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == record {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(record)
}

// ConvertStatusTo populates the provided destination from our PrivateDnsZonesPTRRecord_STATUS
func (record *PrivateDnsZonesPTRRecord_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == record {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(record)
}

// Storage version of v1api20240601.PrivateDnsZonesPTRRecordOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type PrivateDnsZonesPTRRecordOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&PrivateDnsZonesPTRRecord{}, &PrivateDnsZonesPTRRecordList{})
}
