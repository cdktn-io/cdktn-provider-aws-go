// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionCapacityProviderConfig struct {
	// lambda_managed_instances_capacity_provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/lambda_function#lambda_managed_instances_capacity_provider_config LambdaFunction#lambda_managed_instances_capacity_provider_config}
	LambdaManagedInstancesCapacityProviderConfig *LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfig `field:"required" json:"lambdaManagedInstancesCapacityProviderConfig" yaml:"lambdaManagedInstancesCapacityProviderConfig"`
}

