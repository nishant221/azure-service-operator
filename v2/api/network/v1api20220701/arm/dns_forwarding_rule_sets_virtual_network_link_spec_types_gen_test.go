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

func Test_DnsForwardingRuleSetsVirtualNetworkLink_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsForwardingRuleSetsVirtualNetworkLink_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsForwardingRuleSetsVirtualNetworkLink_Spec, DnsForwardingRuleSetsVirtualNetworkLink_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsForwardingRuleSetsVirtualNetworkLink_Spec runs a test to see if a specific instance of DnsForwardingRuleSetsVirtualNetworkLink_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsForwardingRuleSetsVirtualNetworkLink_Spec(subject DnsForwardingRuleSetsVirtualNetworkLink_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsForwardingRuleSetsVirtualNetworkLink_Spec
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

// Generator of DnsForwardingRuleSetsVirtualNetworkLink_Spec instances for property testing - lazily instantiated by
// DnsForwardingRuleSetsVirtualNetworkLink_SpecGenerator()
var dnsForwardingRuleSetsVirtualNetworkLink_SpecGenerator gopter.Gen

// DnsForwardingRuleSetsVirtualNetworkLink_SpecGenerator returns a generator of DnsForwardingRuleSetsVirtualNetworkLink_Spec instances for property testing.
// We first initialize dnsForwardingRuleSetsVirtualNetworkLink_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsForwardingRuleSetsVirtualNetworkLink_SpecGenerator() gopter.Gen {
	if dnsForwardingRuleSetsVirtualNetworkLink_SpecGenerator != nil {
		return dnsForwardingRuleSetsVirtualNetworkLink_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink_Spec(generators)
	dnsForwardingRuleSetsVirtualNetworkLink_SpecGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRuleSetsVirtualNetworkLink_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink_Spec(generators)
	dnsForwardingRuleSetsVirtualNetworkLink_SpecGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRuleSetsVirtualNetworkLink_Spec{}), generators)

	return dnsForwardingRuleSetsVirtualNetworkLink_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualNetworkLinkPropertiesGenerator())
}

func Test_VirtualNetworkLinkProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkLinkProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkLinkProperties, VirtualNetworkLinkPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkLinkProperties runs a test to see if a specific instance of VirtualNetworkLinkProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkLinkProperties(subject VirtualNetworkLinkProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkLinkProperties
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

// Generator of VirtualNetworkLinkProperties instances for property testing - lazily instantiated by
// VirtualNetworkLinkPropertiesGenerator()
var virtualNetworkLinkPropertiesGenerator gopter.Gen

// VirtualNetworkLinkPropertiesGenerator returns a generator of VirtualNetworkLinkProperties instances for property testing.
// We first initialize virtualNetworkLinkPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkLinkPropertiesGenerator() gopter.Gen {
	if virtualNetworkLinkPropertiesGenerator != nil {
		return virtualNetworkLinkPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkLinkProperties(generators)
	virtualNetworkLinkPropertiesGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkLinkProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkLinkProperties(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkLinkProperties(generators)
	virtualNetworkLinkPropertiesGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkLinkProperties{}), generators)

	return virtualNetworkLinkPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkLinkProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkLinkProperties(gens map[string]gopter.Gen) {
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkLinkProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkLinkProperties(gens map[string]gopter.Gen) {
	gens["VirtualNetwork"] = gen.PtrOf(SubResourceGenerator())
}
