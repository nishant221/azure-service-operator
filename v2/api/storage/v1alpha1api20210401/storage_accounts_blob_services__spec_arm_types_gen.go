// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccountsBlobServices_SpecARM struct {
	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	//Name: The name of the blob Service within the specified storage account. Blob Service Name must be 'default'
	Name string `json:"name"`

	//Properties: The properties of a storage account’s Blob service.
	Properties *BlobServicePropertiesPropertiesARM `json:"properties,omitempty"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccountsBlobServices_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (services StorageAccountsBlobServices_SpecARM) GetAPIVersion() string {
	return "2021-04-01"
}

// GetName returns the Name of the resource
func (services StorageAccountsBlobServices_SpecARM) GetName() string {
	return services.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/blobServices"
func (services StorageAccountsBlobServices_SpecARM) GetType() string {
	return "Microsoft.Storage/storageAccounts/blobServices"
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/BlobServicePropertiesProperties
type BlobServicePropertiesPropertiesARM struct {
	//AutomaticSnapshotPolicyEnabled: Deprecated in favor of isVersioningEnabled property.
	AutomaticSnapshotPolicyEnabled *bool `json:"automaticSnapshotPolicyEnabled,omitempty"`

	//ChangeFeed: The blob service properties for change feed events.
	ChangeFeed *ChangeFeedARM `json:"changeFeed,omitempty"`

	//ContainerDeleteRetentionPolicy: The service properties for soft delete.
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicyARM `json:"containerDeleteRetentionPolicy,omitempty"`

	//Cors: Sets the CORS rules. You can include up to five CorsRule elements in the request.
	Cors *CorsRulesARM `json:"cors,omitempty"`

	//DefaultServiceVersion: DefaultServiceVersion indicates the default version to use for requests to the Blob service if an
	//incoming request’s version is not specified. Possible values include version 2008-10-27 and all more recent versions.
	DefaultServiceVersion *string `json:"defaultServiceVersion,omitempty"`

	//DeleteRetentionPolicy: The service properties for soft delete.
	DeleteRetentionPolicy *DeleteRetentionPolicyARM `json:"deleteRetentionPolicy,omitempty"`

	//IsVersioningEnabled: Versioning is enabled if set to true.
	IsVersioningEnabled *bool `json:"isVersioningEnabled,omitempty"`

	//LastAccessTimeTrackingPolicy: The blob service properties for Last access time based tracking policy.
	LastAccessTimeTrackingPolicy *LastAccessTimeTrackingPolicyARM `json:"lastAccessTimeTrackingPolicy,omitempty"`

	//RestorePolicy: The blob service properties for blob restore policy
	RestorePolicy *RestorePolicyPropertiesARM `json:"restorePolicy,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ChangeFeed
type ChangeFeedARM struct {
	//Enabled: Indicates whether change feed event logging is enabled for the Blob service.
	Enabled *bool `json:"enabled,omitempty"`

	//RetentionInDays: Indicates the duration of changeFeed retention in days. Minimum value is 1 day and maximum value is
	//146000 days (400 years). A null value indicates an infinite retention of the change feed.
	RetentionInDays *int `json:"retentionInDays,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/CorsRules
type CorsRulesARM struct {
	//CorsRules: The List of CORS rules. You can include up to five CorsRule elements in the request.
	CorsRules []CorsRuleARM `json:"corsRules,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/DeleteRetentionPolicy
type DeleteRetentionPolicyARM struct {
	//Days: Indicates the number of days that the deleted item should be retained. The minimum specified value can be 1 and
	//the maximum value can be 365.
	Days *int `json:"days,omitempty"`

	//Enabled: Indicates whether DeleteRetentionPolicy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/LastAccessTimeTrackingPolicy
type LastAccessTimeTrackingPolicyARM struct {
	//BlobType: An array of predefined supported blob types. Only blockBlob is the supported value. This field is currently
	//read only
	BlobType []string `json:"blobType,omitempty"`

	//Enable: When set to true last access time based tracking is enabled.
	Enable bool `json:"enable"`

	//Name: Name of the policy. The valid value is AccessTimeTracking. This field is currently read only.
	Name *LastAccessTimeTrackingPolicyName `json:"name,omitempty"`

	//TrackingGranularityInDays: The field specifies blob object tracking granularity in days, typically how often the blob
	//object should be tracked.This field is currently read only with value as 1
	TrackingGranularityInDays *int `json:"trackingGranularityInDays,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/RestorePolicyProperties
type RestorePolicyPropertiesARM struct {
	//Days: how long this blob can be restored. It should be great than zero and less than DeleteRetentionPolicy.days.
	Days *int `json:"days,omitempty"`

	//Enabled: Blob restore is enabled if set to true.
	Enabled bool `json:"enabled"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/CorsRule
type CorsRuleARM struct {
	//AllowedHeaders: Required if CorsRule element is present. A list of headers allowed to be part of the cross-origin
	//request.
	AllowedHeaders []string `json:"allowedHeaders"`

	//AllowedMethods: Required if CorsRule element is present. A list of HTTP methods that are allowed to be executed by the
	//origin.
	AllowedMethods []CorsRuleAllowedMethods `json:"allowedMethods"`

	//AllowedOrigins: Required if CorsRule element is present. A list of origin domains that will be allowed via CORS, or "*"
	//to allow all domains
	AllowedOrigins []string `json:"allowedOrigins"`

	//ExposedHeaders: Required if CorsRule element is present. A list of response headers to expose to CORS clients.
	ExposedHeaders []string `json:"exposedHeaders"`

	//MaxAgeInSeconds: Required if CorsRule element is present. The number of seconds that the client/browser should cache a
	//preflight response.
	MaxAgeInSeconds int `json:"maxAgeInSeconds"`
}
