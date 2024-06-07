// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200202

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

func Test_ApplicationInsightsComponentProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationInsightsComponentProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationInsightsComponentProperties_STATUS_ARM, ApplicationInsightsComponentProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationInsightsComponentProperties_STATUS_ARM runs a test to see if a specific instance of ApplicationInsightsComponentProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationInsightsComponentProperties_STATUS_ARM(subject ApplicationInsightsComponentProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationInsightsComponentProperties_STATUS_ARM
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

// Generator of ApplicationInsightsComponentProperties_STATUS_ARM instances for property testing - lazily instantiated
// by ApplicationInsightsComponentProperties_STATUS_ARMGenerator()
var applicationInsightsComponentProperties_STATUS_ARMGenerator gopter.Gen

// ApplicationInsightsComponentProperties_STATUS_ARMGenerator returns a generator of ApplicationInsightsComponentProperties_STATUS_ARM instances for property testing.
// We first initialize applicationInsightsComponentProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ApplicationInsightsComponentProperties_STATUS_ARMGenerator() gopter.Gen {
	if applicationInsightsComponentProperties_STATUS_ARMGenerator != nil {
		return applicationInsightsComponentProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationInsightsComponentProperties_STATUS_ARM(generators)
	applicationInsightsComponentProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ApplicationInsightsComponentProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationInsightsComponentProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForApplicationInsightsComponentProperties_STATUS_ARM(generators)
	applicationInsightsComponentProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ApplicationInsightsComponentProperties_STATUS_ARM{}), generators)

	return applicationInsightsComponentProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForApplicationInsightsComponentProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationInsightsComponentProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AppId"] = gen.PtrOf(gen.AlphaString())
	gens["ApplicationId"] = gen.PtrOf(gen.AlphaString())
	gens["Application_Type"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_Application_Type_STATUS_Other, ApplicationInsightsComponentProperties_Application_Type_STATUS_Web))
	gens["ConnectionString"] = gen.PtrOf(gen.AlphaString())
	gens["CreationDate"] = gen.PtrOf(gen.AlphaString())
	gens["DisableIpMasking"] = gen.PtrOf(gen.Bool())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["Flow_Type"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_Flow_Type_STATUS_Bluefield))
	gens["ForceCustomerStorageForProfiler"] = gen.PtrOf(gen.Bool())
	gens["HockeyAppId"] = gen.PtrOf(gen.AlphaString())
	gens["HockeyAppToken"] = gen.PtrOf(gen.AlphaString())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
	gens["IngestionMode"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_IngestionMode_STATUS_ApplicationInsights, ApplicationInsightsComponentProperties_IngestionMode_STATUS_ApplicationInsightsWithDiagnosticSettings, ApplicationInsightsComponentProperties_IngestionMode_STATUS_LogAnalytics))
	gens["InstrumentationKey"] = gen.PtrOf(gen.AlphaString())
	gens["LaMigrationDate"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_STATUS_Disabled, PublicNetworkAccessType_STATUS_Enabled))
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_STATUS_Disabled, PublicNetworkAccessType_STATUS_Enabled))
	gens["Request_Source"] = gen.PtrOf(gen.OneConstOf(ApplicationInsightsComponentProperties_Request_Source_STATUS_Rest))
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
	gens["SamplingPercentage"] = gen.PtrOf(gen.Float64())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["WorkspaceResourceId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForApplicationInsightsComponentProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForApplicationInsightsComponentProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PrivateLinkScopedResources"] = gen.SliceOf(PrivateLinkScopedResource_STATUS_ARMGenerator())
}

func Test_Component_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Component_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForComponent_STATUS_ARM, Component_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForComponent_STATUS_ARM runs a test to see if a specific instance of Component_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForComponent_STATUS_ARM(subject Component_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Component_STATUS_ARM
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

// Generator of Component_STATUS_ARM instances for property testing - lazily instantiated by
// Component_STATUS_ARMGenerator()
var component_STATUS_ARMGenerator gopter.Gen

// Component_STATUS_ARMGenerator returns a generator of Component_STATUS_ARM instances for property testing.
// We first initialize component_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Component_STATUS_ARMGenerator() gopter.Gen {
	if component_STATUS_ARMGenerator != nil {
		return component_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComponent_STATUS_ARM(generators)
	component_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Component_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForComponent_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForComponent_STATUS_ARM(generators)
	component_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Component_STATUS_ARM{}), generators)

	return component_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForComponent_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForComponent_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForComponent_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForComponent_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ApplicationInsightsComponentProperties_STATUS_ARMGenerator())
}

func Test_PrivateLinkScopedResource_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkScopedResource_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkScopedResource_STATUS_ARM, PrivateLinkScopedResource_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkScopedResource_STATUS_ARM runs a test to see if a specific instance of PrivateLinkScopedResource_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkScopedResource_STATUS_ARM(subject PrivateLinkScopedResource_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkScopedResource_STATUS_ARM
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

// Generator of PrivateLinkScopedResource_STATUS_ARM instances for property testing - lazily instantiated by
// PrivateLinkScopedResource_STATUS_ARMGenerator()
var privateLinkScopedResource_STATUS_ARMGenerator gopter.Gen

// PrivateLinkScopedResource_STATUS_ARMGenerator returns a generator of PrivateLinkScopedResource_STATUS_ARM instances for property testing.
func PrivateLinkScopedResource_STATUS_ARMGenerator() gopter.Gen {
	if privateLinkScopedResource_STATUS_ARMGenerator != nil {
		return privateLinkScopedResource_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkScopedResource_STATUS_ARM(generators)
	privateLinkScopedResource_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateLinkScopedResource_STATUS_ARM{}), generators)

	return privateLinkScopedResource_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkScopedResource_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkScopedResource_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["ScopeId"] = gen.PtrOf(gen.AlphaString())
}
