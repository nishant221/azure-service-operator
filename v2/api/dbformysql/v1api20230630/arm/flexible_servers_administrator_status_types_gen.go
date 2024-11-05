// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type FlexibleServersAdministrator_STATUS struct {
	// Id: Fully qualified resource ID for the resource. E.g.
	// "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The properties of an administrator.
	Properties *AdministratorProperties_STATUS `json:"properties,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// The properties of an administrator.
type AdministratorProperties_STATUS struct {
	// AdministratorType: Type of the sever administrator.
	AdministratorType *AdministratorProperties_AdministratorType_STATUS `json:"administratorType,omitempty"`

	// IdentityResourceId: The resource id of the identity used for AAD Authentication.
	IdentityResourceId *string `json:"identityResourceId,omitempty"`

	// Login: Login name of the server administrator.
	Login *string `json:"login,omitempty"`

	// Sid: SID (object ID) of the server administrator.
	Sid *string `json:"sid,omitempty"`

	// TenantId: Tenant ID of the administrator.
	TenantId *string `json:"tenantId,omitempty"`
}

type AdministratorProperties_AdministratorType_STATUS string

const AdministratorProperties_AdministratorType_STATUS_ActiveDirectory = AdministratorProperties_AdministratorType_STATUS("ActiveDirectory")

// Mapping from string to AdministratorProperties_AdministratorType_STATUS
var administratorProperties_AdministratorType_STATUS_Values = map[string]AdministratorProperties_AdministratorType_STATUS{
	"activedirectory": AdministratorProperties_AdministratorType_STATUS_ActiveDirectory,
}