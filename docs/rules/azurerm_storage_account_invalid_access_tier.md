<!--- This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT --->

# azurerm_storage_account_invalid_access_tier

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

Allowed values are:
- Hot
- Cool

## Example

```hcl
resource "azurerm_storage_account" "foo" {
  access_tier = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: "..." is an invalid value as access_tier (azurerm_storage_account_invalid_access_tier)

  on template.tf line 2:
  2:   access_tier = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.1.0/docs/rules/azurerm_storage_account_invalid_access_tier.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

This rule is automatically generated from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). If you are uncertain about the warning, check the following API schema referenced by this rule.

https://github.com/Azure/azure-rest-api-specs/tree/master/specification/storage/resource-manager/Microsoft.Storage/stable/2019-06-01/storage.json