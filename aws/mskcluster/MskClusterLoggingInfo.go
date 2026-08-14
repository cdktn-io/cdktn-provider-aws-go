// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterLoggingInfo struct {
	// broker_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/msk_cluster#broker_logs MskCluster#broker_logs}
	BrokerLogs *MskClusterLoggingInfoBrokerLogs `field:"required" json:"brokerLogs" yaml:"brokerLogs"`
}

