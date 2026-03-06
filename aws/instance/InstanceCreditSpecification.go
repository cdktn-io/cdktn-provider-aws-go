// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package instance


type InstanceCreditSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/instance#cpu_credits Instance#cpu_credits}.
	CpuCredits *string `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
}

