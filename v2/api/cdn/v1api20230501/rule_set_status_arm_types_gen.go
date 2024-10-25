// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

type RuleSet_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: The JSON object that contains the properties of the Rule Set to create.
	Properties *RuleSetProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: Read only system data
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// The JSON object that contains the properties of the Rule Set to create.
type RuleSetProperties_STATUS_ARM struct {
	DeploymentStatus *RuleSetProperties_DeploymentStatus_STATUS_ARM `json:"deploymentStatus,omitempty"`

	// ProfileName: The name of the profile which holds the rule set.
	ProfileName *string `json:"profileName,omitempty"`

	// ProvisioningState: Provisioning status
	ProvisioningState *RuleSetProperties_ProvisioningState_STATUS_ARM `json:"provisioningState,omitempty"`
}

type RuleSetProperties_DeploymentStatus_STATUS_ARM string

const (
	RuleSetProperties_DeploymentStatus_STATUS_ARM_Failed     = RuleSetProperties_DeploymentStatus_STATUS_ARM("Failed")
	RuleSetProperties_DeploymentStatus_STATUS_ARM_InProgress = RuleSetProperties_DeploymentStatus_STATUS_ARM("InProgress")
	RuleSetProperties_DeploymentStatus_STATUS_ARM_NotStarted = RuleSetProperties_DeploymentStatus_STATUS_ARM("NotStarted")
	RuleSetProperties_DeploymentStatus_STATUS_ARM_Succeeded  = RuleSetProperties_DeploymentStatus_STATUS_ARM("Succeeded")
)

// Mapping from string to RuleSetProperties_DeploymentStatus_STATUS_ARM
var ruleSetProperties_DeploymentStatus_STATUS_ARM_Values = map[string]RuleSetProperties_DeploymentStatus_STATUS_ARM{
	"failed":     RuleSetProperties_DeploymentStatus_STATUS_ARM_Failed,
	"inprogress": RuleSetProperties_DeploymentStatus_STATUS_ARM_InProgress,
	"notstarted": RuleSetProperties_DeploymentStatus_STATUS_ARM_NotStarted,
	"succeeded":  RuleSetProperties_DeploymentStatus_STATUS_ARM_Succeeded,
}

type RuleSetProperties_ProvisioningState_STATUS_ARM string

const (
	RuleSetProperties_ProvisioningState_STATUS_ARM_Creating  = RuleSetProperties_ProvisioningState_STATUS_ARM("Creating")
	RuleSetProperties_ProvisioningState_STATUS_ARM_Deleting  = RuleSetProperties_ProvisioningState_STATUS_ARM("Deleting")
	RuleSetProperties_ProvisioningState_STATUS_ARM_Failed    = RuleSetProperties_ProvisioningState_STATUS_ARM("Failed")
	RuleSetProperties_ProvisioningState_STATUS_ARM_Succeeded = RuleSetProperties_ProvisioningState_STATUS_ARM("Succeeded")
	RuleSetProperties_ProvisioningState_STATUS_ARM_Updating  = RuleSetProperties_ProvisioningState_STATUS_ARM("Updating")
)

// Mapping from string to RuleSetProperties_ProvisioningState_STATUS_ARM
var ruleSetProperties_ProvisioningState_STATUS_ARM_Values = map[string]RuleSetProperties_ProvisioningState_STATUS_ARM{
	"creating":  RuleSetProperties_ProvisioningState_STATUS_ARM_Creating,
	"deleting":  RuleSetProperties_ProvisioningState_STATUS_ARM_Deleting,
	"failed":    RuleSetProperties_ProvisioningState_STATUS_ARM_Failed,
	"succeeded": RuleSetProperties_ProvisioningState_STATUS_ARM_Succeeded,
	"updating":  RuleSetProperties_ProvisioningState_STATUS_ARM_Updating,
}