// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=eventgrid.azure.com,resources=domains,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eventgrid.azure.com,resources={domains/status,domains/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20200601.Domain
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/resourceDefinitions/domains
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Domains_Spec  `json:"spec,omitempty"`
	Status            Domain_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Domain{}

// GetConditions returns the conditions of the resource
func (domain *Domain) GetConditions() conditions.Conditions {
	return domain.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (domain *Domain) SetConditions(conditions conditions.Conditions) {
	domain.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Domain{}

// AzureName returns the Azure name of the resource
func (domain *Domain) AzureName() string {
	return domain.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (domain Domain) GetAPIVersion() string {
	return "2020-06-01"
}

// GetResourceKind returns the kind of the resource
func (domain *Domain) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (domain *Domain) GetSpec() genruntime.ConvertibleSpec {
	return &domain.Spec
}

// GetStatus returns the status of this resource
func (domain *Domain) GetStatus() genruntime.ConvertibleStatus {
	return &domain.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/domains"
func (domain *Domain) GetType() string {
	return "Microsoft.EventGrid/domains"
}

// NewEmptyStatus returns a new empty (blank) status
func (domain *Domain) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Domain_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (domain *Domain) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(domain.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  domain.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (domain *Domain) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Domain_Status); ok {
		domain.Status = *st
		return nil
	}

	// Convert status to required version
	var st Domain_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	domain.Status = st
	return nil
}

// Hub marks that this Domain is the hub type for conversion
func (domain *Domain) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (domain *Domain) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: domain.Spec.OriginalVersion,
		Kind:    "Domain",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20200601.Domain
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/resourceDefinitions/domains
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

//Storage version of v1alpha1api20200601.Domain_Status
type Domain_Status struct {
	Conditions                 []conditions.Condition                                        `json:"conditions,omitempty"`
	Endpoint                   *string                                                       `json:"endpoint,omitempty"`
	Id                         *string                                                       `json:"id,omitempty"`
	InboundIpRules             []InboundIpRule_Status                                        `json:"inboundIpRules,omitempty"`
	InputSchema                *string                                                       `json:"inputSchema,omitempty"`
	InputSchemaMapping         *InputSchemaMapping_Status                                    `json:"inputSchemaMapping,omitempty"`
	Location                   *string                                                       `json:"location,omitempty"`
	MetricResourceId           *string                                                       `json:"metricResourceId,omitempty"`
	Name                       *string                                                       `json:"name,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_Status_Domain_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag                                        `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                                                       `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *string                                                       `json:"publicNetworkAccess,omitempty"`
	SystemData                 *SystemData_Status                                            `json:"systemData,omitempty"`
	Tags                       map[string]string                                             `json:"tags,omitempty"`
	Type                       *string                                                       `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Domain_Status{}

// ConvertStatusFrom populates our Domain_Status from the provided source
func (domain *Domain_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == domain {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(domain)
}

// ConvertStatusTo populates the provided destination from our Domain_Status
func (domain *Domain_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == domain {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(domain)
}

//Storage version of v1alpha1api20200601.Domains_Spec
type Domains_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName          string                  `json:"azureName"`
	InboundIpRules     []InboundIpRule         `json:"inboundIpRules,omitempty"`
	InputSchema        *string                 `json:"inputSchema,omitempty"`
	InputSchemaMapping *JsonInputSchemaMapping `json:"inputSchemaMapping,omitempty"`
	Location           *string                 `json:"location,omitempty"`
	OriginalVersion    string                  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner               genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag         genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                           `json:"publicNetworkAccess,omitempty"`
	Tags                map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Domains_Spec{}

// ConvertSpecFrom populates our Domains_Spec from the provided source
func (domains *Domains_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == domains {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(domains)
}

// ConvertSpecTo populates the provided destination from our Domains_Spec
func (domains *Domains_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == domains {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(domains)
}

//Storage version of v1alpha1api20200601.InboundIpRule
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/InboundIpRule
type InboundIpRule struct {
	Action      *string                `json:"action,omitempty"`
	IpMask      *string                `json:"ipMask,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200601.InboundIpRule_Status
type InboundIpRule_Status struct {
	Action      *string                `json:"action,omitempty"`
	IpMask      *string                `json:"ipMask,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200601.InputSchemaMapping_Status
type InputSchemaMapping_Status struct {
	InputSchemaMappingType *string                `json:"inputSchemaMappingType,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200601.JsonInputSchemaMapping
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/JsonInputSchemaMapping
type JsonInputSchemaMapping struct {
	InputSchemaMappingType *string                           `json:"inputSchemaMappingType,omitempty"`
	Properties             *JsonInputSchemaMappingProperties `json:"properties,omitempty"`
	PropertyBag            genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200601.PrivateEndpointConnection_Status_Domain_SubResourceEmbedded
type PrivateEndpointConnection_Status_Domain_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200601.SystemData_Status
type SystemData_Status struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200601.JsonInputSchemaMappingProperties
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/JsonInputSchemaMappingProperties
type JsonInputSchemaMappingProperties struct {
	DataVersion *JsonFieldWithDefault  `json:"dataVersion,omitempty"`
	EventTime   *JsonField             `json:"eventTime,omitempty"`
	EventType   *JsonFieldWithDefault  `json:"eventType,omitempty"`
	Id          *JsonField             `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Subject     *JsonFieldWithDefault  `json:"subject,omitempty"`
	Topic       *JsonField             `json:"topic,omitempty"`
}

//Storage version of v1alpha1api20200601.JsonField
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/JsonField
type JsonField struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SourceField *string                `json:"sourceField,omitempty"`
}

//Storage version of v1alpha1api20200601.JsonFieldWithDefault
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/JsonFieldWithDefault
type JsonFieldWithDefault struct {
	DefaultValue *string                `json:"defaultValue,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SourceField  *string                `json:"sourceField,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Domain{}, &DomainList{})
}
