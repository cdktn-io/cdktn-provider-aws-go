// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2inputsource


type Resiliencehubv2InputSourceResourceConfigurationEks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/resiliencehubv2_input_source#cluster_arn Resiliencehubv2InputSource#cluster_arn}.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/resiliencehubv2_input_source#namespaces Resiliencehubv2InputSource#namespaces}.
	Namespaces *[]*string `field:"required" json:"namespaces" yaml:"namespaces"`
}

