// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240601

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20240601/storage"
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

func Test_PrivateDnsZonesVirtualNetworkLink_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesVirtualNetworkLink to hub returns original",
		prop.ForAll(RunResourceConversionTestForPrivateDnsZonesVirtualNetworkLink, PrivateDnsZonesVirtualNetworkLinkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForPrivateDnsZonesVirtualNetworkLink tests if a specific instance of PrivateDnsZonesVirtualNetworkLink round trips to the hub storage version and back losslessly
func RunResourceConversionTestForPrivateDnsZonesVirtualNetworkLink(subject PrivateDnsZonesVirtualNetworkLink) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.PrivateDnsZonesVirtualNetworkLink
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual PrivateDnsZonesVirtualNetworkLink
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

func Test_PrivateDnsZonesVirtualNetworkLink_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesVirtualNetworkLink to PrivateDnsZonesVirtualNetworkLink via AssignProperties_To_PrivateDnsZonesVirtualNetworkLink & AssignProperties_From_PrivateDnsZonesVirtualNetworkLink returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesVirtualNetworkLink, PrivateDnsZonesVirtualNetworkLinkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesVirtualNetworkLink tests if a specific instance of PrivateDnsZonesVirtualNetworkLink can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesVirtualNetworkLink(subject PrivateDnsZonesVirtualNetworkLink) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateDnsZonesVirtualNetworkLink
	err := copied.AssignProperties_To_PrivateDnsZonesVirtualNetworkLink(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesVirtualNetworkLink
	err = actual.AssignProperties_From_PrivateDnsZonesVirtualNetworkLink(&other)
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

func Test_PrivateDnsZonesVirtualNetworkLink_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesVirtualNetworkLink via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink, PrivateDnsZonesVirtualNetworkLinkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink runs a test to see if a specific instance of PrivateDnsZonesVirtualNetworkLink round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink(subject PrivateDnsZonesVirtualNetworkLink) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesVirtualNetworkLink
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

// Generator of PrivateDnsZonesVirtualNetworkLink instances for property testing - lazily instantiated by
// PrivateDnsZonesVirtualNetworkLinkGenerator()
var privateDnsZonesVirtualNetworkLinkGenerator gopter.Gen

// PrivateDnsZonesVirtualNetworkLinkGenerator returns a generator of PrivateDnsZonesVirtualNetworkLink instances for property testing.
func PrivateDnsZonesVirtualNetworkLinkGenerator() gopter.Gen {
	if privateDnsZonesVirtualNetworkLinkGenerator != nil {
		return privateDnsZonesVirtualNetworkLinkGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink(generators)
	privateDnsZonesVirtualNetworkLinkGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesVirtualNetworkLink{}), generators)

	return privateDnsZonesVirtualNetworkLinkGenerator
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink(gens map[string]gopter.Gen) {
	gens["Spec"] = PrivateDnsZonesVirtualNetworkLink_SpecGenerator()
	gens["Status"] = PrivateDnsZonesVirtualNetworkLink_STATUSGenerator()
}

func Test_PrivateDnsZonesVirtualNetworkLinkOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesVirtualNetworkLinkOperatorSpec to PrivateDnsZonesVirtualNetworkLinkOperatorSpec via AssignProperties_To_PrivateDnsZonesVirtualNetworkLinkOperatorSpec & AssignProperties_From_PrivateDnsZonesVirtualNetworkLinkOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesVirtualNetworkLinkOperatorSpec, PrivateDnsZonesVirtualNetworkLinkOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesVirtualNetworkLinkOperatorSpec tests if a specific instance of PrivateDnsZonesVirtualNetworkLinkOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesVirtualNetworkLinkOperatorSpec(subject PrivateDnsZonesVirtualNetworkLinkOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateDnsZonesVirtualNetworkLinkOperatorSpec
	err := copied.AssignProperties_To_PrivateDnsZonesVirtualNetworkLinkOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesVirtualNetworkLinkOperatorSpec
	err = actual.AssignProperties_From_PrivateDnsZonesVirtualNetworkLinkOperatorSpec(&other)
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

func Test_PrivateDnsZonesVirtualNetworkLinkOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesVirtualNetworkLinkOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLinkOperatorSpec, PrivateDnsZonesVirtualNetworkLinkOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLinkOperatorSpec runs a test to see if a specific instance of PrivateDnsZonesVirtualNetworkLinkOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLinkOperatorSpec(subject PrivateDnsZonesVirtualNetworkLinkOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesVirtualNetworkLinkOperatorSpec
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

// Generator of PrivateDnsZonesVirtualNetworkLinkOperatorSpec instances for property testing - lazily instantiated by
// PrivateDnsZonesVirtualNetworkLinkOperatorSpecGenerator()
var privateDnsZonesVirtualNetworkLinkOperatorSpecGenerator gopter.Gen

// PrivateDnsZonesVirtualNetworkLinkOperatorSpecGenerator returns a generator of PrivateDnsZonesVirtualNetworkLinkOperatorSpec instances for property testing.
func PrivateDnsZonesVirtualNetworkLinkOperatorSpecGenerator() gopter.Gen {
	if privateDnsZonesVirtualNetworkLinkOperatorSpecGenerator != nil {
		return privateDnsZonesVirtualNetworkLinkOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	privateDnsZonesVirtualNetworkLinkOperatorSpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesVirtualNetworkLinkOperatorSpec{}), generators)

	return privateDnsZonesVirtualNetworkLinkOperatorSpecGenerator
}

func Test_PrivateDnsZonesVirtualNetworkLink_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesVirtualNetworkLink_STATUS to PrivateDnsZonesVirtualNetworkLink_STATUS via AssignProperties_To_PrivateDnsZonesVirtualNetworkLink_STATUS & AssignProperties_From_PrivateDnsZonesVirtualNetworkLink_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesVirtualNetworkLink_STATUS, PrivateDnsZonesVirtualNetworkLink_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesVirtualNetworkLink_STATUS tests if a specific instance of PrivateDnsZonesVirtualNetworkLink_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesVirtualNetworkLink_STATUS(subject PrivateDnsZonesVirtualNetworkLink_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateDnsZonesVirtualNetworkLink_STATUS
	err := copied.AssignProperties_To_PrivateDnsZonesVirtualNetworkLink_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesVirtualNetworkLink_STATUS
	err = actual.AssignProperties_From_PrivateDnsZonesVirtualNetworkLink_STATUS(&other)
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

func Test_PrivateDnsZonesVirtualNetworkLink_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesVirtualNetworkLink_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink_STATUS, PrivateDnsZonesVirtualNetworkLink_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink_STATUS runs a test to see if a specific instance of PrivateDnsZonesVirtualNetworkLink_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink_STATUS(subject PrivateDnsZonesVirtualNetworkLink_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesVirtualNetworkLink_STATUS
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

// Generator of PrivateDnsZonesVirtualNetworkLink_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZonesVirtualNetworkLink_STATUSGenerator()
var privateDnsZonesVirtualNetworkLink_STATUSGenerator gopter.Gen

// PrivateDnsZonesVirtualNetworkLink_STATUSGenerator returns a generator of PrivateDnsZonesVirtualNetworkLink_STATUS instances for property testing.
// We first initialize privateDnsZonesVirtualNetworkLink_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonesVirtualNetworkLink_STATUSGenerator() gopter.Gen {
	if privateDnsZonesVirtualNetworkLink_STATUSGenerator != nil {
		return privateDnsZonesVirtualNetworkLink_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_STATUS(generators)
	privateDnsZonesVirtualNetworkLink_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesVirtualNetworkLink_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_STATUS(generators)
	privateDnsZonesVirtualNetworkLink_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesVirtualNetworkLink_STATUS{}), generators)

	return privateDnsZonesVirtualNetworkLink_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		VirtualNetworkLinkProperties_ProvisioningState_STATUS_Canceled,
		VirtualNetworkLinkProperties_ProvisioningState_STATUS_Creating,
		VirtualNetworkLinkProperties_ProvisioningState_STATUS_Deleting,
		VirtualNetworkLinkProperties_ProvisioningState_STATUS_Failed,
		VirtualNetworkLinkProperties_ProvisioningState_STATUS_Succeeded,
		VirtualNetworkLinkProperties_ProvisioningState_STATUS_Updating))
	gens["RegistrationEnabled"] = gen.PtrOf(gen.Bool())
	gens["ResolutionPolicy"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkLinkProperties_ResolutionPolicy_STATUS_Default, VirtualNetworkLinkProperties_ResolutionPolicy_STATUS_NxDomainRedirect))
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["VirtualNetworkLinkState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkLinkProperties_VirtualNetworkLinkState_STATUS_Completed, VirtualNetworkLinkProperties_VirtualNetworkLinkState_STATUS_InProgress))
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_STATUS(gens map[string]gopter.Gen) {
	gens["VirtualNetwork"] = gen.PtrOf(SubResource_STATUSGenerator())
}

func Test_PrivateDnsZonesVirtualNetworkLink_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesVirtualNetworkLink_Spec to PrivateDnsZonesVirtualNetworkLink_Spec via AssignProperties_To_PrivateDnsZonesVirtualNetworkLink_Spec & AssignProperties_From_PrivateDnsZonesVirtualNetworkLink_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesVirtualNetworkLink_Spec, PrivateDnsZonesVirtualNetworkLink_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesVirtualNetworkLink_Spec tests if a specific instance of PrivateDnsZonesVirtualNetworkLink_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesVirtualNetworkLink_Spec(subject PrivateDnsZonesVirtualNetworkLink_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateDnsZonesVirtualNetworkLink_Spec
	err := copied.AssignProperties_To_PrivateDnsZonesVirtualNetworkLink_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesVirtualNetworkLink_Spec
	err = actual.AssignProperties_From_PrivateDnsZonesVirtualNetworkLink_Spec(&other)
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

func Test_PrivateDnsZonesVirtualNetworkLink_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesVirtualNetworkLink_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink_Spec, PrivateDnsZonesVirtualNetworkLink_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink_Spec runs a test to see if a specific instance of PrivateDnsZonesVirtualNetworkLink_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink_Spec(subject PrivateDnsZonesVirtualNetworkLink_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesVirtualNetworkLink_Spec
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

// Generator of PrivateDnsZonesVirtualNetworkLink_Spec instances for property testing - lazily instantiated by
// PrivateDnsZonesVirtualNetworkLink_SpecGenerator()
var privateDnsZonesVirtualNetworkLink_SpecGenerator gopter.Gen

// PrivateDnsZonesVirtualNetworkLink_SpecGenerator returns a generator of PrivateDnsZonesVirtualNetworkLink_Spec instances for property testing.
// We first initialize privateDnsZonesVirtualNetworkLink_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonesVirtualNetworkLink_SpecGenerator() gopter.Gen {
	if privateDnsZonesVirtualNetworkLink_SpecGenerator != nil {
		return privateDnsZonesVirtualNetworkLink_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_Spec(generators)
	privateDnsZonesVirtualNetworkLink_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesVirtualNetworkLink_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_Spec(generators)
	privateDnsZonesVirtualNetworkLink_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesVirtualNetworkLink_Spec{}), generators)

	return privateDnsZonesVirtualNetworkLink_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["RegistrationEnabled"] = gen.PtrOf(gen.Bool())
	gens["ResolutionPolicy"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkLinkProperties_ResolutionPolicy_Default, VirtualNetworkLinkProperties_ResolutionPolicy_NxDomainRedirect))
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(PrivateDnsZonesVirtualNetworkLinkOperatorSpecGenerator())
	gens["VirtualNetwork"] = gen.PtrOf(SubResourceGenerator())
}

func Test_SubResource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SubResource to SubResource via AssignProperties_To_SubResource & AssignProperties_From_SubResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSubResource, SubResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSubResource tests if a specific instance of SubResource can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSubResource(subject SubResource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SubResource
	err := copied.AssignProperties_To_SubResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SubResource
	err = actual.AssignProperties_From_SubResource(&other)
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

func Test_SubResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResource, SubResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResource runs a test to see if a specific instance of SubResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResource(subject SubResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource
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

// Generator of SubResource instances for property testing - lazily instantiated by SubResourceGenerator()
var subResourceGenerator gopter.Gen

// SubResourceGenerator returns a generator of SubResource instances for property testing.
func SubResourceGenerator() gopter.Gen {
	if subResourceGenerator != nil {
		return subResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	subResourceGenerator = gen.Struct(reflect.TypeOf(SubResource{}), generators)

	return subResourceGenerator
}

func Test_SubResource_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SubResource_STATUS to SubResource_STATUS via AssignProperties_To_SubResource_STATUS & AssignProperties_From_SubResource_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSubResource_STATUS, SubResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSubResource_STATUS tests if a specific instance of SubResource_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSubResource_STATUS(subject SubResource_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SubResource_STATUS
	err := copied.AssignProperties_To_SubResource_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SubResource_STATUS
	err = actual.AssignProperties_From_SubResource_STATUS(&other)
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

func Test_SubResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResource_STATUS, SubResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResource_STATUS runs a test to see if a specific instance of SubResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResource_STATUS(subject SubResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource_STATUS
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

// Generator of SubResource_STATUS instances for property testing - lazily instantiated by SubResource_STATUSGenerator()
var subResource_STATUSGenerator gopter.Gen

// SubResource_STATUSGenerator returns a generator of SubResource_STATUS instances for property testing.
func SubResource_STATUSGenerator() gopter.Gen {
	if subResource_STATUSGenerator != nil {
		return subResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubResource_STATUS(generators)
	subResource_STATUSGenerator = gen.Struct(reflect.TypeOf(SubResource_STATUS{}), generators)

	return subResource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSubResource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResource_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
