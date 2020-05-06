// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermWindowsVirtualMachineInvalidEvictionPolicyRule checks the pattern is valid
type AzurermWindowsVirtualMachineInvalidEvictionPolicyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermWindowsVirtualMachineInvalidEvictionPolicyRule returns new rule with default attributes
func NewAzurermWindowsVirtualMachineInvalidEvictionPolicyRule() *AzurermWindowsVirtualMachineInvalidEvictionPolicyRule {
	return &AzurermWindowsVirtualMachineInvalidEvictionPolicyRule{
		resourceType:  "azurerm_windows_virtual_machine",
		attributeName: "eviction_policy",
		enum: []string{
			"Deallocate",
			"Delete",
		},
	}
}

// Name returns the rule name
func (r *AzurermWindowsVirtualMachineInvalidEvictionPolicyRule) Name() string {
	return "azurerm_windows_virtual_machine_invalid_eviction_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermWindowsVirtualMachineInvalidEvictionPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermWindowsVirtualMachineInvalidEvictionPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermWindowsVirtualMachineInvalidEvictionPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermWindowsVirtualMachineInvalidEvictionPolicyRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as eviction_policy`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
