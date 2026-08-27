// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package prometheusanomalydetector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PrometheusAnomalyDetectorConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#alias PrometheusAnomalyDetector#alias}.
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#workspace_id PrometheusAnomalyDetector#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#configuration PrometheusAnomalyDetector#configuration}
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#evaluation_interval_in_seconds PrometheusAnomalyDetector#evaluation_interval_in_seconds}.
	EvaluationIntervalInSeconds *float64 `field:"optional" json:"evaluationIntervalInSeconds" yaml:"evaluationIntervalInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#labels PrometheusAnomalyDetector#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// missing_data_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#missing_data_action PrometheusAnomalyDetector#missing_data_action}
	MissingDataAction interface{} `field:"optional" json:"missingDataAction" yaml:"missingDataAction"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#region PrometheusAnomalyDetector#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#tags PrometheusAnomalyDetector#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_anomaly_detector#timeouts PrometheusAnomalyDetector#timeouts}
	Timeouts *PrometheusAnomalyDetectorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

