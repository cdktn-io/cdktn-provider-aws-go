// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchapplication


type OpensearchApplicationDataSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/opensearch_application#data_source_arn OpensearchApplication#data_source_arn}.
	DataSourceArn *string `field:"optional" json:"dataSourceArn" yaml:"dataSourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/opensearch_application#data_source_description OpensearchApplication#data_source_description}.
	DataSourceDescription *string `field:"optional" json:"dataSourceDescription" yaml:"dataSourceDescription"`
}

