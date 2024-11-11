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

// +kubebuilder:rbac:groups=network.azure.com,resources=dnsresolversinboundendpoints,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={dnsresolversinboundendpoints/status,dnsresolversinboundendpoints/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220701.DnsResolversInboundEndpoint
// Generator information:
// - Generated from: /dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/dnsresolver.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}/inboundEndpoints/{inboundEndpointName}
type DnsResolversInboundEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsResolversInboundEndpoint_Spec   `json:"spec,omitempty"`
	Status            DnsResolversInboundEndpoint_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DnsResolversInboundEndpoint{}

// GetConditions returns the conditions of the resource
func (endpoint *DnsResolversInboundEndpoint) GetConditions() conditions.Conditions {
	return endpoint.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (endpoint *DnsResolversInboundEndpoint) SetConditions(conditions conditions.Conditions) {
	endpoint.Status.Conditions = conditions
}

var _ configmaps.Exporter = &DnsResolversInboundEndpoint{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (endpoint *DnsResolversInboundEndpoint) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if endpoint.Spec.OperatorSpec == nil {
		return nil
	}
	return endpoint.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &DnsResolversInboundEndpoint{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (endpoint *DnsResolversInboundEndpoint) SecretDestinationExpressions() []*core.DestinationExpression {
	if endpoint.Spec.OperatorSpec == nil {
		return nil
	}
	return endpoint.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &DnsResolversInboundEndpoint{}

// AzureName returns the Azure name of the resource
func (endpoint *DnsResolversInboundEndpoint) AzureName() string {
	return endpoint.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (endpoint DnsResolversInboundEndpoint) GetAPIVersion() string {
	return "2022-07-01"
}

// GetResourceScope returns the scope of the resource
func (endpoint *DnsResolversInboundEndpoint) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (endpoint *DnsResolversInboundEndpoint) GetSpec() genruntime.ConvertibleSpec {
	return &endpoint.Spec
}

// GetStatus returns the status of this resource
func (endpoint *DnsResolversInboundEndpoint) GetStatus() genruntime.ConvertibleStatus {
	return &endpoint.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (endpoint *DnsResolversInboundEndpoint) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsResolvers/inboundEndpoints"
func (endpoint *DnsResolversInboundEndpoint) GetType() string {
	return "Microsoft.Network/dnsResolvers/inboundEndpoints"
}

// NewEmptyStatus returns a new empty (blank) status
func (endpoint *DnsResolversInboundEndpoint) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DnsResolversInboundEndpoint_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (endpoint *DnsResolversInboundEndpoint) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(endpoint.Spec)
	return endpoint.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (endpoint *DnsResolversInboundEndpoint) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DnsResolversInboundEndpoint_STATUS); ok {
		endpoint.Status = *st
		return nil
	}

	// Convert status to required version
	var st DnsResolversInboundEndpoint_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	endpoint.Status = st
	return nil
}

// Hub marks that this DnsResolversInboundEndpoint is the hub type for conversion
func (endpoint *DnsResolversInboundEndpoint) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (endpoint *DnsResolversInboundEndpoint) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: endpoint.Spec.OriginalVersion,
		Kind:    "DnsResolversInboundEndpoint",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220701.DnsResolversInboundEndpoint
// Generator information:
// - Generated from: /dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/dnsresolver.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}/inboundEndpoints/{inboundEndpointName}
type DnsResolversInboundEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DnsResolversInboundEndpoint `json:"items"`
}

// Storage version of v1api20220701.DnsResolversInboundEndpoint_Spec
type DnsResolversInboundEndpoint_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName        string                                   `json:"azureName,omitempty"`
	IpConfigurations []IpConfiguration                        `json:"ipConfigurations,omitempty"`
	Location         *string                                  `json:"location,omitempty"`
	OperatorSpec     *DnsResolversInboundEndpointOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion  string                                   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/DnsResolver resource
	Owner       *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"DnsResolver"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DnsResolversInboundEndpoint_Spec{}

// ConvertSpecFrom populates our DnsResolversInboundEndpoint_Spec from the provided source
func (endpoint *DnsResolversInboundEndpoint_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(endpoint)
}

// ConvertSpecTo populates the provided destination from our DnsResolversInboundEndpoint_Spec
func (endpoint *DnsResolversInboundEndpoint_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(endpoint)
}

// Storage version of v1api20220701.DnsResolversInboundEndpoint_STATUS
type DnsResolversInboundEndpoint_STATUS struct {
	Conditions        []conditions.Condition   `json:"conditions,omitempty"`
	Etag              *string                  `json:"etag,omitempty"`
	Id                *string                  `json:"id,omitempty"`
	IpConfigurations  []IpConfiguration_STATUS `json:"ipConfigurations,omitempty"`
	Location          *string                  `json:"location,omitempty"`
	Name              *string                  `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
	ProvisioningState *string                  `json:"provisioningState,omitempty"`
	ResourceGuid      *string                  `json:"resourceGuid,omitempty"`
	SystemData        *SystemData_STATUS       `json:"systemData,omitempty"`
	Tags              map[string]string        `json:"tags,omitempty"`
	Type              *string                  `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DnsResolversInboundEndpoint_STATUS{}

// ConvertStatusFrom populates our DnsResolversInboundEndpoint_STATUS from the provided source
func (endpoint *DnsResolversInboundEndpoint_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(endpoint)
}

// ConvertStatusTo populates the provided destination from our DnsResolversInboundEndpoint_STATUS
func (endpoint *DnsResolversInboundEndpoint_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(endpoint)
}

// Storage version of v1api20220701.DnsResolversInboundEndpointOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type DnsResolversInboundEndpointOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20220701.IpConfiguration
// IP configuration.
type IpConfiguration struct {
	PrivateIpAddress          *string                `json:"privateIpAddress,omitempty"`
	PrivateIpAllocationMethod *string                `json:"privateIpAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Subnet                    *SubResource           `json:"subnet,omitempty"`
}

// Storage version of v1api20220701.IpConfiguration_STATUS
// IP configuration.
type IpConfiguration_STATUS struct {
	PrivateIpAddress          *string                `json:"privateIpAddress,omitempty"`
	PrivateIpAllocationMethod *string                `json:"privateIpAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Subnet                    *SubResource_STATUS    `json:"subnet,omitempty"`
}

func init() {
	SchemeBuilder.Register(&DnsResolversInboundEndpoint{}, &DnsResolversInboundEndpointList{})
}
