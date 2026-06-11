// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionEnvironment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/lambda_function#variables LambdaFunction#variables}.
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
}

