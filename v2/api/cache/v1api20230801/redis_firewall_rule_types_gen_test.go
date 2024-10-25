// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230801

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/cache/v1api20230801/storage"
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

func Test_RedisFirewallRule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisFirewallRule to hub returns original",
		prop.ForAll(RunResourceConversionTestForRedisFirewallRule, RedisFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForRedisFirewallRule tests if a specific instance of RedisFirewallRule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForRedisFirewallRule(subject RedisFirewallRule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.RedisFirewallRule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual RedisFirewallRule
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

func Test_RedisFirewallRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisFirewallRule to RedisFirewallRule via AssignProperties_To_RedisFirewallRule & AssignProperties_From_RedisFirewallRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisFirewallRule, RedisFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisFirewallRule tests if a specific instance of RedisFirewallRule can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForRedisFirewallRule(subject RedisFirewallRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.RedisFirewallRule
	err := copied.AssignProperties_To_RedisFirewallRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisFirewallRule
	err = actual.AssignProperties_From_RedisFirewallRule(&other)
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

func Test_RedisFirewallRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisFirewallRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisFirewallRule, RedisFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisFirewallRule runs a test to see if a specific instance of RedisFirewallRule round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisFirewallRule(subject RedisFirewallRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisFirewallRule
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

// Generator of RedisFirewallRule instances for property testing - lazily instantiated by RedisFirewallRuleGenerator()
var redisFirewallRuleGenerator gopter.Gen

// RedisFirewallRuleGenerator returns a generator of RedisFirewallRule instances for property testing.
func RedisFirewallRuleGenerator() gopter.Gen {
	if redisFirewallRuleGenerator != nil {
		return redisFirewallRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisFirewallRule(generators)
	redisFirewallRuleGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRule{}), generators)

	return redisFirewallRuleGenerator
}

// AddRelatedPropertyGeneratorsForRedisFirewallRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisFirewallRule(gens map[string]gopter.Gen) {
	gens["Spec"] = RedisFirewallRule_SpecGenerator()
	gens["Status"] = RedisFirewallRule_STATUSGenerator()
}

func Test_RedisFirewallRule_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisFirewallRule_STATUS to RedisFirewallRule_STATUS via AssignProperties_To_RedisFirewallRule_STATUS & AssignProperties_From_RedisFirewallRule_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisFirewallRule_STATUS, RedisFirewallRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisFirewallRule_STATUS tests if a specific instance of RedisFirewallRule_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForRedisFirewallRule_STATUS(subject RedisFirewallRule_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.RedisFirewallRule_STATUS
	err := copied.AssignProperties_To_RedisFirewallRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisFirewallRule_STATUS
	err = actual.AssignProperties_From_RedisFirewallRule_STATUS(&other)
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

func Test_RedisFirewallRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisFirewallRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisFirewallRule_STATUS, RedisFirewallRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisFirewallRule_STATUS runs a test to see if a specific instance of RedisFirewallRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisFirewallRule_STATUS(subject RedisFirewallRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisFirewallRule_STATUS
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

// Generator of RedisFirewallRule_STATUS instances for property testing - lazily instantiated by
// RedisFirewallRule_STATUSGenerator()
var redisFirewallRule_STATUSGenerator gopter.Gen

// RedisFirewallRule_STATUSGenerator returns a generator of RedisFirewallRule_STATUS instances for property testing.
func RedisFirewallRule_STATUSGenerator() gopter.Gen {
	if redisFirewallRule_STATUSGenerator != nil {
		return redisFirewallRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisFirewallRule_STATUS(generators)
	redisFirewallRule_STATUSGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRule_STATUS{}), generators)

	return redisFirewallRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedisFirewallRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisFirewallRule_STATUS(gens map[string]gopter.Gen) {
	gens["EndIP"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["StartIP"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisFirewallRule_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisFirewallRule_Spec to RedisFirewallRule_Spec via AssignProperties_To_RedisFirewallRule_Spec & AssignProperties_From_RedisFirewallRule_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisFirewallRule_Spec, RedisFirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisFirewallRule_Spec tests if a specific instance of RedisFirewallRule_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForRedisFirewallRule_Spec(subject RedisFirewallRule_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.RedisFirewallRule_Spec
	err := copied.AssignProperties_To_RedisFirewallRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisFirewallRule_Spec
	err = actual.AssignProperties_From_RedisFirewallRule_Spec(&other)
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

func Test_RedisFirewallRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisFirewallRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisFirewallRule_Spec, RedisFirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisFirewallRule_Spec runs a test to see if a specific instance of RedisFirewallRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisFirewallRule_Spec(subject RedisFirewallRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisFirewallRule_Spec
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

// Generator of RedisFirewallRule_Spec instances for property testing - lazily instantiated by
// RedisFirewallRule_SpecGenerator()
var redisFirewallRule_SpecGenerator gopter.Gen

// RedisFirewallRule_SpecGenerator returns a generator of RedisFirewallRule_Spec instances for property testing.
func RedisFirewallRule_SpecGenerator() gopter.Gen {
	if redisFirewallRule_SpecGenerator != nil {
		return redisFirewallRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisFirewallRule_Spec(generators)
	redisFirewallRule_SpecGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRule_Spec{}), generators)

	return redisFirewallRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisFirewallRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisFirewallRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EndIP"] = gen.PtrOf(gen.AlphaString())
	gens["StartIP"] = gen.PtrOf(gen.AlphaString())
}