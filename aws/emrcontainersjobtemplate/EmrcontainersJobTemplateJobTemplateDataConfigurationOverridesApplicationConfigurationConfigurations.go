// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emrcontainersjobtemplate


type EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesApplicationConfigurationConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/emrcontainers_job_template#classification EmrcontainersJobTemplate#classification}.
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/emrcontainers_job_template#properties EmrcontainersJobTemplate#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

