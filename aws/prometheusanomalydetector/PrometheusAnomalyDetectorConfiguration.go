// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package prometheusanomalydetector


type PrometheusAnomalyDetectorConfiguration struct {
	// random_cut_forest block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/prometheus_anomaly_detector#random_cut_forest PrometheusAnomalyDetector#random_cut_forest}
	RandomCutForest interface{} `field:"optional" json:"randomCutForest" yaml:"randomCutForest"`
}

