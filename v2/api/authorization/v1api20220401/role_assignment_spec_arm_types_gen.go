// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RoleAssignment_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Role assignment properties.
	Properties *RoleAssignmentProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RoleAssignment_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-04-01"
func (assignment RoleAssignment_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (assignment *RoleAssignment_Spec_ARM) GetName() string {
	return assignment.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Authorization/roleAssignments"
func (assignment *RoleAssignment_Spec_ARM) GetType() string {
	return "Microsoft.Authorization/roleAssignments"
}

// Role assignment properties.
type RoleAssignmentProperties_ARM struct {
	// Condition: The conditions on the role assignment. This limits the resources it can be assigned to. e.g.:
	// @Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase
	// 'foo_storage_container'
	Condition *string `json:"condition,omitempty"`

	// ConditionVersion: Version of the condition. Currently the only accepted value is '2.0'
	ConditionVersion                   *string `json:"conditionVersion,omitempty"`
	DelegatedManagedIdentityResourceId *string `json:"delegatedManagedIdentityResourceId,omitempty"`

	// Description: Description of role assignment
	Description *string `json:"description,omitempty"`

	// PrincipalId: The principal ID.
	PrincipalId *string `json:"principalId,omitempty" optionalConfigMapPair:"PrincipalId"`

	// PrincipalType: The principal type of the assigned principal ID.
	PrincipalType    *RoleAssignmentProperties_PrincipalType `json:"principalType,omitempty"`
	RoleDefinitionId *string                                 `json:"roleDefinitionId,omitempty"`
}
