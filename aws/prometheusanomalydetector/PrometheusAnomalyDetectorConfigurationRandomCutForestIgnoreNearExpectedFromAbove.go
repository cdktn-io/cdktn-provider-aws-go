// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package prometheusanomalydetector


type PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAbove struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#amount PrometheusAnomalyDetector#amount}.
	Amount *float64 `field:"optional" json:"amount" yaml:"amount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#ratio PrometheusAnomalyDetector#ratio}.
	Ratio *float64 `field:"optional" json:"ratio" yaml:"ratio"`
}

