// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package prometheusanomalydetector


type PrometheusAnomalyDetectorConfigurationRandomCutForest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/prometheus_anomaly_detector#query PrometheusAnomalyDetector#query}.
	Query *string `field:"required" json:"query" yaml:"query"`
	// ignore_near_expected_from_above block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/prometheus_anomaly_detector#ignore_near_expected_from_above PrometheusAnomalyDetector#ignore_near_expected_from_above}
	IgnoreNearExpectedFromAbove interface{} `field:"optional" json:"ignoreNearExpectedFromAbove" yaml:"ignoreNearExpectedFromAbove"`
	// ignore_near_expected_from_below block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/prometheus_anomaly_detector#ignore_near_expected_from_below PrometheusAnomalyDetector#ignore_near_expected_from_below}
	IgnoreNearExpectedFromBelow interface{} `field:"optional" json:"ignoreNearExpectedFromBelow" yaml:"ignoreNearExpectedFromBelow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/prometheus_anomaly_detector#sample_size PrometheusAnomalyDetector#sample_size}.
	SampleSize *float64 `field:"optional" json:"sampleSize" yaml:"sampleSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/prometheus_anomaly_detector#shingle_size PrometheusAnomalyDetector#shingle_size}.
	ShingleSize *float64 `field:"optional" json:"shingleSize" yaml:"shingleSize"`
}

