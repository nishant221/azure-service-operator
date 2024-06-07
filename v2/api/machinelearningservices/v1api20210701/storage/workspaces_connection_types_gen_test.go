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

func Test_WorkspacesConnection_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspacesConnection via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspacesConnection, WorkspacesConnectionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspacesConnection runs a test to see if a specific instance of WorkspacesConnection round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspacesConnection(subject WorkspacesConnection) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspacesConnection
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

// Generator of WorkspacesConnection instances for property testing - lazily instantiated by
// WorkspacesConnectionGenerator()
var workspacesConnectionGenerator gopter.Gen

// WorkspacesConnectionGenerator returns a generator of WorkspacesConnection instances for property testing.
func WorkspacesConnectionGenerator() gopter.Gen {
	if workspacesConnectionGenerator != nil {
		return workspacesConnectionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForWorkspacesConnection(generators)
	workspacesConnectionGenerator = gen.Struct(reflect.TypeOf(WorkspacesConnection{}), generators)

	return workspacesConnectionGenerator
}

// AddRelatedPropertyGeneratorsForWorkspacesConnection is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspacesConnection(gens map[string]gopter.Gen) {
	gens["Spec"] = Workspaces_Connection_SpecGenerator()
	gens["Status"] = Workspaces_Connection_STATUSGenerator()
}

func Test_Workspaces_Connection_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Workspaces_Connection_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaces_Connection_STATUS, Workspaces_Connection_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaces_Connection_STATUS runs a test to see if a specific instance of Workspaces_Connection_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaces_Connection_STATUS(subject Workspaces_Connection_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Workspaces_Connection_STATUS
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

// Generator of Workspaces_Connection_STATUS instances for property testing - lazily instantiated by
// Workspaces_Connection_STATUSGenerator()
var workspaces_Connection_STATUSGenerator gopter.Gen

// Workspaces_Connection_STATUSGenerator returns a generator of Workspaces_Connection_STATUS instances for property testing.
func Workspaces_Connection_STATUSGenerator() gopter.Gen {
	if workspaces_Connection_STATUSGenerator != nil {
		return workspaces_Connection_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaces_Connection_STATUS(generators)
	workspaces_Connection_STATUSGenerator = gen.Struct(reflect.TypeOf(Workspaces_Connection_STATUS{}), generators)

	return workspaces_Connection_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaces_Connection_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaces_Connection_STATUS(gens map[string]gopter.Gen) {
	gens["AuthType"] = gen.PtrOf(gen.AlphaString())
	gens["Category"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
	gens["ValueFormat"] = gen.PtrOf(gen.AlphaString())
}

func Test_Workspaces_Connection_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Workspaces_Connection_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaces_Connection_Spec, Workspaces_Connection_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaces_Connection_Spec runs a test to see if a specific instance of Workspaces_Connection_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaces_Connection_Spec(subject Workspaces_Connection_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Workspaces_Connection_Spec
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

// Generator of Workspaces_Connection_Spec instances for property testing - lazily instantiated by
// Workspaces_Connection_SpecGenerator()
var workspaces_Connection_SpecGenerator gopter.Gen

// Workspaces_Connection_SpecGenerator returns a generator of Workspaces_Connection_Spec instances for property testing.
func Workspaces_Connection_SpecGenerator() gopter.Gen {
	if workspaces_Connection_SpecGenerator != nil {
		return workspaces_Connection_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaces_Connection_Spec(generators)
	workspaces_Connection_SpecGenerator = gen.Struct(reflect.TypeOf(Workspaces_Connection_Spec{}), generators)

	return workspaces_Connection_SpecGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaces_Connection_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaces_Connection_Spec(gens map[string]gopter.Gen) {
	gens["AuthType"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["Category"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Target"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
	gens["ValueFormat"] = gen.PtrOf(gen.AlphaString())
}
