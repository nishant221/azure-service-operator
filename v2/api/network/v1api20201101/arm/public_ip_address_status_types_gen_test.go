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

func Test_DdosSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DdosSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDdosSettings_STATUS, DdosSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDdosSettings_STATUS runs a test to see if a specific instance of DdosSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDdosSettings_STATUS(subject DdosSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DdosSettings_STATUS
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

// Generator of DdosSettings_STATUS instances for property testing - lazily instantiated by
// DdosSettings_STATUSGenerator()
var ddosSettings_STATUSGenerator gopter.Gen

// DdosSettings_STATUSGenerator returns a generator of DdosSettings_STATUS instances for property testing.
// We first initialize ddosSettings_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DdosSettings_STATUSGenerator() gopter.Gen {
	if ddosSettings_STATUSGenerator != nil {
		return ddosSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDdosSettings_STATUS(generators)
	ddosSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(DdosSettings_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDdosSettings_STATUS(generators)
	AddRelatedPropertyGeneratorsForDdosSettings_STATUS(generators)
	ddosSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(DdosSettings_STATUS{}), generators)

	return ddosSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDdosSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDdosSettings_STATUS(gens map[string]gopter.Gen) {
	gens["ProtectedIP"] = gen.PtrOf(gen.Bool())
	gens["ProtectionCoverage"] = gen.PtrOf(gen.OneConstOf(DdosSettings_ProtectionCoverage_STATUS_Basic, DdosSettings_ProtectionCoverage_STATUS_Standard))
}

// AddRelatedPropertyGeneratorsForDdosSettings_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDdosSettings_STATUS(gens map[string]gopter.Gen) {
	gens["DdosCustomPolicy"] = gen.PtrOf(SubResource_STATUSGenerator())
}

func Test_IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded, IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded runs a test to see if a specific instance of IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForIPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded(subject IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded
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

// Generator of IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded instances for property testing - lazily
// instantiated by IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator()
var ipConfiguration_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator gopter.Gen

// IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator returns a generator of IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded instances for property testing.
func IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator() gopter.Gen {
	if ipConfiguration_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator != nil {
		return ipConfiguration_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded(generators)
	ipConfiguration_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded{}), generators)

	return ipConfiguration_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForIPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_IpTag_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IpTag_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIpTag_STATUS, IpTag_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIpTag_STATUS runs a test to see if a specific instance of IpTag_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForIpTag_STATUS(subject IpTag_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IpTag_STATUS
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

// Generator of IpTag_STATUS instances for property testing - lazily instantiated by IpTag_STATUSGenerator()
var ipTag_STATUSGenerator gopter.Gen

// IpTag_STATUSGenerator returns a generator of IpTag_STATUS instances for property testing.
func IpTag_STATUSGenerator() gopter.Gen {
	if ipTag_STATUSGenerator != nil {
		return ipTag_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpTag_STATUS(generators)
	ipTag_STATUSGenerator = gen.Struct(reflect.TypeOf(IpTag_STATUS{}), generators)

	return ipTag_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForIpTag_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIpTag_STATUS(gens map[string]gopter.Gen) {
	gens["IpTagType"] = gen.PtrOf(gen.AlphaString())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
}

func Test_NatGateway_STATUS_PublicIPAddress_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGateway_STATUS_PublicIPAddress_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGateway_STATUS_PublicIPAddress_SubResourceEmbedded, NatGateway_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGateway_STATUS_PublicIPAddress_SubResourceEmbedded runs a test to see if a specific instance of NatGateway_STATUS_PublicIPAddress_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGateway_STATUS_PublicIPAddress_SubResourceEmbedded(subject NatGateway_STATUS_PublicIPAddress_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGateway_STATUS_PublicIPAddress_SubResourceEmbedded
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

// Generator of NatGateway_STATUS_PublicIPAddress_SubResourceEmbedded instances for property testing - lazily
// instantiated by NatGateway_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator()
var natGateway_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator gopter.Gen

// NatGateway_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator returns a generator of NatGateway_STATUS_PublicIPAddress_SubResourceEmbedded instances for property testing.
func NatGateway_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator() gopter.Gen {
	if natGateway_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator != nil {
		return natGateway_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGateway_STATUS_PublicIPAddress_SubResourceEmbedded(generators)
	natGateway_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NatGateway_STATUS_PublicIPAddress_SubResourceEmbedded{}), generators)

	return natGateway_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForNatGateway_STATUS_PublicIPAddress_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNatGateway_STATUS_PublicIPAddress_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_PublicIPAddressDnsSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddressDnsSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddressDnsSettings_STATUS, PublicIPAddressDnsSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddressDnsSettings_STATUS runs a test to see if a specific instance of PublicIPAddressDnsSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddressDnsSettings_STATUS(subject PublicIPAddressDnsSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPAddressDnsSettings_STATUS
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

// Generator of PublicIPAddressDnsSettings_STATUS instances for property testing - lazily instantiated by
// PublicIPAddressDnsSettings_STATUSGenerator()
var publicIPAddressDnsSettings_STATUSGenerator gopter.Gen

// PublicIPAddressDnsSettings_STATUSGenerator returns a generator of PublicIPAddressDnsSettings_STATUS instances for property testing.
func PublicIPAddressDnsSettings_STATUSGenerator() gopter.Gen {
	if publicIPAddressDnsSettings_STATUSGenerator != nil {
		return publicIPAddressDnsSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressDnsSettings_STATUS(generators)
	publicIPAddressDnsSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressDnsSettings_STATUS{}), generators)

	return publicIPAddressDnsSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddressDnsSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddressDnsSettings_STATUS(gens map[string]gopter.Gen) {
	gens["DomainNameLabel"] = gen.PtrOf(gen.AlphaString())
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["ReverseFqdn"] = gen.PtrOf(gen.AlphaString())
}

func Test_PublicIPAddressPropertiesFormat_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddressPropertiesFormat_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddressPropertiesFormat_STATUS, PublicIPAddressPropertiesFormat_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddressPropertiesFormat_STATUS runs a test to see if a specific instance of PublicIPAddressPropertiesFormat_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddressPropertiesFormat_STATUS(subject PublicIPAddressPropertiesFormat_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPAddressPropertiesFormat_STATUS
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

// Generator of PublicIPAddressPropertiesFormat_STATUS instances for property testing - lazily instantiated by
// PublicIPAddressPropertiesFormat_STATUSGenerator()
var publicIPAddressPropertiesFormat_STATUSGenerator gopter.Gen

// PublicIPAddressPropertiesFormat_STATUSGenerator returns a generator of PublicIPAddressPropertiesFormat_STATUS instances for property testing.
// We first initialize publicIPAddressPropertiesFormat_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PublicIPAddressPropertiesFormat_STATUSGenerator() gopter.Gen {
	if publicIPAddressPropertiesFormat_STATUSGenerator != nil {
		return publicIPAddressPropertiesFormat_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressPropertiesFormat_STATUS(generators)
	publicIPAddressPropertiesFormat_STATUSGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressPropertiesFormat_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressPropertiesFormat_STATUS(generators)
	AddRelatedPropertyGeneratorsForPublicIPAddressPropertiesFormat_STATUS(generators)
	publicIPAddressPropertiesFormat_STATUSGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressPropertiesFormat_STATUS{}), generators)

	return publicIPAddressPropertiesFormat_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddressPropertiesFormat_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddressPropertiesFormat_STATUS(gens map[string]gopter.Gen) {
	gens["IdleTimeoutInMinutes"] = gen.PtrOf(gen.Int())
	gens["IpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["MigrationPhase"] = gen.PtrOf(gen.OneConstOf(
		PublicIPAddressPropertiesFormat_MigrationPhase_STATUS_Abort,
		PublicIPAddressPropertiesFormat_MigrationPhase_STATUS_Commit,
		PublicIPAddressPropertiesFormat_MigrationPhase_STATUS_Committed,
		PublicIPAddressPropertiesFormat_MigrationPhase_STATUS_None,
		PublicIPAddressPropertiesFormat_MigrationPhase_STATUS_Prepare))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["PublicIPAddressVersion"] = gen.PtrOf(gen.OneConstOf(IPVersion_STATUS_IPv4, IPVersion_STATUS_IPv6))
	gens["PublicIPAllocationMethod"] = gen.PtrOf(gen.OneConstOf(IPAllocationMethod_STATUS_Dynamic, IPAllocationMethod_STATUS_Static))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPublicIPAddressPropertiesFormat_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPublicIPAddressPropertiesFormat_STATUS(gens map[string]gopter.Gen) {
	gens["DdosSettings"] = gen.PtrOf(DdosSettings_STATUSGenerator())
	gens["DnsSettings"] = gen.PtrOf(PublicIPAddressDnsSettings_STATUSGenerator())
	gens["IpConfiguration"] = gen.PtrOf(IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator())
	gens["IpTags"] = gen.SliceOf(IpTag_STATUSGenerator())
	gens["NatGateway"] = gen.PtrOf(NatGateway_STATUS_PublicIPAddress_SubResourceEmbeddedGenerator())
	gens["PublicIPPrefix"] = gen.PtrOf(SubResource_STATUSGenerator())
}

func Test_PublicIPAddressSku_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddressSku_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddressSku_STATUS, PublicIPAddressSku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddressSku_STATUS runs a test to see if a specific instance of PublicIPAddressSku_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddressSku_STATUS(subject PublicIPAddressSku_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPAddressSku_STATUS
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

// Generator of PublicIPAddressSku_STATUS instances for property testing - lazily instantiated by
// PublicIPAddressSku_STATUSGenerator()
var publicIPAddressSku_STATUSGenerator gopter.Gen

// PublicIPAddressSku_STATUSGenerator returns a generator of PublicIPAddressSku_STATUS instances for property testing.
func PublicIPAddressSku_STATUSGenerator() gopter.Gen {
	if publicIPAddressSku_STATUSGenerator != nil {
		return publicIPAddressSku_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressSku_STATUS(generators)
	publicIPAddressSku_STATUSGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressSku_STATUS{}), generators)

	return publicIPAddressSku_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddressSku_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddressSku_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(PublicIPAddressSku_Name_STATUS_Basic, PublicIPAddressSku_Name_STATUS_Standard))
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(PublicIPAddressSku_Tier_STATUS_Global, PublicIPAddressSku_Tier_STATUS_Regional))
}

func Test_PublicIPAddress_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddress_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddress_STATUS, PublicIPAddress_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddress_STATUS runs a test to see if a specific instance of PublicIPAddress_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddress_STATUS(subject PublicIPAddress_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPAddress_STATUS
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

// Generator of PublicIPAddress_STATUS instances for property testing - lazily instantiated by
// PublicIPAddress_STATUSGenerator()
var publicIPAddress_STATUSGenerator gopter.Gen

// PublicIPAddress_STATUSGenerator returns a generator of PublicIPAddress_STATUS instances for property testing.
// We first initialize publicIPAddress_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PublicIPAddress_STATUSGenerator() gopter.Gen {
	if publicIPAddress_STATUSGenerator != nil {
		return publicIPAddress_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddress_STATUS(generators)
	publicIPAddress_STATUSGenerator = gen.Struct(reflect.TypeOf(PublicIPAddress_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddress_STATUS(generators)
	AddRelatedPropertyGeneratorsForPublicIPAddress_STATUS(generators)
	publicIPAddress_STATUSGenerator = gen.Struct(reflect.TypeOf(PublicIPAddress_STATUS{}), generators)

	return publicIPAddress_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddress_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddress_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPublicIPAddress_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPublicIPAddress_STATUS(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_STATUSGenerator())
	gens["Properties"] = gen.PtrOf(PublicIPAddressPropertiesFormat_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(PublicIPAddressSku_STATUSGenerator())
}