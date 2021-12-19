// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DomainsTopics_SpecARM struct {
	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	//Name: Name of the domain topic.
	Name string `json:"name"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DomainsTopics_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (domainsTopicsSpecARM DomainsTopics_SpecARM) GetAPIVersion() string {
	return "2020-06-01"
}

// GetName returns the Name of the resource
func (domainsTopicsSpecARM DomainsTopics_SpecARM) GetName() string {
	return domainsTopicsSpecARM.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/domains/topics"
func (domainsTopicsSpecARM DomainsTopics_SpecARM) GetType() string {
	return "Microsoft.EventGrid/domains/topics"
}