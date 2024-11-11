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

func Test_BastionHost_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BastionHost via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBastionHost, BastionHostGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBastionHost runs a test to see if a specific instance of BastionHost round trips to JSON and back losslessly
func RunJSONSerializationTestForBastionHost(subject BastionHost) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BastionHost
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

// Generator of BastionHost instances for property testing - lazily instantiated by BastionHostGenerator()
var bastionHostGenerator gopter.Gen

// BastionHostGenerator returns a generator of BastionHost instances for property testing.
func BastionHostGenerator() gopter.Gen {
	if bastionHostGenerator != nil {
		return bastionHostGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForBastionHost(generators)
	bastionHostGenerator = gen.Struct(reflect.TypeOf(BastionHost{}), generators)

	return bastionHostGenerator
}

// AddRelatedPropertyGeneratorsForBastionHost is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBastionHost(gens map[string]gopter.Gen) {
	gens["Spec"] = BastionHost_SpecGenerator()
	gens["Status"] = BastionHost_STATUSGenerator()
}

func Test_BastionHostIPConfiguration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BastionHostIPConfiguration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBastionHostIPConfiguration, BastionHostIPConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBastionHostIPConfiguration runs a test to see if a specific instance of BastionHostIPConfiguration round trips to JSON and back losslessly
func RunJSONSerializationTestForBastionHostIPConfiguration(subject BastionHostIPConfiguration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BastionHostIPConfiguration
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

// Generator of BastionHostIPConfiguration instances for property testing - lazily instantiated by
// BastionHostIPConfigurationGenerator()
var bastionHostIPConfigurationGenerator gopter.Gen

// BastionHostIPConfigurationGenerator returns a generator of BastionHostIPConfiguration instances for property testing.
// We first initialize bastionHostIPConfigurationGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BastionHostIPConfigurationGenerator() gopter.Gen {
	if bastionHostIPConfigurationGenerator != nil {
		return bastionHostIPConfigurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBastionHostIPConfiguration(generators)
	bastionHostIPConfigurationGenerator = gen.Struct(reflect.TypeOf(BastionHostIPConfiguration{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBastionHostIPConfiguration(generators)
	AddRelatedPropertyGeneratorsForBastionHostIPConfiguration(generators)
	bastionHostIPConfigurationGenerator = gen.Struct(reflect.TypeOf(BastionHostIPConfiguration{}), generators)

	return bastionHostIPConfigurationGenerator
}

// AddIndependentPropertyGeneratorsForBastionHostIPConfiguration is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBastionHostIPConfiguration(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateIPAllocationMethod"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBastionHostIPConfiguration is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBastionHostIPConfiguration(gens map[string]gopter.Gen) {
	gens["PublicIPAddress"] = gen.PtrOf(SubResourceGenerator())
	gens["Subnet"] = gen.PtrOf(SubResourceGenerator())
}

func Test_BastionHostIPConfiguration_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BastionHostIPConfiguration_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBastionHostIPConfiguration_STATUS, BastionHostIPConfiguration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBastionHostIPConfiguration_STATUS runs a test to see if a specific instance of BastionHostIPConfiguration_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForBastionHostIPConfiguration_STATUS(subject BastionHostIPConfiguration_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BastionHostIPConfiguration_STATUS
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

// Generator of BastionHostIPConfiguration_STATUS instances for property testing - lazily instantiated by
// BastionHostIPConfiguration_STATUSGenerator()
var bastionHostIPConfiguration_STATUSGenerator gopter.Gen

// BastionHostIPConfiguration_STATUSGenerator returns a generator of BastionHostIPConfiguration_STATUS instances for property testing.
func BastionHostIPConfiguration_STATUSGenerator() gopter.Gen {
	if bastionHostIPConfiguration_STATUSGenerator != nil {
		return bastionHostIPConfiguration_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBastionHostIPConfiguration_STATUS(generators)
	bastionHostIPConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(BastionHostIPConfiguration_STATUS{}), generators)

	return bastionHostIPConfiguration_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForBastionHostIPConfiguration_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBastionHostIPConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_BastionHostOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BastionHostOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBastionHostOperatorSpec, BastionHostOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBastionHostOperatorSpec runs a test to see if a specific instance of BastionHostOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForBastionHostOperatorSpec(subject BastionHostOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BastionHostOperatorSpec
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

// Generator of BastionHostOperatorSpec instances for property testing - lazily instantiated by
// BastionHostOperatorSpecGenerator()
var bastionHostOperatorSpecGenerator gopter.Gen

// BastionHostOperatorSpecGenerator returns a generator of BastionHostOperatorSpec instances for property testing.
func BastionHostOperatorSpecGenerator() gopter.Gen {
	if bastionHostOperatorSpecGenerator != nil {
		return bastionHostOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	bastionHostOperatorSpecGenerator = gen.Struct(reflect.TypeOf(BastionHostOperatorSpec{}), generators)

	return bastionHostOperatorSpecGenerator
}

func Test_BastionHost_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BastionHost_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBastionHost_STATUS, BastionHost_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBastionHost_STATUS runs a test to see if a specific instance of BastionHost_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForBastionHost_STATUS(subject BastionHost_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BastionHost_STATUS
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

// Generator of BastionHost_STATUS instances for property testing - lazily instantiated by BastionHost_STATUSGenerator()
var bastionHost_STATUSGenerator gopter.Gen

// BastionHost_STATUSGenerator returns a generator of BastionHost_STATUS instances for property testing.
// We first initialize bastionHost_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BastionHost_STATUSGenerator() gopter.Gen {
	if bastionHost_STATUSGenerator != nil {
		return bastionHost_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBastionHost_STATUS(generators)
	bastionHost_STATUSGenerator = gen.Struct(reflect.TypeOf(BastionHost_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBastionHost_STATUS(generators)
	AddRelatedPropertyGeneratorsForBastionHost_STATUS(generators)
	bastionHost_STATUSGenerator = gen.Struct(reflect.TypeOf(BastionHost_STATUS{}), generators)

	return bastionHost_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForBastionHost_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBastionHost_STATUS(gens map[string]gopter.Gen) {
	gens["DisableCopyPaste"] = gen.PtrOf(gen.Bool())
	gens["DnsName"] = gen.PtrOf(gen.AlphaString())
	gens["EnableFileCopy"] = gen.PtrOf(gen.Bool())
	gens["EnableIpConnect"] = gen.PtrOf(gen.Bool())
	gens["EnableShareableLink"] = gen.PtrOf(gen.Bool())
	gens["EnableTunneling"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ScaleUnits"] = gen.PtrOf(gen.Int())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBastionHost_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBastionHost_STATUS(gens map[string]gopter.Gen) {
	gens["IpConfigurations"] = gen.SliceOf(BastionHostIPConfiguration_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSGenerator())
}

func Test_BastionHost_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BastionHost_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBastionHost_Spec, BastionHost_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBastionHost_Spec runs a test to see if a specific instance of BastionHost_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForBastionHost_Spec(subject BastionHost_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BastionHost_Spec
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

// Generator of BastionHost_Spec instances for property testing - lazily instantiated by BastionHost_SpecGenerator()
var bastionHost_SpecGenerator gopter.Gen

// BastionHost_SpecGenerator returns a generator of BastionHost_Spec instances for property testing.
// We first initialize bastionHost_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BastionHost_SpecGenerator() gopter.Gen {
	if bastionHost_SpecGenerator != nil {
		return bastionHost_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBastionHost_Spec(generators)
	bastionHost_SpecGenerator = gen.Struct(reflect.TypeOf(BastionHost_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBastionHost_Spec(generators)
	AddRelatedPropertyGeneratorsForBastionHost_Spec(generators)
	bastionHost_SpecGenerator = gen.Struct(reflect.TypeOf(BastionHost_Spec{}), generators)

	return bastionHost_SpecGenerator
}

// AddIndependentPropertyGeneratorsForBastionHost_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBastionHost_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["DisableCopyPaste"] = gen.PtrOf(gen.Bool())
	gens["DnsName"] = gen.PtrOf(gen.AlphaString())
	gens["EnableFileCopy"] = gen.PtrOf(gen.Bool())
	gens["EnableIpConnect"] = gen.PtrOf(gen.Bool())
	gens["EnableShareableLink"] = gen.PtrOf(gen.Bool())
	gens["EnableTunneling"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["ScaleUnits"] = gen.PtrOf(gen.Int())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBastionHost_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBastionHost_Spec(gens map[string]gopter.Gen) {
	gens["IpConfigurations"] = gen.SliceOf(BastionHostIPConfigurationGenerator())
	gens["OperatorSpec"] = gen.PtrOf(BastionHostOperatorSpecGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_Sku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku runs a test to see if a specific instance of Sku round trips to JSON and back losslessly
func RunJSONSerializationTestForSku(subject Sku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku
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

// Generator of Sku instances for property testing - lazily instantiated by SkuGenerator()
var skuGenerator gopter.Gen

// SkuGenerator returns a generator of Sku instances for property testing.
func SkuGenerator() gopter.Gen {
	if skuGenerator != nil {
		return skuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku(generators)
	skuGenerator = gen.Struct(reflect.TypeOf(Sku{}), generators)

	return skuGenerator
}

// AddIndependentPropertyGeneratorsForSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUS, Sku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUS runs a test to see if a specific instance of Sku_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUS(subject Sku_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUS
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

// Generator of Sku_STATUS instances for property testing - lazily instantiated by Sku_STATUSGenerator()
var sku_STATUSGenerator gopter.Gen

// Sku_STATUSGenerator returns a generator of Sku_STATUS instances for property testing.
func Sku_STATUSGenerator() gopter.Gen {
	if sku_STATUSGenerator != nil {
		return sku_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUS(generators)
	sku_STATUSGenerator = gen.Struct(reflect.TypeOf(Sku_STATUS{}), generators)

	return sku_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}
