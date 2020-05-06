// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermRecoveryServicesVaultInvalidSkuRule checks the pattern is valid
type AzurermRecoveryServicesVaultInvalidSkuRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermRecoveryServicesVaultInvalidSkuRule returns new rule with default attributes
func NewAzurermRecoveryServicesVaultInvalidSkuRule() *AzurermRecoveryServicesVaultInvalidSkuRule {
	return &AzurermRecoveryServicesVaultInvalidSkuRule{
		resourceType:  "azurerm_recovery_services_vault",
		attributeName: "sku",
		enum: []string{
			"Standard",
			"RS0",
		},
	}
}

// Name returns the rule name
func (r *AzurermRecoveryServicesVaultInvalidSkuRule) Name() string {
	return "azurerm_recovery_services_vault_invalid_sku"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermRecoveryServicesVaultInvalidSkuRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermRecoveryServicesVaultInvalidSkuRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermRecoveryServicesVaultInvalidSkuRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermRecoveryServicesVaultInvalidSkuRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as sku`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
