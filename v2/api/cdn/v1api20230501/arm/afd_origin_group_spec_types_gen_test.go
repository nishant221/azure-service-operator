// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_AFDOriginGroupProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AFDOriginGroupProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAFDOriginGroupProperties, AFDOriginGroupPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAFDOriginGroupProperties runs a test to see if a specific instance of AFDOriginGroupProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForAFDOriginGroupProperties(subject AFDOriginGroupProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AFDOriginGroupProperties
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

// Generator of AFDOriginGroupProperties instances for property testing - lazily instantiated by
// AFDOriginGroupPropertiesGenerator()
var afdOriginGroupPropertiesGenerator gopter.Gen

// AFDOriginGroupPropertiesGenerator returns a generator of AFDOriginGroupProperties instances for property testing.
// We first initialize afdOriginGroupPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AFDOriginGroupPropertiesGenerator() gopter.Gen {
	if afdOriginGroupPropertiesGenerator != nil {
		return afdOriginGroupPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAFDOriginGroupProperties(generators)
	afdOriginGroupPropertiesGenerator = gen.Struct(reflect.TypeOf(AFDOriginGroupProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAFDOriginGroupProperties(generators)
	AddRelatedPropertyGeneratorsForAFDOriginGroupProperties(generators)
	afdOriginGroupPropertiesGenerator = gen.Struct(reflect.TypeOf(AFDOriginGroupProperties{}), generators)

	return afdOriginGroupPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForAFDOriginGroupProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAFDOriginGroupProperties(gens map[string]gopter.Gen) {
	gens["SessionAffinityState"] = gen.PtrOf(gen.OneConstOf(AFDOriginGroupProperties_SessionAffinityState_Disabled, AFDOriginGroupProperties_SessionAffinityState_Enabled))
	gens["TrafficRestorationTimeToHealedOrNewEndpointsInMinutes"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForAFDOriginGroupProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAFDOriginGroupProperties(gens map[string]gopter.Gen) {
	gens["HealthProbeSettings"] = gen.PtrOf(HealthProbeParametersGenerator())
	gens["LoadBalancingSettings"] = gen.PtrOf(LoadBalancingSettingsParametersGenerator())
}

func Test_AfdOriginGroup_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AfdOriginGroup_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAfdOriginGroup_Spec, AfdOriginGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAfdOriginGroup_Spec runs a test to see if a specific instance of AfdOriginGroup_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForAfdOriginGroup_Spec(subject AfdOriginGroup_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AfdOriginGroup_Spec
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

// Generator of AfdOriginGroup_Spec instances for property testing - lazily instantiated by
// AfdOriginGroup_SpecGenerator()
var afdOriginGroup_SpecGenerator gopter.Gen

// AfdOriginGroup_SpecGenerator returns a generator of AfdOriginGroup_Spec instances for property testing.
// We first initialize afdOriginGroup_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AfdOriginGroup_SpecGenerator() gopter.Gen {
	if afdOriginGroup_SpecGenerator != nil {
		return afdOriginGroup_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdOriginGroup_Spec(generators)
	afdOriginGroup_SpecGenerator = gen.Struct(reflect.TypeOf(AfdOriginGroup_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdOriginGroup_Spec(generators)
	AddRelatedPropertyGeneratorsForAfdOriginGroup_Spec(generators)
	afdOriginGroup_SpecGenerator = gen.Struct(reflect.TypeOf(AfdOriginGroup_Spec{}), generators)

	return afdOriginGroup_SpecGenerator
}

// AddIndependentPropertyGeneratorsForAfdOriginGroup_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAfdOriginGroup_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForAfdOriginGroup_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAfdOriginGroup_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AFDOriginGroupPropertiesGenerator())
}

func Test_HealthProbeParameters_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HealthProbeParameters via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHealthProbeParameters, HealthProbeParametersGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHealthProbeParameters runs a test to see if a specific instance of HealthProbeParameters round trips to JSON and back losslessly
func RunJSONSerializationTestForHealthProbeParameters(subject HealthProbeParameters) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HealthProbeParameters
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

// Generator of HealthProbeParameters instances for property testing - lazily instantiated by
// HealthProbeParametersGenerator()
var healthProbeParametersGenerator gopter.Gen

// HealthProbeParametersGenerator returns a generator of HealthProbeParameters instances for property testing.
func HealthProbeParametersGenerator() gopter.Gen {
	if healthProbeParametersGenerator != nil {
		return healthProbeParametersGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHealthProbeParameters(generators)
	healthProbeParametersGenerator = gen.Struct(reflect.TypeOf(HealthProbeParameters{}), generators)

	return healthProbeParametersGenerator
}

// AddIndependentPropertyGeneratorsForHealthProbeParameters is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHealthProbeParameters(gens map[string]gopter.Gen) {
	gens["ProbeIntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["ProbePath"] = gen.PtrOf(gen.AlphaString())
	gens["ProbeProtocol"] = gen.PtrOf(gen.OneConstOf(HealthProbeParameters_ProbeProtocol_Http, HealthProbeParameters_ProbeProtocol_Https, HealthProbeParameters_ProbeProtocol_NotSet))
	gens["ProbeRequestType"] = gen.PtrOf(gen.OneConstOf(HealthProbeParameters_ProbeRequestType_GET, HealthProbeParameters_ProbeRequestType_HEAD, HealthProbeParameters_ProbeRequestType_NotSet))
}

func Test_LoadBalancingSettingsParameters_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LoadBalancingSettingsParameters via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLoadBalancingSettingsParameters, LoadBalancingSettingsParametersGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLoadBalancingSettingsParameters runs a test to see if a specific instance of LoadBalancingSettingsParameters round trips to JSON and back losslessly
func RunJSONSerializationTestForLoadBalancingSettingsParameters(subject LoadBalancingSettingsParameters) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LoadBalancingSettingsParameters
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

// Generator of LoadBalancingSettingsParameters instances for property testing - lazily instantiated by
// LoadBalancingSettingsParametersGenerator()
var loadBalancingSettingsParametersGenerator gopter.Gen

// LoadBalancingSettingsParametersGenerator returns a generator of LoadBalancingSettingsParameters instances for property testing.
func LoadBalancingSettingsParametersGenerator() gopter.Gen {
	if loadBalancingSettingsParametersGenerator != nil {
		return loadBalancingSettingsParametersGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLoadBalancingSettingsParameters(generators)
	loadBalancingSettingsParametersGenerator = gen.Struct(reflect.TypeOf(LoadBalancingSettingsParameters{}), generators)

	return loadBalancingSettingsParametersGenerator
}

// AddIndependentPropertyGeneratorsForLoadBalancingSettingsParameters is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLoadBalancingSettingsParameters(gens map[string]gopter.Gen) {
	gens["AdditionalLatencyInMilliseconds"] = gen.PtrOf(gen.Int())
	gens["SampleSize"] = gen.PtrOf(gen.Int())
	gens["SuccessfulSamplesRequired"] = gen.PtrOf(gen.Int())
}