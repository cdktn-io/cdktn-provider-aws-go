// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eksnodegroup


type EksNodeGroupWarmPoolConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/eks_node_group#max_group_prepared_capacity EksNodeGroup#max_group_prepared_capacity}.
	MaxGroupPreparedCapacity *float64 `field:"optional" json:"maxGroupPreparedCapacity" yaml:"maxGroupPreparedCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/eks_node_group#min_size EksNodeGroup#min_size}.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/eks_node_group#pool_state EksNodeGroup#pool_state}.
	PoolState *string `field:"optional" json:"poolState" yaml:"poolState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/eks_node_group#reuse_on_scale_in EksNodeGroup#reuse_on_scale_in}.
	ReuseOnScaleIn interface{} `field:"optional" json:"reuseOnScaleIn" yaml:"reuseOnScaleIn"`
}

