// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionTracingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.32.1/docs/resources/lambda_function#mode LambdaFunction#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

