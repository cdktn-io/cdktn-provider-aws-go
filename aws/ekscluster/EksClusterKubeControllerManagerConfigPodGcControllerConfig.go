// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterKubeControllerManagerConfigPodGcControllerConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/eks_cluster#terminated_pod_gc_threshold EksCluster#terminated_pod_gc_threshold}.
	TerminatedPodGcThreshold *float64 `field:"optional" json:"terminatedPodGcThreshold" yaml:"terminatedPodGcThreshold"`
}

