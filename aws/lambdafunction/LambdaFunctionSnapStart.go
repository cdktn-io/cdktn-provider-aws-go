// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionSnapStart struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.37.0/docs/resources/lambda_function#apply_on LambdaFunction#apply_on}.
	ApplyOn *string `field:"required" json:"applyOn" yaml:"applyOn"`
}

