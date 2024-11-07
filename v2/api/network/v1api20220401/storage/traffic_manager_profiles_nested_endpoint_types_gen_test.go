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

func Test_TrafficManagerProfilesNestedEndpoint_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrafficManagerProfilesNestedEndpoint via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrafficManagerProfilesNestedEndpoint, TrafficManagerProfilesNestedEndpointGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrafficManagerProfilesNestedEndpoint runs a test to see if a specific instance of TrafficManagerProfilesNestedEndpoint round trips to JSON and back losslessly
func RunJSONSerializationTestForTrafficManagerProfilesNestedEndpoint(subject TrafficManagerProfilesNestedEndpoint) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrafficManagerProfilesNestedEndpoint
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

// Generator of TrafficManagerProfilesNestedEndpoint instances for property testing - lazily instantiated by
// TrafficManagerProfilesNestedEndpointGenerator()
var trafficManagerProfilesNestedEndpointGenerator gopter.Gen

// TrafficManagerProfilesNestedEndpointGenerator returns a generator of TrafficManagerProfilesNestedEndpoint instances for property testing.
func TrafficManagerProfilesNestedEndpointGenerator() gopter.Gen {
	if trafficManagerProfilesNestedEndpointGenerator != nil {
		return trafficManagerProfilesNestedEndpointGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint(generators)
	trafficManagerProfilesNestedEndpointGenerator = gen.Struct(reflect.TypeOf(TrafficManagerProfilesNestedEndpoint{}), generators)

	return trafficManagerProfilesNestedEndpointGenerator
}

// AddRelatedPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint(gens map[string]gopter.Gen) {
	gens["Spec"] = TrafficManagerProfilesNestedEndpoint_SpecGenerator()
	gens["Status"] = TrafficManagerProfilesNestedEndpoint_STATUSGenerator()
}

func Test_TrafficManagerProfilesNestedEndpointOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrafficManagerProfilesNestedEndpointOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrafficManagerProfilesNestedEndpointOperatorSpec, TrafficManagerProfilesNestedEndpointOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrafficManagerProfilesNestedEndpointOperatorSpec runs a test to see if a specific instance of TrafficManagerProfilesNestedEndpointOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForTrafficManagerProfilesNestedEndpointOperatorSpec(subject TrafficManagerProfilesNestedEndpointOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrafficManagerProfilesNestedEndpointOperatorSpec
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

// Generator of TrafficManagerProfilesNestedEndpointOperatorSpec instances for property testing - lazily instantiated by
// TrafficManagerProfilesNestedEndpointOperatorSpecGenerator()
var trafficManagerProfilesNestedEndpointOperatorSpecGenerator gopter.Gen

// TrafficManagerProfilesNestedEndpointOperatorSpecGenerator returns a generator of TrafficManagerProfilesNestedEndpointOperatorSpec instances for property testing.
func TrafficManagerProfilesNestedEndpointOperatorSpecGenerator() gopter.Gen {
	if trafficManagerProfilesNestedEndpointOperatorSpecGenerator != nil {
		return trafficManagerProfilesNestedEndpointOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	trafficManagerProfilesNestedEndpointOperatorSpecGenerator = gen.Struct(reflect.TypeOf(TrafficManagerProfilesNestedEndpointOperatorSpec{}), generators)

	return trafficManagerProfilesNestedEndpointOperatorSpecGenerator
}

func Test_TrafficManagerProfilesNestedEndpoint_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrafficManagerProfilesNestedEndpoint_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrafficManagerProfilesNestedEndpoint_STATUS, TrafficManagerProfilesNestedEndpoint_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrafficManagerProfilesNestedEndpoint_STATUS runs a test to see if a specific instance of TrafficManagerProfilesNestedEndpoint_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTrafficManagerProfilesNestedEndpoint_STATUS(subject TrafficManagerProfilesNestedEndpoint_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrafficManagerProfilesNestedEndpoint_STATUS
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

// Generator of TrafficManagerProfilesNestedEndpoint_STATUS instances for property testing - lazily instantiated by
// TrafficManagerProfilesNestedEndpoint_STATUSGenerator()
var trafficManagerProfilesNestedEndpoint_STATUSGenerator gopter.Gen

// TrafficManagerProfilesNestedEndpoint_STATUSGenerator returns a generator of TrafficManagerProfilesNestedEndpoint_STATUS instances for property testing.
// We first initialize trafficManagerProfilesNestedEndpoint_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TrafficManagerProfilesNestedEndpoint_STATUSGenerator() gopter.Gen {
	if trafficManagerProfilesNestedEndpoint_STATUSGenerator != nil {
		return trafficManagerProfilesNestedEndpoint_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint_STATUS(generators)
	trafficManagerProfilesNestedEndpoint_STATUSGenerator = gen.Struct(reflect.TypeOf(TrafficManagerProfilesNestedEndpoint_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint_STATUS(generators)
	AddRelatedPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint_STATUS(generators)
	trafficManagerProfilesNestedEndpoint_STATUSGenerator = gen.Struct(reflect.TypeOf(TrafficManagerProfilesNestedEndpoint_STATUS{}), generators)

	return trafficManagerProfilesNestedEndpoint_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint_STATUS(gens map[string]gopter.Gen) {
	gens["AlwaysServe"] = gen.PtrOf(gen.AlphaString())
	gens["EndpointLocation"] = gen.PtrOf(gen.AlphaString())
	gens["EndpointMonitorStatus"] = gen.PtrOf(gen.AlphaString())
	gens["EndpointStatus"] = gen.PtrOf(gen.AlphaString())
	gens["GeoMapping"] = gen.SliceOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["MinChildEndpoints"] = gen.PtrOf(gen.Int())
	gens["MinChildEndpointsIPv4"] = gen.PtrOf(gen.Int())
	gens["MinChildEndpointsIPv6"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
	gens["TargetResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Weight"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint_STATUS(gens map[string]gopter.Gen) {
	gens["CustomHeaders"] = gen.SliceOf(EndpointProperties_CustomHeaders_STATUSGenerator())
	gens["Subnets"] = gen.SliceOf(EndpointProperties_Subnets_STATUSGenerator())
}

func Test_TrafficManagerProfilesNestedEndpoint_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrafficManagerProfilesNestedEndpoint_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrafficManagerProfilesNestedEndpoint_Spec, TrafficManagerProfilesNestedEndpoint_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrafficManagerProfilesNestedEndpoint_Spec runs a test to see if a specific instance of TrafficManagerProfilesNestedEndpoint_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForTrafficManagerProfilesNestedEndpoint_Spec(subject TrafficManagerProfilesNestedEndpoint_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrafficManagerProfilesNestedEndpoint_Spec
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

// Generator of TrafficManagerProfilesNestedEndpoint_Spec instances for property testing - lazily instantiated by
// TrafficManagerProfilesNestedEndpoint_SpecGenerator()
var trafficManagerProfilesNestedEndpoint_SpecGenerator gopter.Gen

// TrafficManagerProfilesNestedEndpoint_SpecGenerator returns a generator of TrafficManagerProfilesNestedEndpoint_Spec instances for property testing.
// We first initialize trafficManagerProfilesNestedEndpoint_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TrafficManagerProfilesNestedEndpoint_SpecGenerator() gopter.Gen {
	if trafficManagerProfilesNestedEndpoint_SpecGenerator != nil {
		return trafficManagerProfilesNestedEndpoint_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint_Spec(generators)
	trafficManagerProfilesNestedEndpoint_SpecGenerator = gen.Struct(reflect.TypeOf(TrafficManagerProfilesNestedEndpoint_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint_Spec(generators)
	AddRelatedPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint_Spec(generators)
	trafficManagerProfilesNestedEndpoint_SpecGenerator = gen.Struct(reflect.TypeOf(TrafficManagerProfilesNestedEndpoint_Spec{}), generators)

	return trafficManagerProfilesNestedEndpoint_SpecGenerator
}

// AddIndependentPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint_Spec(gens map[string]gopter.Gen) {
	gens["AlwaysServe"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["EndpointLocation"] = gen.PtrOf(gen.AlphaString())
	gens["EndpointMonitorStatus"] = gen.PtrOf(gen.AlphaString())
	gens["EndpointStatus"] = gen.PtrOf(gen.AlphaString())
	gens["GeoMapping"] = gen.SliceOf(gen.AlphaString())
	gens["MinChildEndpoints"] = gen.PtrOf(gen.Int())
	gens["MinChildEndpointsIPv4"] = gen.PtrOf(gen.Int())
	gens["MinChildEndpointsIPv6"] = gen.PtrOf(gen.Int())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Weight"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint_Spec(gens map[string]gopter.Gen) {
	gens["CustomHeaders"] = gen.SliceOf(EndpointProperties_CustomHeadersGenerator())
	gens["OperatorSpec"] = gen.PtrOf(TrafficManagerProfilesNestedEndpointOperatorSpecGenerator())
	gens["Subnets"] = gen.SliceOf(EndpointProperties_SubnetsGenerator())
}
