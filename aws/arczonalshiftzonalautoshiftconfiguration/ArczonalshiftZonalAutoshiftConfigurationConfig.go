// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arczonalshiftzonalautoshiftconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ArczonalshiftZonalAutoshiftConfigurationConfig struct {
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
	// The ARN of the managed resource to configure zonal autoshift for (e.g., an Application Load Balancer). Changing this creates a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/arczonalshift_zonal_autoshift_configuration#resource_arn ArczonalshiftZonalAutoshiftConfiguration#resource_arn}
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The status of zonal autoshift. Valid values: `ENABLED`, `DISABLED`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/arczonalshift_zonal_autoshift_configuration#zonal_autoshift_status ArczonalshiftZonalAutoshiftConfiguration#zonal_autoshift_status}
	ZonalAutoshiftStatus *string `field:"required" json:"zonalAutoshiftStatus" yaml:"zonalAutoshiftStatus"`
	// List of time windows during which practice runs are allowed, in the format `Day:HH:MM-Day:HH:MM` (e.g., `Mon:09:00-Mon:17:00`). Cannot be used together with `blocked_windows`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/arczonalshift_zonal_autoshift_configuration#allowed_windows ArczonalshiftZonalAutoshiftConfiguration#allowed_windows}
	AllowedWindows *[]*string `field:"optional" json:"allowedWindows" yaml:"allowedWindows"`
	// List of dates when practice runs should not be started, in the format `YYYY-MM-DD`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/arczonalshift_zonal_autoshift_configuration#blocked_dates ArczonalshiftZonalAutoshiftConfiguration#blocked_dates}
	BlockedDates *[]*string `field:"optional" json:"blockedDates" yaml:"blockedDates"`
	// List of time windows during which practice runs should not be started, in the format `Day:HH:MM-Day:HH:MM` (e.g., `Mon:00:00-Mon:08:00`). Cannot be used together with `allowed_windows`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/arczonalshift_zonal_autoshift_configuration#blocked_windows ArczonalshiftZonalAutoshiftConfiguration#blocked_windows}
	BlockedWindows *[]*string `field:"optional" json:"blockedWindows" yaml:"blockedWindows"`
	// blocking_alarms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/arczonalshift_zonal_autoshift_configuration#blocking_alarms ArczonalshiftZonalAutoshiftConfiguration#blocking_alarms}
	BlockingAlarms interface{} `field:"optional" json:"blockingAlarms" yaml:"blockingAlarms"`
	// outcome_alarms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/arczonalshift_zonal_autoshift_configuration#outcome_alarms ArczonalshiftZonalAutoshiftConfiguration#outcome_alarms}
	OutcomeAlarms interface{} `field:"optional" json:"outcomeAlarms" yaml:"outcomeAlarms"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/arczonalshift_zonal_autoshift_configuration#region ArczonalshiftZonalAutoshiftConfiguration#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

