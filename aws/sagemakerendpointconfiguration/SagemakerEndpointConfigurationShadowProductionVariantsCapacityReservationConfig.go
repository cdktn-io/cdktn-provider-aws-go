// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfiguration


type SagemakerEndpointConfigurationShadowProductionVariantsCapacityReservationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/sagemaker_endpoint_configuration#capacity_reservation_preference SagemakerEndpointConfiguration#capacity_reservation_preference}.
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/sagemaker_endpoint_configuration#ml_reservation_arn SagemakerEndpointConfiguration#ml_reservation_arn}.
	MlReservationArn *string `field:"optional" json:"mlReservationArn" yaml:"mlReservationArn"`
}

