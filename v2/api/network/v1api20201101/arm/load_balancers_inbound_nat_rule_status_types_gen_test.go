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

func Test_InboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded, InboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded runs a test to see if a specific instance of InboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForInboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded(subject InboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded
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

// Generator of InboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded instances for
// property testing - lazily instantiated by
// InboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator()
var inboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator gopter.Gen

// InboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator returns a generator of InboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded instances for property testing.
// We first initialize inboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func InboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator() gopter.Gen {
	if inboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator != nil {
		return inboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded(generators)
	inboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(InboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded(generators)
	AddRelatedPropertyGeneratorsForInboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded(generators)
	inboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(InboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded{}), generators)

	return inboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForInboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["BackendPort"] = gen.PtrOf(gen.Int())
	gens["EnableFloatingIP"] = gen.PtrOf(gen.Bool())
	gens["EnableTcpReset"] = gen.PtrOf(gen.Bool())
	gens["FrontendPort"] = gen.PtrOf(gen.Int())
	gens["IdleTimeoutInMinutes"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(TransportProtocol_STATUS_All, TransportProtocol_STATUS_Tcp, TransportProtocol_STATUS_Udp))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
}

// AddRelatedPropertyGeneratorsForInboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForInboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["BackendIPConfiguration"] = gen.PtrOf(NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator())
	gens["FrontendIPConfiguration"] = gen.PtrOf(SubResource_STATUSGenerator())
}

func Test_LoadBalancersInboundNatRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LoadBalancersInboundNatRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLoadBalancersInboundNatRule_STATUS, LoadBalancersInboundNatRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLoadBalancersInboundNatRule_STATUS runs a test to see if a specific instance of LoadBalancersInboundNatRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForLoadBalancersInboundNatRule_STATUS(subject LoadBalancersInboundNatRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LoadBalancersInboundNatRule_STATUS
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

// Generator of LoadBalancersInboundNatRule_STATUS instances for property testing - lazily instantiated by
// LoadBalancersInboundNatRule_STATUSGenerator()
var loadBalancersInboundNatRule_STATUSGenerator gopter.Gen

// LoadBalancersInboundNatRule_STATUSGenerator returns a generator of LoadBalancersInboundNatRule_STATUS instances for property testing.
// We first initialize loadBalancersInboundNatRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LoadBalancersInboundNatRule_STATUSGenerator() gopter.Gen {
	if loadBalancersInboundNatRule_STATUSGenerator != nil {
		return loadBalancersInboundNatRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLoadBalancersInboundNatRule_STATUS(generators)
	loadBalancersInboundNatRule_STATUSGenerator = gen.Struct(reflect.TypeOf(LoadBalancersInboundNatRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLoadBalancersInboundNatRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForLoadBalancersInboundNatRule_STATUS(generators)
	loadBalancersInboundNatRule_STATUSGenerator = gen.Struct(reflect.TypeOf(LoadBalancersInboundNatRule_STATUS{}), generators)

	return loadBalancersInboundNatRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForLoadBalancersInboundNatRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLoadBalancersInboundNatRule_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForLoadBalancersInboundNatRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLoadBalancersInboundNatRule_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(InboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator())
}

func Test_NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded, NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded runs a test to see if a specific instance of NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded(subject NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded
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

// Generator of NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded instances for
// property testing - lazily instantiated by
// NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator()
var networkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator gopter.Gen

// NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator returns a generator of NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded instances for property testing.
func NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator() gopter.Gen {
	if networkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator != nil {
		return networkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded(generators)
	networkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded{}), generators)

	return networkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
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