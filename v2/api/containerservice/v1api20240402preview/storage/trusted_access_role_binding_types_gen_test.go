// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"encoding/json"
	v20231001s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/storage"
	v20240901s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240901/storage"
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

func Test_TrustedAccessRoleBinding_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from TrustedAccessRoleBinding to hub returns original",
		prop.ForAll(RunResourceConversionTestForTrustedAccessRoleBinding, TrustedAccessRoleBindingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForTrustedAccessRoleBinding tests if a specific instance of TrustedAccessRoleBinding round trips to the hub storage version and back losslessly
func RunResourceConversionTestForTrustedAccessRoleBinding(subject TrustedAccessRoleBinding) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20240901s.TrustedAccessRoleBinding
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual TrustedAccessRoleBinding
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

func Test_TrustedAccessRoleBinding_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from TrustedAccessRoleBinding to TrustedAccessRoleBinding via AssignProperties_To_TrustedAccessRoleBinding & AssignProperties_From_TrustedAccessRoleBinding returns original",
		prop.ForAll(RunPropertyAssignmentTestForTrustedAccessRoleBinding, TrustedAccessRoleBindingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForTrustedAccessRoleBinding tests if a specific instance of TrustedAccessRoleBinding can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForTrustedAccessRoleBinding(subject TrustedAccessRoleBinding) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20231001s.TrustedAccessRoleBinding
	err := copied.AssignProperties_To_TrustedAccessRoleBinding(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual TrustedAccessRoleBinding
	err = actual.AssignProperties_From_TrustedAccessRoleBinding(&other)
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

func Test_TrustedAccessRoleBinding_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrustedAccessRoleBinding via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrustedAccessRoleBinding, TrustedAccessRoleBindingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrustedAccessRoleBinding runs a test to see if a specific instance of TrustedAccessRoleBinding round trips to JSON and back losslessly
func RunJSONSerializationTestForTrustedAccessRoleBinding(subject TrustedAccessRoleBinding) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrustedAccessRoleBinding
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

// Generator of TrustedAccessRoleBinding instances for property testing - lazily instantiated by
// TrustedAccessRoleBindingGenerator()
var trustedAccessRoleBindingGenerator gopter.Gen

// TrustedAccessRoleBindingGenerator returns a generator of TrustedAccessRoleBinding instances for property testing.
func TrustedAccessRoleBindingGenerator() gopter.Gen {
	if trustedAccessRoleBindingGenerator != nil {
		return trustedAccessRoleBindingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding(generators)
	trustedAccessRoleBindingGenerator = gen.Struct(reflect.TypeOf(TrustedAccessRoleBinding{}), generators)

	return trustedAccessRoleBindingGenerator
}

// AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding(gens map[string]gopter.Gen) {
	gens["Spec"] = TrustedAccessRoleBinding_SpecGenerator()
	gens["Status"] = TrustedAccessRoleBinding_STATUSGenerator()
}

func Test_TrustedAccessRoleBindingOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from TrustedAccessRoleBindingOperatorSpec to TrustedAccessRoleBindingOperatorSpec via AssignProperties_To_TrustedAccessRoleBindingOperatorSpec & AssignProperties_From_TrustedAccessRoleBindingOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForTrustedAccessRoleBindingOperatorSpec, TrustedAccessRoleBindingOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForTrustedAccessRoleBindingOperatorSpec tests if a specific instance of TrustedAccessRoleBindingOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForTrustedAccessRoleBindingOperatorSpec(subject TrustedAccessRoleBindingOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20231001s.TrustedAccessRoleBindingOperatorSpec
	err := copied.AssignProperties_To_TrustedAccessRoleBindingOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual TrustedAccessRoleBindingOperatorSpec
	err = actual.AssignProperties_From_TrustedAccessRoleBindingOperatorSpec(&other)
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

func Test_TrustedAccessRoleBindingOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrustedAccessRoleBindingOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrustedAccessRoleBindingOperatorSpec, TrustedAccessRoleBindingOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrustedAccessRoleBindingOperatorSpec runs a test to see if a specific instance of TrustedAccessRoleBindingOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForTrustedAccessRoleBindingOperatorSpec(subject TrustedAccessRoleBindingOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrustedAccessRoleBindingOperatorSpec
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

// Generator of TrustedAccessRoleBindingOperatorSpec instances for property testing - lazily instantiated by
// TrustedAccessRoleBindingOperatorSpecGenerator()
var trustedAccessRoleBindingOperatorSpecGenerator gopter.Gen

// TrustedAccessRoleBindingOperatorSpecGenerator returns a generator of TrustedAccessRoleBindingOperatorSpec instances for property testing.
func TrustedAccessRoleBindingOperatorSpecGenerator() gopter.Gen {
	if trustedAccessRoleBindingOperatorSpecGenerator != nil {
		return trustedAccessRoleBindingOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	trustedAccessRoleBindingOperatorSpecGenerator = gen.Struct(reflect.TypeOf(TrustedAccessRoleBindingOperatorSpec{}), generators)

	return trustedAccessRoleBindingOperatorSpecGenerator
}

func Test_TrustedAccessRoleBinding_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from TrustedAccessRoleBinding_STATUS to TrustedAccessRoleBinding_STATUS via AssignProperties_To_TrustedAccessRoleBinding_STATUS & AssignProperties_From_TrustedAccessRoleBinding_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForTrustedAccessRoleBinding_STATUS, TrustedAccessRoleBinding_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForTrustedAccessRoleBinding_STATUS tests if a specific instance of TrustedAccessRoleBinding_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForTrustedAccessRoleBinding_STATUS(subject TrustedAccessRoleBinding_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20231001s.TrustedAccessRoleBinding_STATUS
	err := copied.AssignProperties_To_TrustedAccessRoleBinding_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual TrustedAccessRoleBinding_STATUS
	err = actual.AssignProperties_From_TrustedAccessRoleBinding_STATUS(&other)
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

func Test_TrustedAccessRoleBinding_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrustedAccessRoleBinding_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrustedAccessRoleBinding_STATUS, TrustedAccessRoleBinding_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrustedAccessRoleBinding_STATUS runs a test to see if a specific instance of TrustedAccessRoleBinding_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTrustedAccessRoleBinding_STATUS(subject TrustedAccessRoleBinding_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrustedAccessRoleBinding_STATUS
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

// Generator of TrustedAccessRoleBinding_STATUS instances for property testing - lazily instantiated by
// TrustedAccessRoleBinding_STATUSGenerator()
var trustedAccessRoleBinding_STATUSGenerator gopter.Gen

// TrustedAccessRoleBinding_STATUSGenerator returns a generator of TrustedAccessRoleBinding_STATUS instances for property testing.
// We first initialize trustedAccessRoleBinding_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TrustedAccessRoleBinding_STATUSGenerator() gopter.Gen {
	if trustedAccessRoleBinding_STATUSGenerator != nil {
		return trustedAccessRoleBinding_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrustedAccessRoleBinding_STATUS(generators)
	trustedAccessRoleBinding_STATUSGenerator = gen.Struct(reflect.TypeOf(TrustedAccessRoleBinding_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrustedAccessRoleBinding_STATUS(generators)
	AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding_STATUS(generators)
	trustedAccessRoleBinding_STATUSGenerator = gen.Struct(reflect.TypeOf(TrustedAccessRoleBinding_STATUS{}), generators)

	return trustedAccessRoleBinding_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTrustedAccessRoleBinding_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTrustedAccessRoleBinding_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Roles"] = gen.SliceOf(gen.AlphaString())
	gens["SourceResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_TrustedAccessRoleBinding_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from TrustedAccessRoleBinding_Spec to TrustedAccessRoleBinding_Spec via AssignProperties_To_TrustedAccessRoleBinding_Spec & AssignProperties_From_TrustedAccessRoleBinding_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForTrustedAccessRoleBinding_Spec, TrustedAccessRoleBinding_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForTrustedAccessRoleBinding_Spec tests if a specific instance of TrustedAccessRoleBinding_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForTrustedAccessRoleBinding_Spec(subject TrustedAccessRoleBinding_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20231001s.TrustedAccessRoleBinding_Spec
	err := copied.AssignProperties_To_TrustedAccessRoleBinding_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual TrustedAccessRoleBinding_Spec
	err = actual.AssignProperties_From_TrustedAccessRoleBinding_Spec(&other)
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

func Test_TrustedAccessRoleBinding_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrustedAccessRoleBinding_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrustedAccessRoleBinding_Spec, TrustedAccessRoleBinding_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrustedAccessRoleBinding_Spec runs a test to see if a specific instance of TrustedAccessRoleBinding_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForTrustedAccessRoleBinding_Spec(subject TrustedAccessRoleBinding_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrustedAccessRoleBinding_Spec
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

// Generator of TrustedAccessRoleBinding_Spec instances for property testing - lazily instantiated by
// TrustedAccessRoleBinding_SpecGenerator()
var trustedAccessRoleBinding_SpecGenerator gopter.Gen

// TrustedAccessRoleBinding_SpecGenerator returns a generator of TrustedAccessRoleBinding_Spec instances for property testing.
// We first initialize trustedAccessRoleBinding_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TrustedAccessRoleBinding_SpecGenerator() gopter.Gen {
	if trustedAccessRoleBinding_SpecGenerator != nil {
		return trustedAccessRoleBinding_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrustedAccessRoleBinding_Spec(generators)
	trustedAccessRoleBinding_SpecGenerator = gen.Struct(reflect.TypeOf(TrustedAccessRoleBinding_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrustedAccessRoleBinding_Spec(generators)
	AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding_Spec(generators)
	trustedAccessRoleBinding_SpecGenerator = gen.Struct(reflect.TypeOf(TrustedAccessRoleBinding_Spec{}), generators)

	return trustedAccessRoleBinding_SpecGenerator
}

// AddIndependentPropertyGeneratorsForTrustedAccessRoleBinding_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTrustedAccessRoleBinding_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Roles"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTrustedAccessRoleBinding_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(TrustedAccessRoleBindingOperatorSpecGenerator())
}
