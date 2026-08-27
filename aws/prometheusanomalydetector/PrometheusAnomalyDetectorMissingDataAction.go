// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package prometheusanomalydetector


type PrometheusAnomalyDetectorMissingDataAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#mark_as_anomaly PrometheusAnomalyDetector#mark_as_anomaly}.
	MarkAsAnomaly interface{} `field:"optional" json:"markAsAnomaly" yaml:"markAsAnomaly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#skip PrometheusAnomalyDetector#skip}.
	Skip interface{} `field:"optional" json:"skip" yaml:"skip"`
}

