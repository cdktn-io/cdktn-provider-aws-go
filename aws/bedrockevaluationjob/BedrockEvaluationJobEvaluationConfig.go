// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobEvaluationConfig struct {
	// automated block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrock_evaluation_job#automated BedrockEvaluationJob#automated}
	Automated interface{} `field:"optional" json:"automated" yaml:"automated"`
	// human block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrock_evaluation_job#human BedrockEvaluationJob#human}
	Human interface{} `field:"optional" json:"human" yaml:"human"`
}

