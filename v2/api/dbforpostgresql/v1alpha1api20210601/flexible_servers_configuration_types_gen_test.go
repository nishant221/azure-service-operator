// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1alpha1api20210601storage"
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

func Test_FlexibleServersConfiguration_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersConfiguration to hub returns original",
		prop.ForAll(RunResourceConversionTestForFlexibleServersConfiguration, FlexibleServersConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForFlexibleServersConfiguration tests if a specific instance of FlexibleServersConfiguration round trips to the hub storage version and back losslessly
func RunResourceConversionTestForFlexibleServersConfiguration(subject FlexibleServersConfiguration) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1alpha1api20210601storage.FlexibleServersConfiguration
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual FlexibleServersConfiguration
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_FlexibleServersConfiguration_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersConfiguration to FlexibleServersConfiguration via AssignPropertiesToFlexibleServersConfiguration & AssignPropertiesFromFlexibleServersConfiguration returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersConfiguration, FlexibleServersConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersConfiguration tests if a specific instance of FlexibleServersConfiguration can be assigned to v1alpha1api20210601storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersConfiguration(subject FlexibleServersConfiguration) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210601storage.FlexibleServersConfiguration
	err := copied.AssignPropertiesToFlexibleServersConfiguration(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersConfiguration
	err = actual.AssignPropertiesFromFlexibleServersConfiguration(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_FlexibleServersConfiguration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersConfiguration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersConfiguration, FlexibleServersConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersConfiguration runs a test to see if a specific instance of FlexibleServersConfiguration round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersConfiguration(subject FlexibleServersConfiguration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersConfiguration
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

// Generator of FlexibleServersConfiguration instances for property testing - lazily instantiated by
//FlexibleServersConfigurationGenerator()
var flexibleServersConfigurationGenerator gopter.Gen

// FlexibleServersConfigurationGenerator returns a generator of FlexibleServersConfiguration instances for property testing.
func FlexibleServersConfigurationGenerator() gopter.Gen {
	if flexibleServersConfigurationGenerator != nil {
		return flexibleServersConfigurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFlexibleServersConfiguration(generators)
	flexibleServersConfigurationGenerator = gen.Struct(reflect.TypeOf(FlexibleServersConfiguration{}), generators)

	return flexibleServersConfigurationGenerator
}

// AddRelatedPropertyGeneratorsForFlexibleServersConfiguration is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersConfiguration(gens map[string]gopter.Gen) {
	gens["Spec"] = FlexibleServersConfigurationsSpecGenerator()
	gens["Status"] = ConfigurationStatusGenerator()
}

func Test_Configuration_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Configuration_Status to Configuration_Status via AssignPropertiesToConfigurationStatus & AssignPropertiesFromConfigurationStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForConfigurationStatus, ConfigurationStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForConfigurationStatus tests if a specific instance of Configuration_Status can be assigned to v1alpha1api20210601storage and back losslessly
func RunPropertyAssignmentTestForConfigurationStatus(subject Configuration_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210601storage.Configuration_Status
	err := copied.AssignPropertiesToConfigurationStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Configuration_Status
	err = actual.AssignPropertiesFromConfigurationStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Configuration_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Configuration_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfigurationStatus, ConfigurationStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfigurationStatus runs a test to see if a specific instance of Configuration_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForConfigurationStatus(subject Configuration_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Configuration_Status
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

// Generator of Configuration_Status instances for property testing - lazily instantiated by
//ConfigurationStatusGenerator()
var configurationStatusGenerator gopter.Gen

// ConfigurationStatusGenerator returns a generator of Configuration_Status instances for property testing.
// We first initialize configurationStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ConfigurationStatusGenerator() gopter.Gen {
	if configurationStatusGenerator != nil {
		return configurationStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationStatus(generators)
	configurationStatusGenerator = gen.Struct(reflect.TypeOf(Configuration_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationStatus(generators)
	AddRelatedPropertyGeneratorsForConfigurationStatus(generators)
	configurationStatusGenerator = gen.Struct(reflect.TypeOf(Configuration_Status{}), generators)

	return configurationStatusGenerator
}

// AddIndependentPropertyGeneratorsForConfigurationStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConfigurationStatus(gens map[string]gopter.Gen) {
	gens["AllowedValues"] = gen.PtrOf(gen.AlphaString())
	gens["DataType"] = gen.PtrOf(gen.OneConstOf(
		ConfigurationPropertiesStatusDataTypeBoolean,
		ConfigurationPropertiesStatusDataTypeEnumeration,
		ConfigurationPropertiesStatusDataTypeInteger,
		ConfigurationPropertiesStatusDataTypeNumeric))
	gens["DefaultValue"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForConfigurationStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForConfigurationStatus(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemDataStatusGenerator())
}

func Test_FlexibleServersConfigurations_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersConfigurations_Spec to FlexibleServersConfigurations_Spec via AssignPropertiesToFlexibleServersConfigurationsSpec & AssignPropertiesFromFlexibleServersConfigurationsSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersConfigurationsSpec, FlexibleServersConfigurationsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersConfigurationsSpec tests if a specific instance of FlexibleServersConfigurations_Spec can be assigned to v1alpha1api20210601storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersConfigurationsSpec(subject FlexibleServersConfigurations_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210601storage.FlexibleServersConfigurations_Spec
	err := copied.AssignPropertiesToFlexibleServersConfigurationsSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersConfigurations_Spec
	err = actual.AssignPropertiesFromFlexibleServersConfigurationsSpec(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_FlexibleServersConfigurations_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersConfigurations_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersConfigurationsSpec, FlexibleServersConfigurationsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersConfigurationsSpec runs a test to see if a specific instance of FlexibleServersConfigurations_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersConfigurationsSpec(subject FlexibleServersConfigurations_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersConfigurations_Spec
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

// Generator of FlexibleServersConfigurations_Spec instances for property testing - lazily instantiated by
//FlexibleServersConfigurationsSpecGenerator()
var flexibleServersConfigurationsSpecGenerator gopter.Gen

// FlexibleServersConfigurationsSpecGenerator returns a generator of FlexibleServersConfigurations_Spec instances for property testing.
func FlexibleServersConfigurationsSpecGenerator() gopter.Gen {
	if flexibleServersConfigurationsSpecGenerator != nil {
		return flexibleServersConfigurationsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersConfigurationsSpec(generators)
	flexibleServersConfigurationsSpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServersConfigurations_Spec{}), generators)

	return flexibleServersConfigurationsSpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersConfigurationsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersConfigurationsSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}