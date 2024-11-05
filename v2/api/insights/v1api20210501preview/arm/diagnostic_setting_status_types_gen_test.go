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

func Test_DiagnosticSetting_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiagnosticSetting_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiagnosticSetting_STATUS, DiagnosticSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiagnosticSetting_STATUS runs a test to see if a specific instance of DiagnosticSetting_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDiagnosticSetting_STATUS(subject DiagnosticSetting_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiagnosticSetting_STATUS
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

// Generator of DiagnosticSetting_STATUS instances for property testing - lazily instantiated by
// DiagnosticSetting_STATUSGenerator()
var diagnosticSetting_STATUSGenerator gopter.Gen

// DiagnosticSetting_STATUSGenerator returns a generator of DiagnosticSetting_STATUS instances for property testing.
// We first initialize diagnosticSetting_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DiagnosticSetting_STATUSGenerator() gopter.Gen {
	if diagnosticSetting_STATUSGenerator != nil {
		return diagnosticSetting_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiagnosticSetting_STATUS(generators)
	diagnosticSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(DiagnosticSetting_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiagnosticSetting_STATUS(generators)
	AddRelatedPropertyGeneratorsForDiagnosticSetting_STATUS(generators)
	diagnosticSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(DiagnosticSetting_STATUS{}), generators)

	return diagnosticSetting_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDiagnosticSetting_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiagnosticSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDiagnosticSetting_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiagnosticSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DiagnosticSettings_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_DiagnosticSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiagnosticSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiagnosticSettings_STATUS, DiagnosticSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiagnosticSettings_STATUS runs a test to see if a specific instance of DiagnosticSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDiagnosticSettings_STATUS(subject DiagnosticSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiagnosticSettings_STATUS
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

// Generator of DiagnosticSettings_STATUS instances for property testing - lazily instantiated by
// DiagnosticSettings_STATUSGenerator()
var diagnosticSettings_STATUSGenerator gopter.Gen

// DiagnosticSettings_STATUSGenerator returns a generator of DiagnosticSettings_STATUS instances for property testing.
// We first initialize diagnosticSettings_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DiagnosticSettings_STATUSGenerator() gopter.Gen {
	if diagnosticSettings_STATUSGenerator != nil {
		return diagnosticSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiagnosticSettings_STATUS(generators)
	diagnosticSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(DiagnosticSettings_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiagnosticSettings_STATUS(generators)
	AddRelatedPropertyGeneratorsForDiagnosticSettings_STATUS(generators)
	diagnosticSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(DiagnosticSettings_STATUS{}), generators)

	return diagnosticSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDiagnosticSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiagnosticSettings_STATUS(gens map[string]gopter.Gen) {
	gens["EventHubAuthorizationRuleId"] = gen.PtrOf(gen.AlphaString())
	gens["EventHubName"] = gen.PtrOf(gen.AlphaString())
	gens["LogAnalyticsDestinationType"] = gen.PtrOf(gen.AlphaString())
	gens["MarketplacePartnerId"] = gen.PtrOf(gen.AlphaString())
	gens["ServiceBusRuleId"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountId"] = gen.PtrOf(gen.AlphaString())
	gens["WorkspaceId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDiagnosticSettings_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiagnosticSettings_STATUS(gens map[string]gopter.Gen) {
	gens["Logs"] = gen.SliceOf(LogSettings_STATUSGenerator())
	gens["Metrics"] = gen.SliceOf(MetricSettings_STATUSGenerator())
}

func Test_LogSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LogSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLogSettings_STATUS, LogSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLogSettings_STATUS runs a test to see if a specific instance of LogSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForLogSettings_STATUS(subject LogSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LogSettings_STATUS
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

// Generator of LogSettings_STATUS instances for property testing - lazily instantiated by LogSettings_STATUSGenerator()
var logSettings_STATUSGenerator gopter.Gen

// LogSettings_STATUSGenerator returns a generator of LogSettings_STATUS instances for property testing.
// We first initialize logSettings_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LogSettings_STATUSGenerator() gopter.Gen {
	if logSettings_STATUSGenerator != nil {
		return logSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLogSettings_STATUS(generators)
	logSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(LogSettings_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLogSettings_STATUS(generators)
	AddRelatedPropertyGeneratorsForLogSettings_STATUS(generators)
	logSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(LogSettings_STATUS{}), generators)

	return logSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForLogSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLogSettings_STATUS(gens map[string]gopter.Gen) {
	gens["Category"] = gen.PtrOf(gen.AlphaString())
	gens["CategoryGroup"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForLogSettings_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLogSettings_STATUS(gens map[string]gopter.Gen) {
	gens["RetentionPolicy"] = gen.PtrOf(RetentionPolicy_STATUSGenerator())
}

func Test_MetricSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricSettings_STATUS, MetricSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricSettings_STATUS runs a test to see if a specific instance of MetricSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricSettings_STATUS(subject MetricSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricSettings_STATUS
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

// Generator of MetricSettings_STATUS instances for property testing - lazily instantiated by
// MetricSettings_STATUSGenerator()
var metricSettings_STATUSGenerator gopter.Gen

// MetricSettings_STATUSGenerator returns a generator of MetricSettings_STATUS instances for property testing.
// We first initialize metricSettings_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MetricSettings_STATUSGenerator() gopter.Gen {
	if metricSettings_STATUSGenerator != nil {
		return metricSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricSettings_STATUS(generators)
	metricSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(MetricSettings_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricSettings_STATUS(generators)
	AddRelatedPropertyGeneratorsForMetricSettings_STATUS(generators)
	metricSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(MetricSettings_STATUS{}), generators)

	return metricSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMetricSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricSettings_STATUS(gens map[string]gopter.Gen) {
	gens["Category"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["TimeGrain"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMetricSettings_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMetricSettings_STATUS(gens map[string]gopter.Gen) {
	gens["RetentionPolicy"] = gen.PtrOf(RetentionPolicy_STATUSGenerator())
}

func Test_RetentionPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RetentionPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRetentionPolicy_STATUS, RetentionPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRetentionPolicy_STATUS runs a test to see if a specific instance of RetentionPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRetentionPolicy_STATUS(subject RetentionPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RetentionPolicy_STATUS
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

// Generator of RetentionPolicy_STATUS instances for property testing - lazily instantiated by
// RetentionPolicy_STATUSGenerator()
var retentionPolicy_STATUSGenerator gopter.Gen

// RetentionPolicy_STATUSGenerator returns a generator of RetentionPolicy_STATUS instances for property testing.
func RetentionPolicy_STATUSGenerator() gopter.Gen {
	if retentionPolicy_STATUSGenerator != nil {
		return retentionPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRetentionPolicy_STATUS(generators)
	retentionPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(RetentionPolicy_STATUS{}), generators)

	return retentionPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRetentionPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRetentionPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_SystemData_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS, SystemData_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS runs a test to see if a specific instance of SystemData_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS(subject SystemData_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS
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

// Generator of SystemData_STATUS instances for property testing - lazily instantiated by SystemData_STATUSGenerator()
var systemData_STATUSGenerator gopter.Gen

// SystemData_STATUSGenerator returns a generator of SystemData_STATUS instances for property testing.
func SystemData_STATUSGenerator() gopter.Gen {
	if systemData_STATUSGenerator != nil {
		return systemData_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS(generators)
	systemData_STATUSGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS{}), generators)

	return systemData_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_CreatedByType_STATUS_Application,
		SystemData_CreatedByType_STATUS_Key,
		SystemData_CreatedByType_STATUS_ManagedIdentity,
		SystemData_CreatedByType_STATUS_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_STATUS_Application,
		SystemData_LastModifiedByType_STATUS_Key,
		SystemData_LastModifiedByType_STATUS_ManagedIdentity,
		SystemData_LastModifiedByType_STATUS_User))
}