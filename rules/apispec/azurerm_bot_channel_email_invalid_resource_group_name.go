// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermBotChannelEmailInvalidResourceGroupNameRule checks the pattern is valid
type AzurermBotChannelEmailInvalidResourceGroupNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermBotChannelEmailInvalidResourceGroupNameRule returns new rule with default attributes
func NewAzurermBotChannelEmailInvalidResourceGroupNameRule() *AzurermBotChannelEmailInvalidResourceGroupNameRule {
	return &AzurermBotChannelEmailInvalidResourceGroupNameRule{
		resourceType:  "azurerm_bot_channel_email",
		attributeName: "resource_group_name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`),
	}
}

// Name returns the rule name
func (r *AzurermBotChannelEmailInvalidResourceGroupNameRule) Name() string {
	return "azurerm_bot_channel_email_invalid_resource_group_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermBotChannelEmailInvalidResourceGroupNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermBotChannelEmailInvalidResourceGroupNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermBotChannelEmailInvalidResourceGroupNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermBotChannelEmailInvalidResourceGroupNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
