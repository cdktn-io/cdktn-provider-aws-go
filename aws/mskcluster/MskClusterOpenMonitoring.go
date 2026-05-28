// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterOpenMonitoring struct {
	// prometheus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/msk_cluster#prometheus MskCluster#prometheus}
	Prometheus *MskClusterOpenMonitoringPrometheus `field:"required" json:"prometheus" yaml:"prometheus"`
}

