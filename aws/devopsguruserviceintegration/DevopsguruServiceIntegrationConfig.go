// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsguruserviceintegration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DevopsguruServiceIntegrationConfig struct {
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
	// kms_server_side_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/devopsguru_service_integration#kms_server_side_encryption DevopsguruServiceIntegration#kms_server_side_encryption}
	KmsServerSideEncryption interface{} `field:"optional" json:"kmsServerSideEncryption" yaml:"kmsServerSideEncryption"`
	// logs_anomaly_detection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/devopsguru_service_integration#logs_anomaly_detection DevopsguruServiceIntegration#logs_anomaly_detection}
	LogsAnomalyDetection interface{} `field:"optional" json:"logsAnomalyDetection" yaml:"logsAnomalyDetection"`
	// ops_center block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/devopsguru_service_integration#ops_center DevopsguruServiceIntegration#ops_center}
	OpsCenter interface{} `field:"optional" json:"opsCenter" yaml:"opsCenter"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/devopsguru_service_integration#region DevopsguruServiceIntegration#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

