// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermSentinelAlertRuleMsSecurityIncidentInvalidProductFilterRule checks the pattern is valid
type AzurermSentinelAlertRuleMsSecurityIncidentInvalidProductFilterRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermSentinelAlertRuleMsSecurityIncidentInvalidProductFilterRule returns new rule with default attributes
func NewAzurermSentinelAlertRuleMsSecurityIncidentInvalidProductFilterRule() *AzurermSentinelAlertRuleMsSecurityIncidentInvalidProductFilterRule {
	return &AzurermSentinelAlertRuleMsSecurityIncidentInvalidProductFilterRule{
		resourceType:  "azurerm_sentinel_alert_rule_ms_security_incident",
		attributeName: "product_filter",
		enum: []string{
			"Microsoft Cloud App Security",
			"Azure Security Center",
			"Azure Advanced Threat Protection",
			"Azure Active Directory Identity Protection",
			"Azure Security Center for IoT",
		},
	}
}

// Name returns the rule name
func (r *AzurermSentinelAlertRuleMsSecurityIncidentInvalidProductFilterRule) Name() string {
	return "azurerm_sentinel_alert_rule_ms_security_incident_invalid_product_filter"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermSentinelAlertRuleMsSecurityIncidentInvalidProductFilterRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermSentinelAlertRuleMsSecurityIncidentInvalidProductFilterRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermSentinelAlertRuleMsSecurityIncidentInvalidProductFilterRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermSentinelAlertRuleMsSecurityIncidentInvalidProductFilterRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as product_filter`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
