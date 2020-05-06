// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermHealthcareServiceInvalidCosmosdbThroughputRule checks the pattern is valid
type AzurermHealthcareServiceInvalidCosmosdbThroughputRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAzurermHealthcareServiceInvalidCosmosdbThroughputRule returns new rule with default attributes
func NewAzurermHealthcareServiceInvalidCosmosdbThroughputRule() *AzurermHealthcareServiceInvalidCosmosdbThroughputRule {
	return &AzurermHealthcareServiceInvalidCosmosdbThroughputRule{
		resourceType:  "azurerm_healthcare_service",
		attributeName: "cosmosdb_throughput",
		max:           10000,
		min:           400,
	}
}

// Name returns the rule name
func (r *AzurermHealthcareServiceInvalidCosmosdbThroughputRule) Name() string {
	return "azurerm_healthcare_service_invalid_cosmosdb_throughput"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermHealthcareServiceInvalidCosmosdbThroughputRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermHealthcareServiceInvalidCosmosdbThroughputRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermHealthcareServiceInvalidCosmosdbThroughputRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermHealthcareServiceInvalidCosmosdbThroughputRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val int
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if val > r.max {
				runner.EmitIssue(
					r,
					"cosmosdb_throughput must be 10000 or less",
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			if val < r.min {
				runner.EmitIssue(
					r,
					"cosmosdb_throughput must be 400 or higher",
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
