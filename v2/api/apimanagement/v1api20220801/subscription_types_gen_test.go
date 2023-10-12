// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

import (
	"encoding/json"
	v20220801s "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801storage"
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

func Test_Subscription_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Subscription to hub returns original",
		prop.ForAll(RunResourceConversionTestForSubscription, SubscriptionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSubscription tests if a specific instance of Subscription round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSubscription(subject Subscription) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20220801s.Subscription
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Subscription
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Subscription_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Subscription to Subscription via AssignProperties_To_Subscription & AssignProperties_From_Subscription returns original",
		prop.ForAll(RunPropertyAssignmentTestForSubscription, SubscriptionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSubscription tests if a specific instance of Subscription can be assigned to v1api20220801storage and back losslessly
func RunPropertyAssignmentTestForSubscription(subject Subscription) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220801s.Subscription
	err := copied.AssignProperties_To_Subscription(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Subscription
	err = actual.AssignProperties_From_Subscription(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Subscription_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Subscription via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubscription, SubscriptionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubscription runs a test to see if a specific instance of Subscription round trips to JSON and back losslessly
func RunJSONSerializationTestForSubscription(subject Subscription) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Subscription
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

// Generator of Subscription instances for property testing - lazily instantiated by SubscriptionGenerator()
var subscriptionGenerator gopter.Gen

// SubscriptionGenerator returns a generator of Subscription instances for property testing.
func SubscriptionGenerator() gopter.Gen {
	if subscriptionGenerator != nil {
		return subscriptionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSubscription(generators)
	subscriptionGenerator = gen.Struct(reflect.TypeOf(Subscription{}), generators)

	return subscriptionGenerator
}

// AddRelatedPropertyGeneratorsForSubscription is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSubscription(gens map[string]gopter.Gen) {
	gens["Spec"] = Service_Subscription_SpecGenerator()
	gens["Status"] = Service_Subscription_STATUSGenerator()
}

func Test_Service_Subscription_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Service_Subscription_Spec to Service_Subscription_Spec via AssignProperties_To_Service_Subscription_Spec & AssignProperties_From_Service_Subscription_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForService_Subscription_Spec, Service_Subscription_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForService_Subscription_Spec tests if a specific instance of Service_Subscription_Spec can be assigned to v1api20220801storage and back losslessly
func RunPropertyAssignmentTestForService_Subscription_Spec(subject Service_Subscription_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220801s.Service_Subscription_Spec
	err := copied.AssignProperties_To_Service_Subscription_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Service_Subscription_Spec
	err = actual.AssignProperties_From_Service_Subscription_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Service_Subscription_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_Subscription_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_Subscription_Spec, Service_Subscription_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_Subscription_Spec runs a test to see if a specific instance of Service_Subscription_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForService_Subscription_Spec(subject Service_Subscription_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_Subscription_Spec
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

// Generator of Service_Subscription_Spec instances for property testing - lazily instantiated by
// Service_Subscription_SpecGenerator()
var service_Subscription_SpecGenerator gopter.Gen

// Service_Subscription_SpecGenerator returns a generator of Service_Subscription_Spec instances for property testing.
func Service_Subscription_SpecGenerator() gopter.Gen {
	if service_Subscription_SpecGenerator != nil {
		return service_Subscription_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_Subscription_Spec(generators)
	service_Subscription_SpecGenerator = gen.Struct(reflect.TypeOf(Service_Subscription_Spec{}), generators)

	return service_Subscription_SpecGenerator
}

// AddIndependentPropertyGeneratorsForService_Subscription_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_Subscription_Spec(gens map[string]gopter.Gen) {
	gens["AllowTracing"] = gen.PtrOf(gen.Bool())
	gens["AzureName"] = gen.AlphaString()
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(
		SubscriptionCreateParameterProperties_State_Active,
		SubscriptionCreateParameterProperties_State_Cancelled,
		SubscriptionCreateParameterProperties_State_Expired,
		SubscriptionCreateParameterProperties_State_Rejected,
		SubscriptionCreateParameterProperties_State_Submitted,
		SubscriptionCreateParameterProperties_State_Suspended))
}

func Test_Service_Subscription_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Service_Subscription_STATUS to Service_Subscription_STATUS via AssignProperties_To_Service_Subscription_STATUS & AssignProperties_From_Service_Subscription_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForService_Subscription_STATUS, Service_Subscription_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForService_Subscription_STATUS tests if a specific instance of Service_Subscription_STATUS can be assigned to v1api20220801storage and back losslessly
func RunPropertyAssignmentTestForService_Subscription_STATUS(subject Service_Subscription_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220801s.Service_Subscription_STATUS
	err := copied.AssignProperties_To_Service_Subscription_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Service_Subscription_STATUS
	err = actual.AssignProperties_From_Service_Subscription_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Service_Subscription_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_Subscription_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_Subscription_STATUS, Service_Subscription_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_Subscription_STATUS runs a test to see if a specific instance of Service_Subscription_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForService_Subscription_STATUS(subject Service_Subscription_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_Subscription_STATUS
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

// Generator of Service_Subscription_STATUS instances for property testing - lazily instantiated by
// Service_Subscription_STATUSGenerator()
var service_Subscription_STATUSGenerator gopter.Gen

// Service_Subscription_STATUSGenerator returns a generator of Service_Subscription_STATUS instances for property testing.
func Service_Subscription_STATUSGenerator() gopter.Gen {
	if service_Subscription_STATUSGenerator != nil {
		return service_Subscription_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_Subscription_STATUS(generators)
	service_Subscription_STATUSGenerator = gen.Struct(reflect.TypeOf(Service_Subscription_STATUS{}), generators)

	return service_Subscription_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForService_Subscription_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_Subscription_STATUS(gens map[string]gopter.Gen) {
	gens["AllowTracing"] = gen.PtrOf(gen.Bool())
	gens["CreatedDate"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["EndDate"] = gen.PtrOf(gen.AlphaString())
	gens["ExpirationDate"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["NotificationDate"] = gen.PtrOf(gen.AlphaString())
	gens["OwnerId"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.AlphaString())
	gens["StartDate"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(
		SubscriptionContractProperties_State_STATUS_Active,
		SubscriptionContractProperties_State_STATUS_Cancelled,
		SubscriptionContractProperties_State_STATUS_Expired,
		SubscriptionContractProperties_State_STATUS_Rejected,
		SubscriptionContractProperties_State_STATUS_Submitted,
		SubscriptionContractProperties_State_STATUS_Suspended))
	gens["StateComment"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}