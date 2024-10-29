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

func Test_ApiContactInformation_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApiContactInformation via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApiContactInformation, ApiContactInformationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApiContactInformation runs a test to see if a specific instance of ApiContactInformation round trips to JSON and back losslessly
func RunJSONSerializationTestForApiContactInformation(subject ApiContactInformation) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApiContactInformation
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

// Generator of ApiContactInformation instances for property testing - lazily instantiated by
// ApiContactInformationGenerator()
var apiContactInformationGenerator gopter.Gen

// ApiContactInformationGenerator returns a generator of ApiContactInformation instances for property testing.
func ApiContactInformationGenerator() gopter.Gen {
	if apiContactInformationGenerator != nil {
		return apiContactInformationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiContactInformation(generators)
	apiContactInformationGenerator = gen.Struct(reflect.TypeOf(ApiContactInformation{}), generators)

	return apiContactInformationGenerator
}

// AddIndependentPropertyGeneratorsForApiContactInformation is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApiContactInformation(gens map[string]gopter.Gen) {
	gens["Email"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Url"] = gen.PtrOf(gen.AlphaString())
}

func Test_ApiCreateOrUpdateProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApiCreateOrUpdateProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApiCreateOrUpdateProperties, ApiCreateOrUpdatePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApiCreateOrUpdateProperties runs a test to see if a specific instance of ApiCreateOrUpdateProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForApiCreateOrUpdateProperties(subject ApiCreateOrUpdateProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApiCreateOrUpdateProperties
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

// Generator of ApiCreateOrUpdateProperties instances for property testing - lazily instantiated by
// ApiCreateOrUpdatePropertiesGenerator()
var apiCreateOrUpdatePropertiesGenerator gopter.Gen

// ApiCreateOrUpdatePropertiesGenerator returns a generator of ApiCreateOrUpdateProperties instances for property testing.
// We first initialize apiCreateOrUpdatePropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ApiCreateOrUpdatePropertiesGenerator() gopter.Gen {
	if apiCreateOrUpdatePropertiesGenerator != nil {
		return apiCreateOrUpdatePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiCreateOrUpdateProperties(generators)
	apiCreateOrUpdatePropertiesGenerator = gen.Struct(reflect.TypeOf(ApiCreateOrUpdateProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiCreateOrUpdateProperties(generators)
	AddRelatedPropertyGeneratorsForApiCreateOrUpdateProperties(generators)
	apiCreateOrUpdatePropertiesGenerator = gen.Struct(reflect.TypeOf(ApiCreateOrUpdateProperties{}), generators)

	return apiCreateOrUpdatePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForApiCreateOrUpdateProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApiCreateOrUpdateProperties(gens map[string]gopter.Gen) {
	gens["APIVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ApiRevision"] = gen.PtrOf(gen.AlphaString())
	gens["ApiRevisionDescription"] = gen.PtrOf(gen.AlphaString())
	gens["ApiType"] = gen.PtrOf(gen.OneConstOf(
		ApiCreateOrUpdateProperties_ApiType_Graphql,
		ApiCreateOrUpdateProperties_ApiType_Grpc,
		ApiCreateOrUpdateProperties_ApiType_Http,
		ApiCreateOrUpdateProperties_ApiType_Odata,
		ApiCreateOrUpdateProperties_ApiType_Soap,
		ApiCreateOrUpdateProperties_ApiType_Websocket))
	gens["ApiVersionDescription"] = gen.PtrOf(gen.AlphaString())
	gens["ApiVersionSetId"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["Format"] = gen.PtrOf(gen.OneConstOf(
		ApiCreateOrUpdateProperties_Format_GraphqlLink,
		ApiCreateOrUpdateProperties_Format_Grpc,
		ApiCreateOrUpdateProperties_Format_GrpcLink,
		ApiCreateOrUpdateProperties_Format_Odata,
		ApiCreateOrUpdateProperties_Format_OdataLink,
		ApiCreateOrUpdateProperties_Format_Openapi,
		ApiCreateOrUpdateProperties_Format_OpenapiJson,
		ApiCreateOrUpdateProperties_Format_OpenapiJsonLink,
		ApiCreateOrUpdateProperties_Format_OpenapiLink,
		ApiCreateOrUpdateProperties_Format_SwaggerJson,
		ApiCreateOrUpdateProperties_Format_SwaggerLinkJson,
		ApiCreateOrUpdateProperties_Format_WadlLinkJson,
		ApiCreateOrUpdateProperties_Format_WadlXml,
		ApiCreateOrUpdateProperties_Format_Wsdl,
		ApiCreateOrUpdateProperties_Format_WsdlLink))
	gens["IsCurrent"] = gen.PtrOf(gen.Bool())
	gens["Path"] = gen.PtrOf(gen.AlphaString())
	gens["Protocols"] = gen.SliceOf(gen.OneConstOf(
		ApiCreateOrUpdateProperties_Protocols_Http,
		ApiCreateOrUpdateProperties_Protocols_Https,
		ApiCreateOrUpdateProperties_Protocols_Ws,
		ApiCreateOrUpdateProperties_Protocols_Wss))
	gens["ServiceUrl"] = gen.PtrOf(gen.AlphaString())
	gens["SourceApiId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionRequired"] = gen.PtrOf(gen.Bool())
	gens["TermsOfServiceUrl"] = gen.PtrOf(gen.AlphaString())
	gens["TranslateRequiredQueryParameters"] = gen.PtrOf(gen.OneConstOf(ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters_Query, ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters_Template))
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		ApiCreateOrUpdateProperties_Type_Graphql,
		ApiCreateOrUpdateProperties_Type_Grpc,
		ApiCreateOrUpdateProperties_Type_Http,
		ApiCreateOrUpdateProperties_Type_Odata,
		ApiCreateOrUpdateProperties_Type_Soap,
		ApiCreateOrUpdateProperties_Type_Websocket))
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForApiCreateOrUpdateProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForApiCreateOrUpdateProperties(gens map[string]gopter.Gen) {
	gens["ApiVersionSet"] = gen.PtrOf(ApiVersionSetContractDetailsGenerator())
	gens["AuthenticationSettings"] = gen.PtrOf(AuthenticationSettingsContractGenerator())
	gens["Contact"] = gen.PtrOf(ApiContactInformationGenerator())
	gens["License"] = gen.PtrOf(ApiLicenseInformationGenerator())
	gens["SubscriptionKeyParameterNames"] = gen.PtrOf(SubscriptionKeyParameterNamesContractGenerator())
	gens["WsdlSelector"] = gen.PtrOf(ApiCreateOrUpdateProperties_WsdlSelectorGenerator())
}

func Test_ApiCreateOrUpdateProperties_WsdlSelector_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApiCreateOrUpdateProperties_WsdlSelector via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApiCreateOrUpdateProperties_WsdlSelector, ApiCreateOrUpdateProperties_WsdlSelectorGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApiCreateOrUpdateProperties_WsdlSelector runs a test to see if a specific instance of ApiCreateOrUpdateProperties_WsdlSelector round trips to JSON and back losslessly
func RunJSONSerializationTestForApiCreateOrUpdateProperties_WsdlSelector(subject ApiCreateOrUpdateProperties_WsdlSelector) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApiCreateOrUpdateProperties_WsdlSelector
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

// Generator of ApiCreateOrUpdateProperties_WsdlSelector instances for property testing - lazily instantiated by
// ApiCreateOrUpdateProperties_WsdlSelectorGenerator()
var apiCreateOrUpdateProperties_WsdlSelectorGenerator gopter.Gen

// ApiCreateOrUpdateProperties_WsdlSelectorGenerator returns a generator of ApiCreateOrUpdateProperties_WsdlSelector instances for property testing.
func ApiCreateOrUpdateProperties_WsdlSelectorGenerator() gopter.Gen {
	if apiCreateOrUpdateProperties_WsdlSelectorGenerator != nil {
		return apiCreateOrUpdateProperties_WsdlSelectorGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiCreateOrUpdateProperties_WsdlSelector(generators)
	apiCreateOrUpdateProperties_WsdlSelectorGenerator = gen.Struct(reflect.TypeOf(ApiCreateOrUpdateProperties_WsdlSelector{}), generators)

	return apiCreateOrUpdateProperties_WsdlSelectorGenerator
}

// AddIndependentPropertyGeneratorsForApiCreateOrUpdateProperties_WsdlSelector is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApiCreateOrUpdateProperties_WsdlSelector(gens map[string]gopter.Gen) {
	gens["WsdlEndpointName"] = gen.PtrOf(gen.AlphaString())
	gens["WsdlServiceName"] = gen.PtrOf(gen.AlphaString())
}

func Test_ApiLicenseInformation_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApiLicenseInformation via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApiLicenseInformation, ApiLicenseInformationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApiLicenseInformation runs a test to see if a specific instance of ApiLicenseInformation round trips to JSON and back losslessly
func RunJSONSerializationTestForApiLicenseInformation(subject ApiLicenseInformation) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApiLicenseInformation
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

// Generator of ApiLicenseInformation instances for property testing - lazily instantiated by
// ApiLicenseInformationGenerator()
var apiLicenseInformationGenerator gopter.Gen

// ApiLicenseInformationGenerator returns a generator of ApiLicenseInformation instances for property testing.
func ApiLicenseInformationGenerator() gopter.Gen {
	if apiLicenseInformationGenerator != nil {
		return apiLicenseInformationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiLicenseInformation(generators)
	apiLicenseInformationGenerator = gen.Struct(reflect.TypeOf(ApiLicenseInformation{}), generators)

	return apiLicenseInformationGenerator
}

// AddIndependentPropertyGeneratorsForApiLicenseInformation is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApiLicenseInformation(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Url"] = gen.PtrOf(gen.AlphaString())
}

func Test_ApiVersionSetContractDetails_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApiVersionSetContractDetails via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApiVersionSetContractDetails, ApiVersionSetContractDetailsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApiVersionSetContractDetails runs a test to see if a specific instance of ApiVersionSetContractDetails round trips to JSON and back losslessly
func RunJSONSerializationTestForApiVersionSetContractDetails(subject ApiVersionSetContractDetails) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApiVersionSetContractDetails
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

// Generator of ApiVersionSetContractDetails instances for property testing - lazily instantiated by
// ApiVersionSetContractDetailsGenerator()
var apiVersionSetContractDetailsGenerator gopter.Gen

// ApiVersionSetContractDetailsGenerator returns a generator of ApiVersionSetContractDetails instances for property testing.
func ApiVersionSetContractDetailsGenerator() gopter.Gen {
	if apiVersionSetContractDetailsGenerator != nil {
		return apiVersionSetContractDetailsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiVersionSetContractDetails(generators)
	apiVersionSetContractDetailsGenerator = gen.Struct(reflect.TypeOf(ApiVersionSetContractDetails{}), generators)

	return apiVersionSetContractDetailsGenerator
}

// AddIndependentPropertyGeneratorsForApiVersionSetContractDetails is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApiVersionSetContractDetails(gens map[string]gopter.Gen) {
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["VersionHeaderName"] = gen.PtrOf(gen.AlphaString())
	gens["VersionQueryName"] = gen.PtrOf(gen.AlphaString())
	gens["VersioningScheme"] = gen.PtrOf(gen.OneConstOf(ApiVersionSetContractDetails_VersioningScheme_Header, ApiVersionSetContractDetails_VersioningScheme_Query, ApiVersionSetContractDetails_VersioningScheme_Segment))
}

func Test_Api_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Api_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApi_Spec, Api_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApi_Spec runs a test to see if a specific instance of Api_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForApi_Spec(subject Api_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Api_Spec
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

// Generator of Api_Spec instances for property testing - lazily instantiated by Api_SpecGenerator()
var api_SpecGenerator gopter.Gen

// Api_SpecGenerator returns a generator of Api_Spec instances for property testing.
// We first initialize api_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Api_SpecGenerator() gopter.Gen {
	if api_SpecGenerator != nil {
		return api_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApi_Spec(generators)
	api_SpecGenerator = gen.Struct(reflect.TypeOf(Api_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApi_Spec(generators)
	AddRelatedPropertyGeneratorsForApi_Spec(generators)
	api_SpecGenerator = gen.Struct(reflect.TypeOf(Api_Spec{}), generators)

	return api_SpecGenerator
}

// AddIndependentPropertyGeneratorsForApi_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApi_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForApi_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForApi_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ApiCreateOrUpdatePropertiesGenerator())
}

func Test_AuthenticationSettingsContract_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthenticationSettingsContract via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthenticationSettingsContract, AuthenticationSettingsContractGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthenticationSettingsContract runs a test to see if a specific instance of AuthenticationSettingsContract round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthenticationSettingsContract(subject AuthenticationSettingsContract) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthenticationSettingsContract
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

// Generator of AuthenticationSettingsContract instances for property testing - lazily instantiated by
// AuthenticationSettingsContractGenerator()
var authenticationSettingsContractGenerator gopter.Gen

// AuthenticationSettingsContractGenerator returns a generator of AuthenticationSettingsContract instances for property testing.
func AuthenticationSettingsContractGenerator() gopter.Gen {
	if authenticationSettingsContractGenerator != nil {
		return authenticationSettingsContractGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAuthenticationSettingsContract(generators)
	authenticationSettingsContractGenerator = gen.Struct(reflect.TypeOf(AuthenticationSettingsContract{}), generators)

	return authenticationSettingsContractGenerator
}

// AddRelatedPropertyGeneratorsForAuthenticationSettingsContract is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthenticationSettingsContract(gens map[string]gopter.Gen) {
	gens["OAuth2"] = gen.PtrOf(OAuth2AuthenticationSettingsContractGenerator())
	gens["OAuth2AuthenticationSettings"] = gen.SliceOf(OAuth2AuthenticationSettingsContractGenerator())
	gens["Openid"] = gen.PtrOf(OpenIdAuthenticationSettingsContractGenerator())
	gens["OpenidAuthenticationSettings"] = gen.SliceOf(OpenIdAuthenticationSettingsContractGenerator())
}

func Test_OAuth2AuthenticationSettingsContract_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of OAuth2AuthenticationSettingsContract via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForOAuth2AuthenticationSettingsContract, OAuth2AuthenticationSettingsContractGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForOAuth2AuthenticationSettingsContract runs a test to see if a specific instance of OAuth2AuthenticationSettingsContract round trips to JSON and back losslessly
func RunJSONSerializationTestForOAuth2AuthenticationSettingsContract(subject OAuth2AuthenticationSettingsContract) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual OAuth2AuthenticationSettingsContract
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

// Generator of OAuth2AuthenticationSettingsContract instances for property testing - lazily instantiated by
// OAuth2AuthenticationSettingsContractGenerator()
var oAuth2AuthenticationSettingsContractGenerator gopter.Gen

// OAuth2AuthenticationSettingsContractGenerator returns a generator of OAuth2AuthenticationSettingsContract instances for property testing.
func OAuth2AuthenticationSettingsContractGenerator() gopter.Gen {
	if oAuth2AuthenticationSettingsContractGenerator != nil {
		return oAuth2AuthenticationSettingsContractGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOAuth2AuthenticationSettingsContract(generators)
	oAuth2AuthenticationSettingsContractGenerator = gen.Struct(reflect.TypeOf(OAuth2AuthenticationSettingsContract{}), generators)

	return oAuth2AuthenticationSettingsContractGenerator
}

// AddIndependentPropertyGeneratorsForOAuth2AuthenticationSettingsContract is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForOAuth2AuthenticationSettingsContract(gens map[string]gopter.Gen) {
	gens["AuthorizationServerId"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.AlphaString())
}

func Test_OpenIdAuthenticationSettingsContract_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of OpenIdAuthenticationSettingsContract via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForOpenIdAuthenticationSettingsContract, OpenIdAuthenticationSettingsContractGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForOpenIdAuthenticationSettingsContract runs a test to see if a specific instance of OpenIdAuthenticationSettingsContract round trips to JSON and back losslessly
func RunJSONSerializationTestForOpenIdAuthenticationSettingsContract(subject OpenIdAuthenticationSettingsContract) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual OpenIdAuthenticationSettingsContract
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

// Generator of OpenIdAuthenticationSettingsContract instances for property testing - lazily instantiated by
// OpenIdAuthenticationSettingsContractGenerator()
var openIdAuthenticationSettingsContractGenerator gopter.Gen

// OpenIdAuthenticationSettingsContractGenerator returns a generator of OpenIdAuthenticationSettingsContract instances for property testing.
func OpenIdAuthenticationSettingsContractGenerator() gopter.Gen {
	if openIdAuthenticationSettingsContractGenerator != nil {
		return openIdAuthenticationSettingsContractGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOpenIdAuthenticationSettingsContract(generators)
	openIdAuthenticationSettingsContractGenerator = gen.Struct(reflect.TypeOf(OpenIdAuthenticationSettingsContract{}), generators)

	return openIdAuthenticationSettingsContractGenerator
}

// AddIndependentPropertyGeneratorsForOpenIdAuthenticationSettingsContract is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForOpenIdAuthenticationSettingsContract(gens map[string]gopter.Gen) {
	gens["BearerTokenSendingMethods"] = gen.SliceOf(gen.OneConstOf(BearerTokenSendingMethodsContract_AuthorizationHeader, BearerTokenSendingMethodsContract_Query))
	gens["OpenidProviderId"] = gen.PtrOf(gen.AlphaString())
}

func Test_SubscriptionKeyParameterNamesContract_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubscriptionKeyParameterNamesContract via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubscriptionKeyParameterNamesContract, SubscriptionKeyParameterNamesContractGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubscriptionKeyParameterNamesContract runs a test to see if a specific instance of SubscriptionKeyParameterNamesContract round trips to JSON and back losslessly
func RunJSONSerializationTestForSubscriptionKeyParameterNamesContract(subject SubscriptionKeyParameterNamesContract) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubscriptionKeyParameterNamesContract
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

// Generator of SubscriptionKeyParameterNamesContract instances for property testing - lazily instantiated by
// SubscriptionKeyParameterNamesContractGenerator()
var subscriptionKeyParameterNamesContractGenerator gopter.Gen

// SubscriptionKeyParameterNamesContractGenerator returns a generator of SubscriptionKeyParameterNamesContract instances for property testing.
func SubscriptionKeyParameterNamesContractGenerator() gopter.Gen {
	if subscriptionKeyParameterNamesContractGenerator != nil {
		return subscriptionKeyParameterNamesContractGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubscriptionKeyParameterNamesContract(generators)
	subscriptionKeyParameterNamesContractGenerator = gen.Struct(reflect.TypeOf(SubscriptionKeyParameterNamesContract{}), generators)

	return subscriptionKeyParameterNamesContractGenerator
}

// AddIndependentPropertyGeneratorsForSubscriptionKeyParameterNamesContract is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubscriptionKeyParameterNamesContract(gens map[string]gopter.Gen) {
	gens["Header"] = gen.PtrOf(gen.AlphaString())
	gens["Query"] = gen.PtrOf(gen.AlphaString())
}
