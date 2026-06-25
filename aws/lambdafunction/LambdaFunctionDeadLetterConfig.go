// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionDeadLetterConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/lambda_function#target_arn LambdaFunction#target_arn}.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
}

