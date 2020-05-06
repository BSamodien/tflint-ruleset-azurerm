// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermFirewallApplicationRuleCollectionInvalidActionRule checks the pattern is valid
type AzurermFirewallApplicationRuleCollectionInvalidActionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermFirewallApplicationRuleCollectionInvalidActionRule returns new rule with default attributes
func NewAzurermFirewallApplicationRuleCollectionInvalidActionRule() *AzurermFirewallApplicationRuleCollectionInvalidActionRule {
	return &AzurermFirewallApplicationRuleCollectionInvalidActionRule{
		resourceType:  "azurerm_firewall_application_rule_collection",
		attributeName: "action",
		enum: []string{
			"Allow",
			"Deny",
		},
	}
}

// Name returns the rule name
func (r *AzurermFirewallApplicationRuleCollectionInvalidActionRule) Name() string {
	return "azurerm_firewall_application_rule_collection_invalid_action"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermFirewallApplicationRuleCollectionInvalidActionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermFirewallApplicationRuleCollectionInvalidActionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermFirewallApplicationRuleCollectionInvalidActionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermFirewallApplicationRuleCollectionInvalidActionRule) Check(runner tflint.Runner) error {
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
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as action`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
