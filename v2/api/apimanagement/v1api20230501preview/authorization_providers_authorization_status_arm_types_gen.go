// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501preview

type AuthorizationProvidersAuthorization_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the Authorization Contract.
	Properties *AuthorizationContractProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Authorization details.
type AuthorizationContractProperties_STATUS_ARM struct {
	// AuthorizationType: Authorization type options
	AuthorizationType *AuthorizationContractProperties_AuthorizationType_STATUS_ARM `json:"authorizationType,omitempty"`

	// Error: Authorization error details.
	Error *AuthorizationError_STATUS_ARM `json:"error,omitempty"`

	// Oauth2GrantType: OAuth2 grant type options
	Oauth2GrantType *AuthorizationContractProperties_Oauth2GrantType_STATUS_ARM `json:"oauth2grantType,omitempty"`

	// Parameters: Authorization parameters
	Parameters map[string]string `json:"parameters,omitempty"`

	// Status: Status of the Authorization
	Status *string `json:"status,omitempty"`
}

type AuthorizationContractProperties_AuthorizationType_STATUS_ARM string

const AuthorizationContractProperties_AuthorizationType_STATUS_ARM_OAuth2 = AuthorizationContractProperties_AuthorizationType_STATUS_ARM("OAuth2")

// Mapping from string to AuthorizationContractProperties_AuthorizationType_STATUS_ARM
var authorizationContractProperties_AuthorizationType_STATUS_ARM_Values = map[string]AuthorizationContractProperties_AuthorizationType_STATUS_ARM{
	"oauth2": AuthorizationContractProperties_AuthorizationType_STATUS_ARM_OAuth2,
}

type AuthorizationContractProperties_Oauth2GrantType_STATUS_ARM string

const (
	AuthorizationContractProperties_Oauth2GrantType_STATUS_ARM_AuthorizationCode = AuthorizationContractProperties_Oauth2GrantType_STATUS_ARM("AuthorizationCode")
	AuthorizationContractProperties_Oauth2GrantType_STATUS_ARM_ClientCredentials = AuthorizationContractProperties_Oauth2GrantType_STATUS_ARM("ClientCredentials")
)

// Mapping from string to AuthorizationContractProperties_Oauth2GrantType_STATUS_ARM
var authorizationContractProperties_Oauth2GrantType_STATUS_ARM_Values = map[string]AuthorizationContractProperties_Oauth2GrantType_STATUS_ARM{
	"authorizationcode": AuthorizationContractProperties_Oauth2GrantType_STATUS_ARM_AuthorizationCode,
	"clientcredentials": AuthorizationContractProperties_Oauth2GrantType_STATUS_ARM_ClientCredentials,
}

// Authorization error details.
type AuthorizationError_STATUS_ARM struct {
	// Code: Error code
	Code *string `json:"code,omitempty"`

	// Message: Error message
	Message *string `json:"message,omitempty"`
}