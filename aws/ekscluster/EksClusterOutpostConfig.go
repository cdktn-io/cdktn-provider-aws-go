// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterOutpostConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/eks_cluster#control_plane_instance_type EksCluster#control_plane_instance_type}.
	ControlPlaneInstanceType *string `field:"required" json:"controlPlaneInstanceType" yaml:"controlPlaneInstanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/eks_cluster#outpost_arns EksCluster#outpost_arns}.
	OutpostArns *[]*string `field:"required" json:"outpostArns" yaml:"outpostArns"`
	// control_plane_placement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/eks_cluster#control_plane_placement EksCluster#control_plane_placement}
	ControlPlanePlacement *EksClusterOutpostConfigControlPlanePlacement `field:"optional" json:"controlPlanePlacement" yaml:"controlPlanePlacement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/eks_cluster#etcd_instance_type EksCluster#etcd_instance_type}.
	EtcdInstanceType *string `field:"optional" json:"etcdInstanceType" yaml:"etcdInstanceType"`
	// etcd_placement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/eks_cluster#etcd_placement EksCluster#etcd_placement}
	EtcdPlacement *EksClusterOutpostConfigEtcdPlacement `field:"optional" json:"etcdPlacement" yaml:"etcdPlacement"`
}

