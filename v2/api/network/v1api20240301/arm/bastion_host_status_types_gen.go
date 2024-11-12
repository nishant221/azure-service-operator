// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

// Bastion Host resource.
type BastionHost_STATUS struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Represents the bastion host resource.
	Properties *BastionHostPropertiesFormat_STATUS `json:"properties,omitempty"`

	// Sku: The sku of this Bastion Host.
	Sku *Sku_STATUS `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`

	// Zones: A list of availability zones denoting where the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

// Properties of the Bastion Host.
type BastionHostPropertiesFormat_STATUS struct {
	// DisableCopyPaste: Enable/Disable Copy/Paste feature of the Bastion Host resource.
	DisableCopyPaste *bool `json:"disableCopyPaste,omitempty"`

	// DnsName: FQDN for the endpoint on which bastion host is accessible.
	DnsName *string `json:"dnsName,omitempty"`

	// EnableFileCopy: Enable/Disable File Copy feature of the Bastion Host resource.
	EnableFileCopy *bool `json:"enableFileCopy,omitempty"`

	// EnableIpConnect: Enable/Disable IP Connect feature of the Bastion Host resource.
	EnableIpConnect *bool `json:"enableIpConnect,omitempty"`

	// EnableKerberos: Enable/Disable Kerberos feature of the Bastion Host resource.
	EnableKerberos *bool `json:"enableKerberos,omitempty"`

	// EnableSessionRecording: Enable/Disable Session Recording feature of the Bastion Host resource.
	EnableSessionRecording *bool `json:"enableSessionRecording,omitempty"`

	// EnableShareableLink: Enable/Disable Shareable Link of the Bastion Host resource.
	EnableShareableLink *bool `json:"enableShareableLink,omitempty"`

	// EnableTunneling: Enable/Disable Tunneling feature of the Bastion Host resource.
	EnableTunneling *bool `json:"enableTunneling,omitempty"`

	// IpConfigurations: IP configuration of the Bastion Host resource.
	IpConfigurations []BastionHostIPConfiguration_STATUS             `json:"ipConfigurations,omitempty"`
	NetworkAcls      *BastionHostPropertiesFormat_NetworkAcls_STATUS `json:"networkAcls,omitempty"`

	// ProvisioningState: The provisioning state of the bastion host resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ScaleUnits: The scale units for the Bastion Host resource.
	ScaleUnits *int `json:"scaleUnits,omitempty"`

	// VirtualNetwork: Reference to an existing virtual network required for Developer Bastion Host only.
	VirtualNetwork *SubResource_STATUS `json:"virtualNetwork,omitempty"`
}

// The sku of this Bastion Host.
type Sku_STATUS struct {
	// Name: The name of the sku of this Bastion Host.
	Name *Sku_Name_STATUS `json:"name,omitempty"`
}

// IP configuration of an Bastion Host.
type BastionHostIPConfiguration_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type BastionHostPropertiesFormat_NetworkAcls_STATUS struct {
	// IpRules: Sets the IP ACL rules for Developer Bastion Host.
	IpRules []IPRule_STATUS `json:"ipRules,omitempty"`
}

// The current provisioning state.
type ProvisioningState_STATUS string

const (
	ProvisioningState_STATUS_Deleting  = ProvisioningState_STATUS("Deleting")
	ProvisioningState_STATUS_Failed    = ProvisioningState_STATUS("Failed")
	ProvisioningState_STATUS_Succeeded = ProvisioningState_STATUS("Succeeded")
	ProvisioningState_STATUS_Updating  = ProvisioningState_STATUS("Updating")
)

// Mapping from string to ProvisioningState_STATUS
var provisioningState_STATUS_Values = map[string]ProvisioningState_STATUS{
	"deleting":  ProvisioningState_STATUS_Deleting,
	"failed":    ProvisioningState_STATUS_Failed,
	"succeeded": ProvisioningState_STATUS_Succeeded,
	"updating":  ProvisioningState_STATUS_Updating,
}

type Sku_Name_STATUS string

const (
	Sku_Name_STATUS_Basic     = Sku_Name_STATUS("Basic")
	Sku_Name_STATUS_Developer = Sku_Name_STATUS("Developer")
	Sku_Name_STATUS_Premium   = Sku_Name_STATUS("Premium")
	Sku_Name_STATUS_Standard  = Sku_Name_STATUS("Standard")
)

// Mapping from string to Sku_Name_STATUS
var sku_Name_STATUS_Values = map[string]Sku_Name_STATUS{
	"basic":     Sku_Name_STATUS_Basic,
	"developer": Sku_Name_STATUS_Developer,
	"premium":   Sku_Name_STATUS_Premium,
	"standard":  Sku_Name_STATUS_Standard,
}

// Reference to another subresource.
type SubResource_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type IPRule_STATUS struct {
	// AddressPrefix: Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
	AddressPrefix *string `json:"addressPrefix,omitempty"`
}
