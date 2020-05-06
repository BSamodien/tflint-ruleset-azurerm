// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermDevspaceControllerInvalidNameRule checks the pattern is valid
type AzurermDevspaceControllerInvalidNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermDevspaceControllerInvalidNameRule returns new rule with default attributes
func NewAzurermDevspaceControllerInvalidNameRule() *AzurermDevspaceControllerInvalidNameRule {
	return &AzurermDevspaceControllerInvalidNameRule{
		resourceType:  "azurerm_devspace_controller",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9]([_-]*[a-zA-Z0-9])*$`),
	}
}

// Name returns the rule name
func (r *AzurermDevspaceControllerInvalidNameRule) Name() string {
	return "azurerm_devspace_controller_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDevspaceControllerInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDevspaceControllerInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDevspaceControllerInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermDevspaceControllerInvalidNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9]([_-]*[a-zA-Z0-9])*$`),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
