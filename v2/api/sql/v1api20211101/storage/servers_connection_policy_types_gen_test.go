// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

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

func Test_ServersConnectionPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersConnectionPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersConnectionPolicy, ServersConnectionPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersConnectionPolicy runs a test to see if a specific instance of ServersConnectionPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForServersConnectionPolicy(subject ServersConnectionPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersConnectionPolicy
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

// Generator of ServersConnectionPolicy instances for property testing - lazily instantiated by
// ServersConnectionPolicyGenerator()
var serversConnectionPolicyGenerator gopter.Gen

// ServersConnectionPolicyGenerator returns a generator of ServersConnectionPolicy instances for property testing.
func ServersConnectionPolicyGenerator() gopter.Gen {
	if serversConnectionPolicyGenerator != nil {
		return serversConnectionPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersConnectionPolicy(generators)
	serversConnectionPolicyGenerator = gen.Struct(reflect.TypeOf(ServersConnectionPolicy{}), generators)

	return serversConnectionPolicyGenerator
}

// AddRelatedPropertyGeneratorsForServersConnectionPolicy is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersConnectionPolicy(gens map[string]gopter.Gen) {
	gens["Spec"] = ServersConnectionPolicy_SpecGenerator()
	gens["Status"] = ServersConnectionPolicy_STATUSGenerator()
}

func Test_ServersConnectionPolicyOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersConnectionPolicyOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersConnectionPolicyOperatorSpec, ServersConnectionPolicyOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersConnectionPolicyOperatorSpec runs a test to see if a specific instance of ServersConnectionPolicyOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersConnectionPolicyOperatorSpec(subject ServersConnectionPolicyOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersConnectionPolicyOperatorSpec
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

// Generator of ServersConnectionPolicyOperatorSpec instances for property testing - lazily instantiated by
// ServersConnectionPolicyOperatorSpecGenerator()
var serversConnectionPolicyOperatorSpecGenerator gopter.Gen

// ServersConnectionPolicyOperatorSpecGenerator returns a generator of ServersConnectionPolicyOperatorSpec instances for property testing.
func ServersConnectionPolicyOperatorSpecGenerator() gopter.Gen {
	if serversConnectionPolicyOperatorSpecGenerator != nil {
		return serversConnectionPolicyOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	serversConnectionPolicyOperatorSpecGenerator = gen.Struct(reflect.TypeOf(ServersConnectionPolicyOperatorSpec{}), generators)

	return serversConnectionPolicyOperatorSpecGenerator
}

func Test_ServersConnectionPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersConnectionPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersConnectionPolicy_STATUS, ServersConnectionPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersConnectionPolicy_STATUS runs a test to see if a specific instance of ServersConnectionPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServersConnectionPolicy_STATUS(subject ServersConnectionPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersConnectionPolicy_STATUS
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

// Generator of ServersConnectionPolicy_STATUS instances for property testing - lazily instantiated by
// ServersConnectionPolicy_STATUSGenerator()
var serversConnectionPolicy_STATUSGenerator gopter.Gen

// ServersConnectionPolicy_STATUSGenerator returns a generator of ServersConnectionPolicy_STATUS instances for property testing.
func ServersConnectionPolicy_STATUSGenerator() gopter.Gen {
	if serversConnectionPolicy_STATUSGenerator != nil {
		return serversConnectionPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersConnectionPolicy_STATUS(generators)
	serversConnectionPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(ServersConnectionPolicy_STATUS{}), generators)

	return serversConnectionPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServersConnectionPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersConnectionPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["ConnectionType"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServersConnectionPolicy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersConnectionPolicy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersConnectionPolicy_Spec, ServersConnectionPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersConnectionPolicy_Spec runs a test to see if a specific instance of ServersConnectionPolicy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersConnectionPolicy_Spec(subject ServersConnectionPolicy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersConnectionPolicy_Spec
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

// Generator of ServersConnectionPolicy_Spec instances for property testing - lazily instantiated by
// ServersConnectionPolicy_SpecGenerator()
var serversConnectionPolicy_SpecGenerator gopter.Gen

// ServersConnectionPolicy_SpecGenerator returns a generator of ServersConnectionPolicy_Spec instances for property testing.
// We first initialize serversConnectionPolicy_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServersConnectionPolicy_SpecGenerator() gopter.Gen {
	if serversConnectionPolicy_SpecGenerator != nil {
		return serversConnectionPolicy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersConnectionPolicy_Spec(generators)
	serversConnectionPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(ServersConnectionPolicy_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersConnectionPolicy_Spec(generators)
	AddRelatedPropertyGeneratorsForServersConnectionPolicy_Spec(generators)
	serversConnectionPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(ServersConnectionPolicy_Spec{}), generators)

	return serversConnectionPolicy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServersConnectionPolicy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersConnectionPolicy_Spec(gens map[string]gopter.Gen) {
	gens["ConnectionType"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForServersConnectionPolicy_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersConnectionPolicy_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(ServersConnectionPolicyOperatorSpecGenerator())
}
