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

func Test_DnsResolversOutboundEndpoint_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsResolversOutboundEndpoint_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsResolversOutboundEndpoint_STATUS, DnsResolversOutboundEndpoint_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsResolversOutboundEndpoint_STATUS runs a test to see if a specific instance of DnsResolversOutboundEndpoint_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsResolversOutboundEndpoint_STATUS(subject DnsResolversOutboundEndpoint_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsResolversOutboundEndpoint_STATUS
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

// Generator of DnsResolversOutboundEndpoint_STATUS instances for property testing - lazily instantiated by
// DnsResolversOutboundEndpoint_STATUSGenerator()
var dnsResolversOutboundEndpoint_STATUSGenerator gopter.Gen

// DnsResolversOutboundEndpoint_STATUSGenerator returns a generator of DnsResolversOutboundEndpoint_STATUS instances for property testing.
// We first initialize dnsResolversOutboundEndpoint_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsResolversOutboundEndpoint_STATUSGenerator() gopter.Gen {
	if dnsResolversOutboundEndpoint_STATUSGenerator != nil {
		return dnsResolversOutboundEndpoint_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolversOutboundEndpoint_STATUS(generators)
	dnsResolversOutboundEndpoint_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsResolversOutboundEndpoint_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolversOutboundEndpoint_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsResolversOutboundEndpoint_STATUS(generators)
	dnsResolversOutboundEndpoint_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsResolversOutboundEndpoint_STATUS{}), generators)

	return dnsResolversOutboundEndpoint_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsResolversOutboundEndpoint_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsResolversOutboundEndpoint_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsResolversOutboundEndpoint_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsResolversOutboundEndpoint_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(OutboundEndpointProperties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_OutboundEndpointProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of OutboundEndpointProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForOutboundEndpointProperties_STATUS, OutboundEndpointProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForOutboundEndpointProperties_STATUS runs a test to see if a specific instance of OutboundEndpointProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForOutboundEndpointProperties_STATUS(subject OutboundEndpointProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual OutboundEndpointProperties_STATUS
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

// Generator of OutboundEndpointProperties_STATUS instances for property testing - lazily instantiated by
// OutboundEndpointProperties_STATUSGenerator()
var outboundEndpointProperties_STATUSGenerator gopter.Gen

// OutboundEndpointProperties_STATUSGenerator returns a generator of OutboundEndpointProperties_STATUS instances for property testing.
// We first initialize outboundEndpointProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func OutboundEndpointProperties_STATUSGenerator() gopter.Gen {
	if outboundEndpointProperties_STATUSGenerator != nil {
		return outboundEndpointProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOutboundEndpointProperties_STATUS(generators)
	outboundEndpointProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(OutboundEndpointProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOutboundEndpointProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForOutboundEndpointProperties_STATUS(generators)
	outboundEndpointProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(OutboundEndpointProperties_STATUS{}), generators)

	return outboundEndpointProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForOutboundEndpointProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForOutboundEndpointProperties_STATUS(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DnsresolverProvisioningState_STATUS_Canceled,
		DnsresolverProvisioningState_STATUS_Creating,
		DnsresolverProvisioningState_STATUS_Deleting,
		DnsresolverProvisioningState_STATUS_Failed,
		DnsresolverProvisioningState_STATUS_Succeeded,
		DnsresolverProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForOutboundEndpointProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForOutboundEndpointProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Subnet"] = gen.PtrOf(DnsresolverSubResource_STATUSGenerator())
}