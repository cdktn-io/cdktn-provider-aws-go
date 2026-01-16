// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftidcapplication


type RedshiftIdcApplicationServiceIntegration struct {
	// lake_formation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/redshift_idc_application#lake_formation RedshiftIdcApplication#lake_formation}
	LakeFormation interface{} `field:"optional" json:"lakeFormation" yaml:"lakeFormation"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/redshift_idc_application#redshift RedshiftIdcApplication#redshift}
	Redshift interface{} `field:"optional" json:"redshift" yaml:"redshift"`
	// s3_access_grants block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/redshift_idc_application#s3_access_grants RedshiftIdcApplication#s3_access_grants}
	S3AccessGrants interface{} `field:"optional" json:"s3AccessGrants" yaml:"s3AccessGrants"`
}

