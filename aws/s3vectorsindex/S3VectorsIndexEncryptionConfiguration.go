// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3vectorsindex


type S3VectorsIndexEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/s3vectors_index#kms_key_arn S3VectorsIndex#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/s3vectors_index#sse_type S3VectorsIndex#sse_type}.
	SseType *string `field:"optional" json:"sseType" yaml:"sseType"`
}

