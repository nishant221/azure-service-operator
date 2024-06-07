// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_DataPlaneAadOrApiKeyAuthOption_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DataPlaneAadOrApiKeyAuthOption_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDataPlaneAadOrApiKeyAuthOption_STATUS_ARM, DataPlaneAadOrApiKeyAuthOption_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDataPlaneAadOrApiKeyAuthOption_STATUS_ARM runs a test to see if a specific instance of DataPlaneAadOrApiKeyAuthOption_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDataPlaneAadOrApiKeyAuthOption_STATUS_ARM(subject DataPlaneAadOrApiKeyAuthOption_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DataPlaneAadOrApiKeyAuthOption_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of DataPlaneAadOrApiKeyAuthOption_STATUS_ARM instances for property testing - lazily instantiated by
// DataPlaneAadOrApiKeyAuthOption_STATUS_ARMGenerator()
var dataPlaneAadOrApiKeyAuthOption_STATUS_ARMGenerator gopter.Gen

// DataPlaneAadOrApiKeyAuthOption_STATUS_ARMGenerator returns a generator of DataPlaneAadOrApiKeyAuthOption_STATUS_ARM instances for property testing.
func DataPlaneAadOrApiKeyAuthOption_STATUS_ARMGenerator() gopter.Gen {
	if dataPlaneAadOrApiKeyAuthOption_STATUS_ARMGenerator != nil {
		return dataPlaneAadOrApiKeyAuthOption_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDataPlaneAadOrApiKeyAuthOption_STATUS_ARM(generators)
	dataPlaneAadOrApiKeyAuthOption_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DataPlaneAadOrApiKeyAuthOption_STATUS_ARM{}), generators)

	return dataPlaneAadOrApiKeyAuthOption_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDataPlaneAadOrApiKeyAuthOption_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDataPlaneAadOrApiKeyAuthOption_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AadAuthFailureMode"] = gen.PtrOf(gen.OneConstOf(DataPlaneAadOrApiKeyAuthOption_AadAuthFailureMode_STATUS_Http401WithBearerChallenge, DataPlaneAadOrApiKeyAuthOption_AadAuthFailureMode_STATUS_Http403))
}

func Test_DataPlaneAuthOptions_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DataPlaneAuthOptions_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDataPlaneAuthOptions_STATUS_ARM, DataPlaneAuthOptions_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDataPlaneAuthOptions_STATUS_ARM runs a test to see if a specific instance of DataPlaneAuthOptions_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDataPlaneAuthOptions_STATUS_ARM(subject DataPlaneAuthOptions_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DataPlaneAuthOptions_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of DataPlaneAuthOptions_STATUS_ARM instances for property testing - lazily instantiated by
// DataPlaneAuthOptions_STATUS_ARMGenerator()
var dataPlaneAuthOptions_STATUS_ARMGenerator gopter.Gen

// DataPlaneAuthOptions_STATUS_ARMGenerator returns a generator of DataPlaneAuthOptions_STATUS_ARM instances for property testing.
func DataPlaneAuthOptions_STATUS_ARMGenerator() gopter.Gen {
	if dataPlaneAuthOptions_STATUS_ARMGenerator != nil {
		return dataPlaneAuthOptions_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDataPlaneAuthOptions_STATUS_ARM(generators)
	dataPlaneAuthOptions_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DataPlaneAuthOptions_STATUS_ARM{}), generators)

	return dataPlaneAuthOptions_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForDataPlaneAuthOptions_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDataPlaneAuthOptions_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AadOrApiKey"] = gen.PtrOf(DataPlaneAadOrApiKeyAuthOption_STATUS_ARMGenerator())
}

func Test_EncryptionWithCmk_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionWithCmk_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionWithCmk_STATUS_ARM, EncryptionWithCmk_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionWithCmk_STATUS_ARM runs a test to see if a specific instance of EncryptionWithCmk_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionWithCmk_STATUS_ARM(subject EncryptionWithCmk_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionWithCmk_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of EncryptionWithCmk_STATUS_ARM instances for property testing - lazily instantiated by
// EncryptionWithCmk_STATUS_ARMGenerator()
var encryptionWithCmk_STATUS_ARMGenerator gopter.Gen

// EncryptionWithCmk_STATUS_ARMGenerator returns a generator of EncryptionWithCmk_STATUS_ARM instances for property testing.
func EncryptionWithCmk_STATUS_ARMGenerator() gopter.Gen {
	if encryptionWithCmk_STATUS_ARMGenerator != nil {
		return encryptionWithCmk_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionWithCmk_STATUS_ARM(generators)
	encryptionWithCmk_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionWithCmk_STATUS_ARM{}), generators)

	return encryptionWithCmk_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionWithCmk_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionWithCmk_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["EncryptionComplianceStatus"] = gen.PtrOf(gen.OneConstOf(EncryptionWithCmk_EncryptionComplianceStatus_STATUS_Compliant, EncryptionWithCmk_EncryptionComplianceStatus_STATUS_NonCompliant))
	gens["Enforcement"] = gen.PtrOf(gen.OneConstOf(EncryptionWithCmk_Enforcement_STATUS_Disabled, EncryptionWithCmk_Enforcement_STATUS_Enabled, EncryptionWithCmk_Enforcement_STATUS_Unspecified))
}

func Test_Identity_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentity_STATUS_ARM, Identity_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentity_STATUS_ARM runs a test to see if a specific instance of Identity_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentity_STATUS_ARM(subject Identity_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Identity_STATUS_ARM instances for property testing - lazily instantiated by
// Identity_STATUS_ARMGenerator()
var identity_STATUS_ARMGenerator gopter.Gen

// Identity_STATUS_ARMGenerator returns a generator of Identity_STATUS_ARM instances for property testing.
func Identity_STATUS_ARMGenerator() gopter.Gen {
	if identity_STATUS_ARMGenerator != nil {
		return identity_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_STATUS_ARM(generators)
	identity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Identity_STATUS_ARM{}), generators)

	return identity_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(Identity_Type_STATUS_None, Identity_Type_STATUS_SystemAssigned))
}

func Test_IpRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IpRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIpRule_STATUS_ARM, IpRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIpRule_STATUS_ARM runs a test to see if a specific instance of IpRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIpRule_STATUS_ARM(subject IpRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IpRule_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of IpRule_STATUS_ARM instances for property testing - lazily instantiated by IpRule_STATUS_ARMGenerator()
var ipRule_STATUS_ARMGenerator gopter.Gen

// IpRule_STATUS_ARMGenerator returns a generator of IpRule_STATUS_ARM instances for property testing.
func IpRule_STATUS_ARMGenerator() gopter.Gen {
	if ipRule_STATUS_ARMGenerator != nil {
		return ipRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpRule_STATUS_ARM(generators)
	ipRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(IpRule_STATUS_ARM{}), generators)

	return ipRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIpRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIpRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkRuleSet_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkRuleSet_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkRuleSet_STATUS_ARM, NetworkRuleSet_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkRuleSet_STATUS_ARM runs a test to see if a specific instance of NetworkRuleSet_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkRuleSet_STATUS_ARM(subject NetworkRuleSet_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkRuleSet_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of NetworkRuleSet_STATUS_ARM instances for property testing - lazily instantiated by
// NetworkRuleSet_STATUS_ARMGenerator()
var networkRuleSet_STATUS_ARMGenerator gopter.Gen

// NetworkRuleSet_STATUS_ARMGenerator returns a generator of NetworkRuleSet_STATUS_ARM instances for property testing.
func NetworkRuleSet_STATUS_ARMGenerator() gopter.Gen {
	if networkRuleSet_STATUS_ARMGenerator != nil {
		return networkRuleSet_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNetworkRuleSet_STATUS_ARM(generators)
	networkRuleSet_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkRuleSet_STATUS_ARM{}), generators)

	return networkRuleSet_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForNetworkRuleSet_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkRuleSet_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["IpRules"] = gen.SliceOf(IpRule_STATUS_ARMGenerator())
}

func Test_PrivateEndpointConnection_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_STATUS_ARM, PrivateEndpointConnection_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_STATUS_ARM runs a test to see if a specific instance of PrivateEndpointConnection_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_STATUS_ARM(subject PrivateEndpointConnection_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of PrivateEndpointConnection_STATUS_ARM instances for property testing - lazily instantiated by
// PrivateEndpointConnection_STATUS_ARMGenerator()
var privateEndpointConnection_STATUS_ARMGenerator gopter.Gen

// PrivateEndpointConnection_STATUS_ARMGenerator returns a generator of PrivateEndpointConnection_STATUS_ARM instances for property testing.
func PrivateEndpointConnection_STATUS_ARMGenerator() gopter.Gen {
	if privateEndpointConnection_STATUS_ARMGenerator != nil {
		return privateEndpointConnection_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_ARM(generators)
	privateEndpointConnection_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS_ARM{}), generators)

	return privateEndpointConnection_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_SearchServiceProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SearchServiceProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSearchServiceProperties_STATUS_ARM, SearchServiceProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSearchServiceProperties_STATUS_ARM runs a test to see if a specific instance of SearchServiceProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSearchServiceProperties_STATUS_ARM(subject SearchServiceProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SearchServiceProperties_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SearchServiceProperties_STATUS_ARM instances for property testing - lazily instantiated by
// SearchServiceProperties_STATUS_ARMGenerator()
var searchServiceProperties_STATUS_ARMGenerator gopter.Gen

// SearchServiceProperties_STATUS_ARMGenerator returns a generator of SearchServiceProperties_STATUS_ARM instances for property testing.
// We first initialize searchServiceProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SearchServiceProperties_STATUS_ARMGenerator() gopter.Gen {
	if searchServiceProperties_STATUS_ARMGenerator != nil {
		return searchServiceProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSearchServiceProperties_STATUS_ARM(generators)
	searchServiceProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SearchServiceProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSearchServiceProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForSearchServiceProperties_STATUS_ARM(generators)
	searchServiceProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SearchServiceProperties_STATUS_ARM{}), generators)

	return searchServiceProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSearchServiceProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSearchServiceProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["HostingMode"] = gen.PtrOf(gen.OneConstOf(SearchServiceProperties_HostingMode_STATUS_Default, SearchServiceProperties_HostingMode_STATUS_HighDensity))
	gens["PartitionCount"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(SearchServiceProperties_ProvisioningState_STATUS_Failed, SearchServiceProperties_ProvisioningState_STATUS_Provisioning, SearchServiceProperties_ProvisioningState_STATUS_Succeeded))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(SearchServiceProperties_PublicNetworkAccess_STATUS_Disabled, SearchServiceProperties_PublicNetworkAccess_STATUS_Enabled))
	gens["ReplicaCount"] = gen.PtrOf(gen.Int())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		SearchServiceProperties_Status_STATUS_Degraded,
		SearchServiceProperties_Status_STATUS_Deleting,
		SearchServiceProperties_Status_STATUS_Disabled,
		SearchServiceProperties_Status_STATUS_Error,
		SearchServiceProperties_Status_STATUS_Provisioning,
		SearchServiceProperties_Status_STATUS_Running))
	gens["StatusDetails"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSearchServiceProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSearchServiceProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AuthOptions"] = gen.PtrOf(DataPlaneAuthOptions_STATUS_ARMGenerator())
	gens["EncryptionWithCmk"] = gen.PtrOf(EncryptionWithCmk_STATUS_ARMGenerator())
	gens["NetworkRuleSet"] = gen.PtrOf(NetworkRuleSet_STATUS_ARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_STATUS_ARMGenerator())
	gens["SharedPrivateLinkResources"] = gen.SliceOf(SharedPrivateLinkResource_STATUS_ARMGenerator())
}

func Test_SearchService_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SearchService_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSearchService_STATUS_ARM, SearchService_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSearchService_STATUS_ARM runs a test to see if a specific instance of SearchService_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSearchService_STATUS_ARM(subject SearchService_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SearchService_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SearchService_STATUS_ARM instances for property testing - lazily instantiated by
// SearchService_STATUS_ARMGenerator()
var searchService_STATUS_ARMGenerator gopter.Gen

// SearchService_STATUS_ARMGenerator returns a generator of SearchService_STATUS_ARM instances for property testing.
// We first initialize searchService_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SearchService_STATUS_ARMGenerator() gopter.Gen {
	if searchService_STATUS_ARMGenerator != nil {
		return searchService_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSearchService_STATUS_ARM(generators)
	searchService_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SearchService_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSearchService_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForSearchService_STATUS_ARM(generators)
	searchService_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SearchService_STATUS_ARM{}), generators)

	return searchService_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSearchService_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSearchService_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSearchService_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSearchService_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(Identity_STATUS_ARMGenerator())
	gens["Properties"] = gen.PtrOf(SearchServiceProperties_STATUS_ARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUS_ARMGenerator())
}

func Test_SharedPrivateLinkResource_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SharedPrivateLinkResource_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSharedPrivateLinkResource_STATUS_ARM, SharedPrivateLinkResource_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSharedPrivateLinkResource_STATUS_ARM runs a test to see if a specific instance of SharedPrivateLinkResource_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSharedPrivateLinkResource_STATUS_ARM(subject SharedPrivateLinkResource_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SharedPrivateLinkResource_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SharedPrivateLinkResource_STATUS_ARM instances for property testing - lazily instantiated by
// SharedPrivateLinkResource_STATUS_ARMGenerator()
var sharedPrivateLinkResource_STATUS_ARMGenerator gopter.Gen

// SharedPrivateLinkResource_STATUS_ARMGenerator returns a generator of SharedPrivateLinkResource_STATUS_ARM instances for property testing.
func SharedPrivateLinkResource_STATUS_ARMGenerator() gopter.Gen {
	if sharedPrivateLinkResource_STATUS_ARMGenerator != nil {
		return sharedPrivateLinkResource_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResource_STATUS_ARM(generators)
	sharedPrivateLinkResource_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResource_STATUS_ARM{}), generators)

	return sharedPrivateLinkResource_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSharedPrivateLinkResource_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSharedPrivateLinkResource_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUS_ARM, Sku_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUS_ARM runs a test to see if a specific instance of Sku_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUS_ARM(subject Sku_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Sku_STATUS_ARM instances for property testing - lazily instantiated by Sku_STATUS_ARMGenerator()
var sku_STATUS_ARMGenerator gopter.Gen

// Sku_STATUS_ARMGenerator returns a generator of Sku_STATUS_ARM instances for property testing.
func Sku_STATUS_ARMGenerator() gopter.Gen {
	if sku_STATUS_ARMGenerator != nil {
		return sku_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUS_ARM(generators)
	sku_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Sku_STATUS_ARM{}), generators)

	return sku_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		Sku_Name_STATUS_Basic,
		Sku_Name_STATUS_Free,
		Sku_Name_STATUS_Standard,
		Sku_Name_STATUS_Standard2,
		Sku_Name_STATUS_Standard3,
		Sku_Name_STATUS_Storage_Optimized_L1,
		Sku_Name_STATUS_Storage_Optimized_L2))
}
