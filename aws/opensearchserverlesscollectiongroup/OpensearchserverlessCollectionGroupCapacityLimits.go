// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchserverlesscollectiongroup


type OpensearchserverlessCollectionGroupCapacityLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/opensearchserverless_collection_group#max_indexing_capacity_in_ocu OpensearchserverlessCollectionGroup#max_indexing_capacity_in_ocu}.
	MaxIndexingCapacityInOcu *float64 `field:"optional" json:"maxIndexingCapacityInOcu" yaml:"maxIndexingCapacityInOcu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/opensearchserverless_collection_group#max_search_capacity_in_ocu OpensearchserverlessCollectionGroup#max_search_capacity_in_ocu}.
	MaxSearchCapacityInOcu *float64 `field:"optional" json:"maxSearchCapacityInOcu" yaml:"maxSearchCapacityInOcu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/opensearchserverless_collection_group#min_indexing_capacity_in_ocu OpensearchserverlessCollectionGroup#min_indexing_capacity_in_ocu}.
	MinIndexingCapacityInOcu *float64 `field:"optional" json:"minIndexingCapacityInOcu" yaml:"minIndexingCapacityInOcu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/opensearchserverless_collection_group#min_search_capacity_in_ocu OpensearchserverlessCollectionGroup#min_search_capacity_in_ocu}.
	MinSearchCapacityInOcu *float64 `field:"optional" json:"minSearchCapacityInOcu" yaml:"minSearchCapacityInOcu"`
}

