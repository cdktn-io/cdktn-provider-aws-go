// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreamwritetable


type TimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocation struct {
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/timestreamwrite_table#s3_configuration TimestreamwriteTable#s3_configuration}
	S3Configuration *TimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocationS3Configuration `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

