// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftidcapplication


type RedshiftIdcApplicationServiceIntegrationS3AccessGrants struct {
	// read_write_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/redshift_idc_application#read_write_access RedshiftIdcApplication#read_write_access}
	ReadWriteAccess interface{} `field:"optional" json:"readWriteAccess" yaml:"readWriteAccess"`
}

