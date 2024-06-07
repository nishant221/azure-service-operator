// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/storage"
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

func Test_ServersAzureADOnlyAuthentication_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersAzureADOnlyAuthentication to hub returns original",
		prop.ForAll(RunResourceConversionTestForServersAzureADOnlyAuthentication, ServersAzureADOnlyAuthenticationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForServersAzureADOnlyAuthentication tests if a specific instance of ServersAzureADOnlyAuthentication round trips to the hub storage version and back losslessly
func RunResourceConversionTestForServersAzureADOnlyAuthentication(subject ServersAzureADOnlyAuthentication) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.ServersAzureADOnlyAuthentication
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ServersAzureADOnlyAuthentication
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

func Test_ServersAzureADOnlyAuthentication_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersAzureADOnlyAuthentication to ServersAzureADOnlyAuthentication via AssignProperties_To_ServersAzureADOnlyAuthentication & AssignProperties_From_ServersAzureADOnlyAuthentication returns original",
		prop.ForAll(RunPropertyAssignmentTestForServersAzureADOnlyAuthentication, ServersAzureADOnlyAuthenticationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServersAzureADOnlyAuthentication tests if a specific instance of ServersAzureADOnlyAuthentication can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForServersAzureADOnlyAuthentication(subject ServersAzureADOnlyAuthentication) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.ServersAzureADOnlyAuthentication
	err := copied.AssignProperties_To_ServersAzureADOnlyAuthentication(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServersAzureADOnlyAuthentication
	err = actual.AssignProperties_From_ServersAzureADOnlyAuthentication(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ServersAzureADOnlyAuthentication_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersAzureADOnlyAuthentication via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersAzureADOnlyAuthentication, ServersAzureADOnlyAuthenticationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersAzureADOnlyAuthentication runs a test to see if a specific instance of ServersAzureADOnlyAuthentication round trips to JSON and back losslessly
func RunJSONSerializationTestForServersAzureADOnlyAuthentication(subject ServersAzureADOnlyAuthentication) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersAzureADOnlyAuthentication
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

// Generator of ServersAzureADOnlyAuthentication instances for property testing - lazily instantiated by
// ServersAzureADOnlyAuthenticationGenerator()
var serversAzureADOnlyAuthenticationGenerator gopter.Gen

// ServersAzureADOnlyAuthenticationGenerator returns a generator of ServersAzureADOnlyAuthentication instances for property testing.
func ServersAzureADOnlyAuthenticationGenerator() gopter.Gen {
	if serversAzureADOnlyAuthenticationGenerator != nil {
		return serversAzureADOnlyAuthenticationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersAzureADOnlyAuthentication(generators)
	serversAzureADOnlyAuthenticationGenerator = gen.Struct(reflect.TypeOf(ServersAzureADOnlyAuthentication{}), generators)

	return serversAzureADOnlyAuthenticationGenerator
}

// AddRelatedPropertyGeneratorsForServersAzureADOnlyAuthentication is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersAzureADOnlyAuthentication(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_AzureADOnlyAuthentication_SpecGenerator()
	gens["Status"] = Servers_AzureADOnlyAuthentication_STATUSGenerator()
}

func Test_Servers_AzureADOnlyAuthentication_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_AzureADOnlyAuthentication_STATUS to Servers_AzureADOnlyAuthentication_STATUS via AssignProperties_To_Servers_AzureADOnlyAuthentication_STATUS & AssignProperties_From_Servers_AzureADOnlyAuthentication_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_AzureADOnlyAuthentication_STATUS, Servers_AzureADOnlyAuthentication_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_AzureADOnlyAuthentication_STATUS tests if a specific instance of Servers_AzureADOnlyAuthentication_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForServers_AzureADOnlyAuthentication_STATUS(subject Servers_AzureADOnlyAuthentication_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Servers_AzureADOnlyAuthentication_STATUS
	err := copied.AssignProperties_To_Servers_AzureADOnlyAuthentication_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_AzureADOnlyAuthentication_STATUS
	err = actual.AssignProperties_From_Servers_AzureADOnlyAuthentication_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Servers_AzureADOnlyAuthentication_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_AzureADOnlyAuthentication_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_AzureADOnlyAuthentication_STATUS, Servers_AzureADOnlyAuthentication_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_AzureADOnlyAuthentication_STATUS runs a test to see if a specific instance of Servers_AzureADOnlyAuthentication_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_AzureADOnlyAuthentication_STATUS(subject Servers_AzureADOnlyAuthentication_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_AzureADOnlyAuthentication_STATUS
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

// Generator of Servers_AzureADOnlyAuthentication_STATUS instances for property testing - lazily instantiated by
// Servers_AzureADOnlyAuthentication_STATUSGenerator()
var servers_AzureADOnlyAuthentication_STATUSGenerator gopter.Gen

// Servers_AzureADOnlyAuthentication_STATUSGenerator returns a generator of Servers_AzureADOnlyAuthentication_STATUS instances for property testing.
func Servers_AzureADOnlyAuthentication_STATUSGenerator() gopter.Gen {
	if servers_AzureADOnlyAuthentication_STATUSGenerator != nil {
		return servers_AzureADOnlyAuthentication_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_AzureADOnlyAuthentication_STATUS(generators)
	servers_AzureADOnlyAuthentication_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_AzureADOnlyAuthentication_STATUS{}), generators)

	return servers_AzureADOnlyAuthentication_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_AzureADOnlyAuthentication_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_AzureADOnlyAuthentication_STATUS(gens map[string]gopter.Gen) {
	gens["AzureADOnlyAuthentication"] = gen.PtrOf(gen.Bool())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_Servers_AzureADOnlyAuthentication_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_AzureADOnlyAuthentication_Spec to Servers_AzureADOnlyAuthentication_Spec via AssignProperties_To_Servers_AzureADOnlyAuthentication_Spec & AssignProperties_From_Servers_AzureADOnlyAuthentication_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_AzureADOnlyAuthentication_Spec, Servers_AzureADOnlyAuthentication_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_AzureADOnlyAuthentication_Spec tests if a specific instance of Servers_AzureADOnlyAuthentication_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForServers_AzureADOnlyAuthentication_Spec(subject Servers_AzureADOnlyAuthentication_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Servers_AzureADOnlyAuthentication_Spec
	err := copied.AssignProperties_To_Servers_AzureADOnlyAuthentication_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_AzureADOnlyAuthentication_Spec
	err = actual.AssignProperties_From_Servers_AzureADOnlyAuthentication_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Servers_AzureADOnlyAuthentication_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_AzureADOnlyAuthentication_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_AzureADOnlyAuthentication_Spec, Servers_AzureADOnlyAuthentication_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_AzureADOnlyAuthentication_Spec runs a test to see if a specific instance of Servers_AzureADOnlyAuthentication_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_AzureADOnlyAuthentication_Spec(subject Servers_AzureADOnlyAuthentication_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_AzureADOnlyAuthentication_Spec
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

// Generator of Servers_AzureADOnlyAuthentication_Spec instances for property testing - lazily instantiated by
// Servers_AzureADOnlyAuthentication_SpecGenerator()
var servers_AzureADOnlyAuthentication_SpecGenerator gopter.Gen

// Servers_AzureADOnlyAuthentication_SpecGenerator returns a generator of Servers_AzureADOnlyAuthentication_Spec instances for property testing.
func Servers_AzureADOnlyAuthentication_SpecGenerator() gopter.Gen {
	if servers_AzureADOnlyAuthentication_SpecGenerator != nil {
		return servers_AzureADOnlyAuthentication_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_AzureADOnlyAuthentication_Spec(generators)
	servers_AzureADOnlyAuthentication_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_AzureADOnlyAuthentication_Spec{}), generators)

	return servers_AzureADOnlyAuthentication_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_AzureADOnlyAuthentication_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_AzureADOnlyAuthentication_Spec(gens map[string]gopter.Gen) {
	gens["AzureADOnlyAuthentication"] = gen.PtrOf(gen.Bool())
}
