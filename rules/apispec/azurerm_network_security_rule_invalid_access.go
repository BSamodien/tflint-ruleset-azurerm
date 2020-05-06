// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermNetworkSecurityRuleInvalidAccessRule checks the pattern is valid
type AzurermNetworkSecurityRuleInvalidAccessRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermNetworkSecurityRuleInvalidAccessRule returns new rule with default attributes
func NewAzurermNetworkSecurityRuleInvalidAccessRule() *AzurermNetworkSecurityRuleInvalidAccessRule {
	return &AzurermNetworkSecurityRuleInvalidAccessRule{
		resourceType:  "azurerm_network_security_rule",
		attributeName: "access",
		enum: []string{
			"Allow",
			"Deny",
		},
	}
}

// Name returns the rule name
func (r *AzurermNetworkSecurityRuleInvalidAccessRule) Name() string {
	return "azurerm_network_security_rule_invalid_access"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermNetworkSecurityRuleInvalidAccessRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermNetworkSecurityRuleInvalidAccessRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermNetworkSecurityRuleInvalidAccessRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermNetworkSecurityRuleInvalidAccessRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as access`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
