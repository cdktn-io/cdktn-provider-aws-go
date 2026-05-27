// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobDebugHookConfigCollectionConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_training_job#collection_name SagemakerTrainingJob#collection_name}.
	CollectionName *string `field:"optional" json:"collectionName" yaml:"collectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_training_job#collection_parameters SagemakerTrainingJob#collection_parameters}.
	CollectionParameters *map[string]*string `field:"optional" json:"collectionParameters" yaml:"collectionParameters"`
}

