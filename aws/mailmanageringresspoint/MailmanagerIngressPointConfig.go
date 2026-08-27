// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanageringresspoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MailmanagerIngressPointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#name MailmanagerIngressPoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#rule_set_id MailmanagerIngressPoint#rule_set_id}.
	RuleSetId *string `field:"required" json:"ruleSetId" yaml:"ruleSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#traffic_policy_id MailmanagerIngressPoint#traffic_policy_id}.
	TrafficPolicyId *string `field:"required" json:"trafficPolicyId" yaml:"trafficPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#type MailmanagerIngressPoint#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// ingress_point_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#ingress_point_configuration MailmanagerIngressPoint#ingress_point_configuration}
	IngressPointConfiguration interface{} `field:"optional" json:"ingressPointConfiguration" yaml:"ingressPointConfiguration"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#network_configuration MailmanagerIngressPoint#network_configuration}
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#region MailmanagerIngressPoint#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#tags MailmanagerIngressPoint#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#timeouts MailmanagerIngressPoint#timeouts}
	Timeouts *MailmanagerIngressPointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#tls_policy MailmanagerIngressPoint#tls_policy}.
	TlsPolicy *string `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
}

