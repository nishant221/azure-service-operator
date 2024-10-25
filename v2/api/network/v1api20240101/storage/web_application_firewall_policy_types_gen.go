// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=webapplicationfirewallpolicies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={webapplicationfirewallpolicies/status,webapplicationfirewallpolicies/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20240101.WebApplicationFirewallPolicy
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2024-01-01/webapplicationfirewall.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}
type WebApplicationFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebApplicationFirewallPolicy_Spec   `json:"spec,omitempty"`
	Status            WebApplicationFirewallPolicy_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &WebApplicationFirewallPolicy{}

// GetConditions returns the conditions of the resource
func (policy *WebApplicationFirewallPolicy) GetConditions() conditions.Conditions {
	return policy.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (policy *WebApplicationFirewallPolicy) SetConditions(conditions conditions.Conditions) {
	policy.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &WebApplicationFirewallPolicy{}

// AzureName returns the Azure name of the resource
func (policy *WebApplicationFirewallPolicy) AzureName() string {
	return policy.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-01-01"
func (policy WebApplicationFirewallPolicy) GetAPIVersion() string {
	return "2024-01-01"
}

// GetResourceScope returns the scope of the resource
func (policy *WebApplicationFirewallPolicy) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (policy *WebApplicationFirewallPolicy) GetSpec() genruntime.ConvertibleSpec {
	return &policy.Spec
}

// GetStatus returns the status of this resource
func (policy *WebApplicationFirewallPolicy) GetStatus() genruntime.ConvertibleStatus {
	return &policy.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (policy *WebApplicationFirewallPolicy) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies"
func (policy *WebApplicationFirewallPolicy) GetType() string {
	return "Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies"
}

// NewEmptyStatus returns a new empty (blank) status
func (policy *WebApplicationFirewallPolicy) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &WebApplicationFirewallPolicy_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (policy *WebApplicationFirewallPolicy) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(policy.Spec)
	return policy.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (policy *WebApplicationFirewallPolicy) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*WebApplicationFirewallPolicy_STATUS); ok {
		policy.Status = *st
		return nil
	}

	// Convert status to required version
	var st WebApplicationFirewallPolicy_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// Hub marks that this WebApplicationFirewallPolicy is the hub type for conversion
func (policy *WebApplicationFirewallPolicy) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *WebApplicationFirewallPolicy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion,
		Kind:    "WebApplicationFirewallPolicy",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20240101.WebApplicationFirewallPolicy
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2024-01-01/webapplicationfirewall.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}
type WebApplicationFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebApplicationFirewallPolicy `json:"items"`
}

// Storage version of v1api20240101.WebApplicationFirewallPolicy_Spec
type WebApplicationFirewallPolicy_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                             `json:"azureName,omitempty"`
	CustomRules     []WebApplicationFirewallCustomRule `json:"customRules,omitempty"`
	Location        *string                            `json:"location,omitempty"`
	ManagedRules    *ManagedRulesDefinition            `json:"managedRules,omitempty"`
	OriginalVersion string                             `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner          *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PolicySettings *PolicySettings                    `json:"policySettings,omitempty"`
	PropertyBag    genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags           map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &WebApplicationFirewallPolicy_Spec{}

// ConvertSpecFrom populates our WebApplicationFirewallPolicy_Spec from the provided source
func (policy *WebApplicationFirewallPolicy_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(policy)
}

// ConvertSpecTo populates the provided destination from our WebApplicationFirewallPolicy_Spec
func (policy *WebApplicationFirewallPolicy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(policy)
}

// Storage version of v1api20240101.WebApplicationFirewallPolicy_STATUS
type WebApplicationFirewallPolicy_STATUS struct {
	ApplicationGateways []ApplicationGateway_STATUS_ApplicationGatewayWebApplicationFirewallPolicy_SubResourceEmbedded `json:"applicationGateways,omitempty"`
	Conditions          []conditions.Condition                                                                         `json:"conditions,omitempty"`
	CustomRules         []WebApplicationFirewallCustomRule_STATUS                                                      `json:"customRules,omitempty"`
	Etag                *string                                                                                        `json:"etag,omitempty"`
	HttpListeners       []SubResource_STATUS                                                                           `json:"httpListeners,omitempty"`
	Id                  *string                                                                                        `json:"id,omitempty"`
	Location            *string                                                                                        `json:"location,omitempty"`
	ManagedRules        *ManagedRulesDefinition_STATUS                                                                 `json:"managedRules,omitempty"`
	Name                *string                                                                                        `json:"name,omitempty"`
	PathBasedRules      []SubResource_STATUS                                                                           `json:"pathBasedRules,omitempty"`
	PolicySettings      *PolicySettings_STATUS                                                                         `json:"policySettings,omitempty"`
	PropertyBag         genruntime.PropertyBag                                                                         `json:"$propertyBag,omitempty"`
	ProvisioningState   *string                                                                                        `json:"provisioningState,omitempty"`
	ResourceState       *string                                                                                        `json:"resourceState,omitempty"`
	Tags                map[string]string                                                                              `json:"tags,omitempty"`
	Type                *string                                                                                        `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &WebApplicationFirewallPolicy_STATUS{}

// ConvertStatusFrom populates our WebApplicationFirewallPolicy_STATUS from the provided source
func (policy *WebApplicationFirewallPolicy_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(policy)
}

// ConvertStatusTo populates the provided destination from our WebApplicationFirewallPolicy_STATUS
func (policy *WebApplicationFirewallPolicy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(policy)
}

// Storage version of v1api20240101.ApplicationGateway_STATUS_ApplicationGatewayWebApplicationFirewallPolicy_SubResourceEmbedded
// Application gateway resource.
type ApplicationGateway_STATUS_ApplicationGatewayWebApplicationFirewallPolicy_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240101.ManagedRulesDefinition
// Allow to exclude some variable satisfy the condition for the WAF check.
type ManagedRulesDefinition struct {
	Exclusions      []OwaspCrsExclusionEntry `json:"exclusions,omitempty"`
	ManagedRuleSets []ManagedRuleSet         `json:"managedRuleSets,omitempty"`
	PropertyBag     genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240101.ManagedRulesDefinition_STATUS
// Allow to exclude some variable satisfy the condition for the WAF check.
type ManagedRulesDefinition_STATUS struct {
	Exclusions      []OwaspCrsExclusionEntry_STATUS `json:"exclusions,omitempty"`
	ManagedRuleSets []ManagedRuleSet_STATUS         `json:"managedRuleSets,omitempty"`
	PropertyBag     genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240101.PolicySettings
// Defines contents of a web application firewall global configuration.
type PolicySettings struct {
	CustomBlockResponseBody           *string                      `json:"customBlockResponseBody,omitempty"`
	CustomBlockResponseStatusCode     *int                         `json:"customBlockResponseStatusCode,omitempty"`
	FileUploadEnforcement             *bool                        `json:"fileUploadEnforcement,omitempty"`
	FileUploadLimitInMb               *int                         `json:"fileUploadLimitInMb,omitempty"`
	JsChallengeCookieExpirationInMins *int                         `json:"jsChallengeCookieExpirationInMins,omitempty"`
	LogScrubbing                      *PolicySettings_LogScrubbing `json:"logScrubbing,omitempty"`
	MaxRequestBodySizeInKb            *int                         `json:"maxRequestBodySizeInKb,omitempty"`
	Mode                              *string                      `json:"mode,omitempty"`
	PropertyBag                       genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	RequestBodyCheck                  *bool                        `json:"requestBodyCheck,omitempty"`
	RequestBodyEnforcement            *bool                        `json:"requestBodyEnforcement,omitempty"`
	RequestBodyInspectLimitInKB       *int                         `json:"requestBodyInspectLimitInKB,omitempty"`
	State                             *string                      `json:"state,omitempty"`
}

// Storage version of v1api20240101.PolicySettings_STATUS
// Defines contents of a web application firewall global configuration.
type PolicySettings_STATUS struct {
	CustomBlockResponseBody           *string                             `json:"customBlockResponseBody,omitempty"`
	CustomBlockResponseStatusCode     *int                                `json:"customBlockResponseStatusCode,omitempty"`
	FileUploadEnforcement             *bool                               `json:"fileUploadEnforcement,omitempty"`
	FileUploadLimitInMb               *int                                `json:"fileUploadLimitInMb,omitempty"`
	JsChallengeCookieExpirationInMins *int                                `json:"jsChallengeCookieExpirationInMins,omitempty"`
	LogScrubbing                      *PolicySettings_LogScrubbing_STATUS `json:"logScrubbing,omitempty"`
	MaxRequestBodySizeInKb            *int                                `json:"maxRequestBodySizeInKb,omitempty"`
	Mode                              *string                             `json:"mode,omitempty"`
	PropertyBag                       genruntime.PropertyBag              `json:"$propertyBag,omitempty"`
	RequestBodyCheck                  *bool                               `json:"requestBodyCheck,omitempty"`
	RequestBodyEnforcement            *bool                               `json:"requestBodyEnforcement,omitempty"`
	RequestBodyInspectLimitInKB       *int                                `json:"requestBodyInspectLimitInKB,omitempty"`
	State                             *string                             `json:"state,omitempty"`
}

// Storage version of v1api20240101.SubResource_STATUS
// Reference to another subresource.
type SubResource_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240101.WebApplicationFirewallCustomRule
// Defines contents of a web application rule.
type WebApplicationFirewallCustomRule struct {
	Action             *string                `json:"action,omitempty"`
	GroupByUserSession []GroupByUserSession   `json:"groupByUserSession,omitempty"`
	MatchConditions    []MatchCondition       `json:"matchConditions,omitempty"`
	Name               *string                `json:"name,omitempty"`
	Priority           *int                   `json:"priority,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RateLimitDuration  *string                `json:"rateLimitDuration,omitempty"`
	RateLimitThreshold *int                   `json:"rateLimitThreshold,omitempty"`
	RuleType           *string                `json:"ruleType,omitempty"`
	State              *string                `json:"state,omitempty"`
}

// Storage version of v1api20240101.WebApplicationFirewallCustomRule_STATUS
// Defines contents of a web application rule.
type WebApplicationFirewallCustomRule_STATUS struct {
	Action             *string                     `json:"action,omitempty"`
	Etag               *string                     `json:"etag,omitempty"`
	GroupByUserSession []GroupByUserSession_STATUS `json:"groupByUserSession,omitempty"`
	MatchConditions    []MatchCondition_STATUS     `json:"matchConditions,omitempty"`
	Name               *string                     `json:"name,omitempty"`
	Priority           *int                        `json:"priority,omitempty"`
	PropertyBag        genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	RateLimitDuration  *string                     `json:"rateLimitDuration,omitempty"`
	RateLimitThreshold *int                        `json:"rateLimitThreshold,omitempty"`
	RuleType           *string                     `json:"ruleType,omitempty"`
	State              *string                     `json:"state,omitempty"`
}

// Storage version of v1api20240101.GroupByUserSession
// Define user session identifier group by clauses.
type GroupByUserSession struct {
	GroupByVariables []GroupByVariable      `json:"groupByVariables,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240101.GroupByUserSession_STATUS
// Define user session identifier group by clauses.
type GroupByUserSession_STATUS struct {
	GroupByVariables []GroupByVariable_STATUS `json:"groupByVariables,omitempty"`
	PropertyBag      genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240101.ManagedRuleSet
// Defines a managed rule set.
type ManagedRuleSet struct {
	PropertyBag        genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	RuleGroupOverrides []ManagedRuleGroupOverride `json:"ruleGroupOverrides,omitempty"`
	RuleSetType        *string                    `json:"ruleSetType,omitempty"`
	RuleSetVersion     *string                    `json:"ruleSetVersion,omitempty"`
}

// Storage version of v1api20240101.ManagedRuleSet_STATUS
// Defines a managed rule set.
type ManagedRuleSet_STATUS struct {
	PropertyBag        genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	RuleGroupOverrides []ManagedRuleGroupOverride_STATUS `json:"ruleGroupOverrides,omitempty"`
	RuleSetType        *string                           `json:"ruleSetType,omitempty"`
	RuleSetVersion     *string                           `json:"ruleSetVersion,omitempty"`
}

// Storage version of v1api20240101.MatchCondition
// Define match conditions.
type MatchCondition struct {
	MatchValues      []string               `json:"matchValues,omitempty"`
	MatchVariables   []MatchVariable        `json:"matchVariables,omitempty"`
	NegationConditon *bool                  `json:"negationConditon,omitempty"`
	Operator         *string                `json:"operator,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Transforms       []string               `json:"transforms,omitempty"`
}

// Storage version of v1api20240101.MatchCondition_STATUS
// Define match conditions.
type MatchCondition_STATUS struct {
	MatchValues      []string               `json:"matchValues,omitempty"`
	MatchVariables   []MatchVariable_STATUS `json:"matchVariables,omitempty"`
	NegationConditon *bool                  `json:"negationConditon,omitempty"`
	Operator         *string                `json:"operator,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Transforms       []string               `json:"transforms,omitempty"`
}

// Storage version of v1api20240101.OwaspCrsExclusionEntry
// Allow to exclude some variable satisfy the condition for the WAF check.
type OwaspCrsExclusionEntry struct {
	ExclusionManagedRuleSets []ExclusionManagedRuleSet `json:"exclusionManagedRuleSets,omitempty"`
	MatchVariable            *string                   `json:"matchVariable,omitempty"`
	PropertyBag              genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	Selector                 *string                   `json:"selector,omitempty"`
	SelectorMatchOperator    *string                   `json:"selectorMatchOperator,omitempty"`
}

// Storage version of v1api20240101.OwaspCrsExclusionEntry_STATUS
// Allow to exclude some variable satisfy the condition for the WAF check.
type OwaspCrsExclusionEntry_STATUS struct {
	ExclusionManagedRuleSets []ExclusionManagedRuleSet_STATUS `json:"exclusionManagedRuleSets,omitempty"`
	MatchVariable            *string                          `json:"matchVariable,omitempty"`
	PropertyBag              genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	Selector                 *string                          `json:"selector,omitempty"`
	SelectorMatchOperator    *string                          `json:"selectorMatchOperator,omitempty"`
}

// Storage version of v1api20240101.PolicySettings_LogScrubbing
type PolicySettings_LogScrubbing struct {
	PropertyBag    genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	ScrubbingRules []WebApplicationFirewallScrubbingRules `json:"scrubbingRules,omitempty"`
	State          *string                                `json:"state,omitempty"`
}

// Storage version of v1api20240101.PolicySettings_LogScrubbing_STATUS
type PolicySettings_LogScrubbing_STATUS struct {
	PropertyBag    genruntime.PropertyBag                        `json:"$propertyBag,omitempty"`
	ScrubbingRules []WebApplicationFirewallScrubbingRules_STATUS `json:"scrubbingRules,omitempty"`
	State          *string                                       `json:"state,omitempty"`
}

// Storage version of v1api20240101.ExclusionManagedRuleSet
// Defines a managed rule set for Exclusions.
type ExclusionManagedRuleSet struct {
	PropertyBag    genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	RuleGroups     []ExclusionManagedRuleGroup `json:"ruleGroups,omitempty"`
	RuleSetType    *string                     `json:"ruleSetType,omitempty"`
	RuleSetVersion *string                     `json:"ruleSetVersion,omitempty"`
}

// Storage version of v1api20240101.ExclusionManagedRuleSet_STATUS
// Defines a managed rule set for Exclusions.
type ExclusionManagedRuleSet_STATUS struct {
	PropertyBag    genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RuleGroups     []ExclusionManagedRuleGroup_STATUS `json:"ruleGroups,omitempty"`
	RuleSetType    *string                            `json:"ruleSetType,omitempty"`
	RuleSetVersion *string                            `json:"ruleSetVersion,omitempty"`
}

// Storage version of v1api20240101.GroupByVariable
// Define user session group by clause variables.
type GroupByVariable struct {
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VariableName *string                `json:"variableName,omitempty"`
}

// Storage version of v1api20240101.GroupByVariable_STATUS
// Define user session group by clause variables.
type GroupByVariable_STATUS struct {
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VariableName *string                `json:"variableName,omitempty"`
}

// Storage version of v1api20240101.ManagedRuleGroupOverride
// Defines a managed rule group override setting.
type ManagedRuleGroupOverride struct {
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RuleGroupName *string                `json:"ruleGroupName,omitempty"`
	Rules         []ManagedRuleOverride  `json:"rules,omitempty"`
}

// Storage version of v1api20240101.ManagedRuleGroupOverride_STATUS
// Defines a managed rule group override setting.
type ManagedRuleGroupOverride_STATUS struct {
	PropertyBag   genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	RuleGroupName *string                      `json:"ruleGroupName,omitempty"`
	Rules         []ManagedRuleOverride_STATUS `json:"rules,omitempty"`
}

// Storage version of v1api20240101.MatchVariable
// Define match variables.
type MatchVariable struct {
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Selector     *string                `json:"selector,omitempty"`
	VariableName *string                `json:"variableName,omitempty"`
}

// Storage version of v1api20240101.MatchVariable_STATUS
// Define match variables.
type MatchVariable_STATUS struct {
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Selector     *string                `json:"selector,omitempty"`
	VariableName *string                `json:"variableName,omitempty"`
}

// Storage version of v1api20240101.WebApplicationFirewallScrubbingRules
// Allow certain variables to be scrubbed on WAF logs
type WebApplicationFirewallScrubbingRules struct {
	MatchVariable         *string                `json:"matchVariable,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Selector              *string                `json:"selector,omitempty"`
	SelectorMatchOperator *string                `json:"selectorMatchOperator,omitempty"`
	State                 *string                `json:"state,omitempty"`
}

// Storage version of v1api20240101.WebApplicationFirewallScrubbingRules_STATUS
// Allow certain variables to be scrubbed on WAF logs
type WebApplicationFirewallScrubbingRules_STATUS struct {
	MatchVariable         *string                `json:"matchVariable,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Selector              *string                `json:"selector,omitempty"`
	SelectorMatchOperator *string                `json:"selectorMatchOperator,omitempty"`
	State                 *string                `json:"state,omitempty"`
}

// Storage version of v1api20240101.ExclusionManagedRuleGroup
// Defines a managed rule group to use for exclusion.
type ExclusionManagedRuleGroup struct {
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RuleGroupName *string                `json:"ruleGroupName,omitempty"`
	Rules         []ExclusionManagedRule `json:"rules,omitempty"`
}

// Storage version of v1api20240101.ExclusionManagedRuleGroup_STATUS
// Defines a managed rule group to use for exclusion.
type ExclusionManagedRuleGroup_STATUS struct {
	PropertyBag   genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	RuleGroupName *string                       `json:"ruleGroupName,omitempty"`
	Rules         []ExclusionManagedRule_STATUS `json:"rules,omitempty"`
}

// Storage version of v1api20240101.ManagedRuleOverride
// Defines a managed rule group override setting.
type ManagedRuleOverride struct {
	Action      *string                `json:"action,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RuleId      *string                `json:"ruleId,omitempty"`
	State       *string                `json:"state,omitempty"`
}

// Storage version of v1api20240101.ManagedRuleOverride_STATUS
// Defines a managed rule group override setting.
type ManagedRuleOverride_STATUS struct {
	Action      *string                `json:"action,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RuleId      *string                `json:"ruleId,omitempty"`
	State       *string                `json:"state,omitempty"`
}

// Storage version of v1api20240101.ExclusionManagedRule
// Defines a managed rule to use for exclusion.
type ExclusionManagedRule struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RuleId      *string                `json:"ruleId,omitempty"`
}

// Storage version of v1api20240101.ExclusionManagedRule_STATUS
// Defines a managed rule to use for exclusion.
type ExclusionManagedRule_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RuleId      *string                `json:"ruleId,omitempty"`
}

func init() {
	SchemeBuilder.Register(&WebApplicationFirewallPolicy{}, &WebApplicationFirewallPolicyList{})
}