// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermSecurityCenterSubscriptionPricingInvalidTierRule checks the pattern is valid
type AzurermSecurityCenterSubscriptionPricingInvalidTierRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermSecurityCenterSubscriptionPricingInvalidTierRule returns new rule with default attributes
func NewAzurermSecurityCenterSubscriptionPricingInvalidTierRule() *AzurermSecurityCenterSubscriptionPricingInvalidTierRule {
	return &AzurermSecurityCenterSubscriptionPricingInvalidTierRule{
		resourceType:  "azurerm_security_center_subscription_pricing",
		attributeName: "tier",
		enum: []string{
			"Free",
			"Standard",
		},
	}
}

// Name returns the rule name
func (r *AzurermSecurityCenterSubscriptionPricingInvalidTierRule) Name() string {
	return "azurerm_security_center_subscription_pricing_invalid_tier"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermSecurityCenterSubscriptionPricingInvalidTierRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermSecurityCenterSubscriptionPricingInvalidTierRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermSecurityCenterSubscriptionPricingInvalidTierRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermSecurityCenterSubscriptionPricingInvalidTierRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as tier`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
