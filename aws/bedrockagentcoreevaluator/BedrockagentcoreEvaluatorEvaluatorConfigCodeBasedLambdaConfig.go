// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreevaluator


type BedrockagentcoreEvaluatorEvaluatorConfigCodeBasedLambdaConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_evaluator#lambda_arn BedrockagentcoreEvaluator#lambda_arn}.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_evaluator#lambda_timeout_in_seconds BedrockagentcoreEvaluator#lambda_timeout_in_seconds}.
	LambdaTimeoutInSeconds *float64 `field:"optional" json:"lambdaTimeoutInSeconds" yaml:"lambdaTimeoutInSeconds"`
}

