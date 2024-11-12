// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20240601/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=bastionhosts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={bastionhosts/status,bastionhosts/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20240301.BastionHost
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2024-03-01/bastionHost.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}
type BastionHost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BastionHost_Spec   `json:"spec,omitempty"`
	Status            BastionHost_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &BastionHost{}

// GetConditions returns the conditions of the resource
func (host *BastionHost) GetConditions() conditions.Conditions {
	return host.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (host *BastionHost) SetConditions(conditions conditions.Conditions) {
	host.Status.Conditions = conditions
}

var _ configmaps.Exporter = &BastionHost{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (host *BastionHost) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if host.Spec.OperatorSpec == nil {
		return nil
	}
	return host.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &BastionHost{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (host *BastionHost) SecretDestinationExpressions() []*core.DestinationExpression {
	if host.Spec.OperatorSpec == nil {
		return nil
	}
	return host.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &BastionHost{}

// AzureName returns the Azure name of the resource
func (host *BastionHost) AzureName() string {
	return host.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-03-01"
func (host BastionHost) GetAPIVersion() string {
	return "2024-03-01"
}

// GetResourceScope returns the scope of the resource
func (host *BastionHost) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (host *BastionHost) GetSpec() genruntime.ConvertibleSpec {
	return &host.Spec
}

// GetStatus returns the status of this resource
func (host *BastionHost) GetStatus() genruntime.ConvertibleStatus {
	return &host.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (host *BastionHost) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/bastionHosts"
func (host *BastionHost) GetType() string {
	return "Microsoft.Network/bastionHosts"
}

// NewEmptyStatus returns a new empty (blank) status
func (host *BastionHost) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &BastionHost_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (host *BastionHost) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(host.Spec)
	return host.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (host *BastionHost) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*BastionHost_STATUS); ok {
		host.Status = *st
		return nil
	}

	// Convert status to required version
	var st BastionHost_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	host.Status = st
	return nil
}

// Hub marks that this BastionHost is the hub type for conversion
func (host *BastionHost) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (host *BastionHost) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: host.Spec.OriginalVersion,
		Kind:    "BastionHost",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20240301.BastionHost
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2024-03-01/bastionHost.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}
type BastionHostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BastionHost `json:"items"`
}

// Storage version of v1api20240301.APIVersion
// +kubebuilder:validation:Enum={"2024-03-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2024-03-01")

// Storage version of v1api20240301.BastionHost_Spec
type BastionHost_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName              string                                   `json:"azureName,omitempty"`
	DisableCopyPaste       *bool                                    `json:"disableCopyPaste,omitempty"`
	DnsName                *string                                  `json:"dnsName,omitempty"`
	EnableFileCopy         *bool                                    `json:"enableFileCopy,omitempty"`
	EnableIpConnect        *bool                                    `json:"enableIpConnect,omitempty"`
	EnableKerberos         *bool                                    `json:"enableKerberos,omitempty"`
	EnableSessionRecording *bool                                    `json:"enableSessionRecording,omitempty"`
	EnableShareableLink    *bool                                    `json:"enableShareableLink,omitempty"`
	EnableTunneling        *bool                                    `json:"enableTunneling,omitempty"`
	IpConfigurations       []BastionHostIPConfiguration             `json:"ipConfigurations,omitempty"`
	Location               *string                                  `json:"location,omitempty"`
	NetworkAcls            *BastionHostPropertiesFormat_NetworkAcls `json:"networkAcls,omitempty"`
	OperatorSpec           *BastionHostOperatorSpec                 `json:"operatorSpec,omitempty"`
	OriginalVersion        string                                   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner          *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag    genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ScaleUnits     *int                               `json:"scaleUnits,omitempty"`
	Sku            *Sku                               `json:"sku,omitempty"`
	Tags           map[string]string                  `json:"tags,omitempty"`
	VirtualNetwork *SubResource                       `json:"virtualNetwork,omitempty"`
	Zones          []string                           `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleSpec = &BastionHost_Spec{}

// ConvertSpecFrom populates our BastionHost_Spec from the provided source
func (host *BastionHost_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == host {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(host)
}

// ConvertSpecTo populates the provided destination from our BastionHost_Spec
func (host *BastionHost_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == host {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(host)
}

// Storage version of v1api20240301.BastionHost_STATUS
// Bastion Host resource.
type BastionHost_STATUS struct {
	Conditions             []conditions.Condition                          `json:"conditions,omitempty"`
	DisableCopyPaste       *bool                                           `json:"disableCopyPaste,omitempty"`
	DnsName                *string                                         `json:"dnsName,omitempty"`
	EnableFileCopy         *bool                                           `json:"enableFileCopy,omitempty"`
	EnableIpConnect        *bool                                           `json:"enableIpConnect,omitempty"`
	EnableKerberos         *bool                                           `json:"enableKerberos,omitempty"`
	EnableSessionRecording *bool                                           `json:"enableSessionRecording,omitempty"`
	EnableShareableLink    *bool                                           `json:"enableShareableLink,omitempty"`
	EnableTunneling        *bool                                           `json:"enableTunneling,omitempty"`
	Etag                   *string                                         `json:"etag,omitempty"`
	Id                     *string                                         `json:"id,omitempty"`
	IpConfigurations       []BastionHostIPConfiguration_STATUS             `json:"ipConfigurations,omitempty"`
	Location               *string                                         `json:"location,omitempty"`
	Name                   *string                                         `json:"name,omitempty"`
	NetworkAcls            *BastionHostPropertiesFormat_NetworkAcls_STATUS `json:"networkAcls,omitempty"`
	PropertyBag            genruntime.PropertyBag                          `json:"$propertyBag,omitempty"`
	ProvisioningState      *string                                         `json:"provisioningState,omitempty"`
	ScaleUnits             *int                                            `json:"scaleUnits,omitempty"`
	Sku                    *Sku_STATUS                                     `json:"sku,omitempty"`
	Tags                   map[string]string                               `json:"tags,omitempty"`
	Type                   *string                                         `json:"type,omitempty"`
	VirtualNetwork         *SubResource_STATUS                             `json:"virtualNetwork,omitempty"`
	Zones                  []string                                        `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleStatus = &BastionHost_STATUS{}

// ConvertStatusFrom populates our BastionHost_STATUS from the provided source
func (host *BastionHost_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == host {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(host)
}

// ConvertStatusTo populates the provided destination from our BastionHost_STATUS
func (host *BastionHost_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == host {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(host)
}

// Storage version of v1api20240301.BastionHostIPConfiguration
// IP configuration of an Bastion Host.
type BastionHostIPConfiguration struct {
	Name                      *string                `json:"name,omitempty"`
	PrivateIPAllocationMethod *string                `json:"privateIPAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicIPAddress           *SubResource           `json:"publicIPAddress,omitempty"`
	Subnet                    *SubResource           `json:"subnet,omitempty"`
}

// Storage version of v1api20240301.BastionHostIPConfiguration_STATUS
// IP configuration of an Bastion Host.
type BastionHostIPConfiguration_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240301.BastionHostOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type BastionHostOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20240301.BastionHostPropertiesFormat_NetworkAcls
type BastionHostPropertiesFormat_NetworkAcls struct {
	IpRules     []IPRule               `json:"ipRules,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240301.BastionHostPropertiesFormat_NetworkAcls_STATUS
type BastionHostPropertiesFormat_NetworkAcls_STATUS struct {
	IpRules     []IPRule_STATUS        `json:"ipRules,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240301.Sku
// The sku of this Bastion Host.
type Sku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240301.Sku_STATUS
// The sku of this Bastion Host.
type Sku_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240301.SubResource
// Reference to another subresource.
type SubResource struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// AssignProperties_From_SubResource populates our SubResource from the provided source SubResource
func (resource *SubResource) AssignProperties_From_SubResource(source *storage.SubResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Reference
	if source.Reference != nil {
		reference := source.Reference.Copy()
		resource.Reference = &reference
	} else {
		resource.Reference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// Invoke the augmentConversionForSubResource interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSubResource); ok {
		err := augmentedResource.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SubResource populates the provided destination SubResource from our SubResource
func (resource *SubResource) AssignProperties_To_SubResource(destination *storage.SubResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Reference
	if resource.Reference != nil {
		reference := resource.Reference.Copy()
		destination.Reference = &reference
	} else {
		destination.Reference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSubResource interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSubResource); ok {
		err := augmentedResource.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20240301.SubResource_STATUS
// Reference to another subresource.
type SubResource_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_SubResource_STATUS populates our SubResource_STATUS from the provided source SubResource_STATUS
func (resource *SubResource_STATUS) AssignProperties_From_SubResource_STATUS(source *storage.SubResource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// Invoke the augmentConversionForSubResource_STATUS interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSubResource_STATUS); ok {
		err := augmentedResource.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SubResource_STATUS populates the provided destination SubResource_STATUS from our SubResource_STATUS
func (resource *SubResource_STATUS) AssignProperties_To_SubResource_STATUS(destination *storage.SubResource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSubResource_STATUS interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSubResource_STATUS); ok {
		err := augmentedResource.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForSubResource interface {
	AssignPropertiesFrom(src *storage.SubResource) error
	AssignPropertiesTo(dst *storage.SubResource) error
}

type augmentConversionForSubResource_STATUS interface {
	AssignPropertiesFrom(src *storage.SubResource_STATUS) error
	AssignPropertiesTo(dst *storage.SubResource_STATUS) error
}

// Storage version of v1api20240301.IPRule
type IPRule struct {
	AddressPrefix *string                `json:"addressPrefix,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240301.IPRule_STATUS
type IPRule_STATUS struct {
	AddressPrefix *string                `json:"addressPrefix,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&BastionHost{}, &BastionHostList{})
}