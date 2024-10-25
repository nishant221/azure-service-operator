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

func Test_VirtualMachineScaleSetExtensionProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineScaleSetExtensionProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineScaleSetExtensionProperties_STATUS, VirtualMachineScaleSetExtensionProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineScaleSetExtensionProperties_STATUS runs a test to see if a specific instance of VirtualMachineScaleSetExtensionProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineScaleSetExtensionProperties_STATUS(subject VirtualMachineScaleSetExtensionProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineScaleSetExtensionProperties_STATUS
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

// Generator of VirtualMachineScaleSetExtensionProperties_STATUS instances for property testing - lazily instantiated by
// VirtualMachineScaleSetExtensionProperties_STATUSGenerator()
var virtualMachineScaleSetExtensionProperties_STATUSGenerator gopter.Gen

// VirtualMachineScaleSetExtensionProperties_STATUSGenerator returns a generator of VirtualMachineScaleSetExtensionProperties_STATUS instances for property testing.
func VirtualMachineScaleSetExtensionProperties_STATUSGenerator() gopter.Gen {
	if virtualMachineScaleSetExtensionProperties_STATUSGenerator != nil {
		return virtualMachineScaleSetExtensionProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineScaleSetExtensionProperties_STATUS(generators)
	virtualMachineScaleSetExtensionProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualMachineScaleSetExtensionProperties_STATUS{}), generators)

	return virtualMachineScaleSetExtensionProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachineScaleSetExtensionProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachineScaleSetExtensionProperties_STATUS(gens map[string]gopter.Gen) {
	gens["AutoUpgradeMinorVersion"] = gen.PtrOf(gen.Bool())
	gens["EnableAutomaticUpgrade"] = gen.PtrOf(gen.Bool())
	gens["ForceUpdateTag"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisionAfterExtensions"] = gen.SliceOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Publisher"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}

func Test_VirtualMachineScaleSetsExtension_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineScaleSetsExtension_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineScaleSetsExtension_STATUS, VirtualMachineScaleSetsExtension_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineScaleSetsExtension_STATUS runs a test to see if a specific instance of VirtualMachineScaleSetsExtension_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineScaleSetsExtension_STATUS(subject VirtualMachineScaleSetsExtension_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineScaleSetsExtension_STATUS
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

// Generator of VirtualMachineScaleSetsExtension_STATUS instances for property testing - lazily instantiated by
// VirtualMachineScaleSetsExtension_STATUSGenerator()
var virtualMachineScaleSetsExtension_STATUSGenerator gopter.Gen

// VirtualMachineScaleSetsExtension_STATUSGenerator returns a generator of VirtualMachineScaleSetsExtension_STATUS instances for property testing.
// We first initialize virtualMachineScaleSetsExtension_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachineScaleSetsExtension_STATUSGenerator() gopter.Gen {
	if virtualMachineScaleSetsExtension_STATUSGenerator != nil {
		return virtualMachineScaleSetsExtension_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineScaleSetsExtension_STATUS(generators)
	virtualMachineScaleSetsExtension_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualMachineScaleSetsExtension_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineScaleSetsExtension_STATUS(generators)
	AddRelatedPropertyGeneratorsForVirtualMachineScaleSetsExtension_STATUS(generators)
	virtualMachineScaleSetsExtension_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualMachineScaleSetsExtension_STATUS{}), generators)

	return virtualMachineScaleSetsExtension_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachineScaleSetsExtension_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachineScaleSetsExtension_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachineScaleSetsExtension_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachineScaleSetsExtension_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualMachineScaleSetExtensionProperties_STATUSGenerator())
}