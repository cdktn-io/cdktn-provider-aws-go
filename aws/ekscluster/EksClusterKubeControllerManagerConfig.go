// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterKubeControllerManagerConfig struct {
	// horizontal_pod_autoscaler_controller_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/eks_cluster#horizontal_pod_autoscaler_controller_config EksCluster#horizontal_pod_autoscaler_controller_config}
	HorizontalPodAutoscalerControllerConfig *EksClusterKubeControllerManagerConfigHorizontalPodAutoscalerControllerConfig `field:"optional" json:"horizontalPodAutoscalerControllerConfig" yaml:"horizontalPodAutoscalerControllerConfig"`
	// pod_gc_controller_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/eks_cluster#pod_gc_controller_config EksCluster#pod_gc_controller_config}
	PodGcControllerConfig *EksClusterKubeControllerManagerConfigPodGcControllerConfig `field:"optional" json:"podGcControllerConfig" yaml:"podGcControllerConfig"`
}

