// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobVpcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_training_job#security_group_ids SagemakerTrainingJob#security_group_ids}.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_training_job#subnets SagemakerTrainingJob#subnets}.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

