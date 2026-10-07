// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase


type OdbAutonomousDatabaseAdminPasswordSourceCustomerManagedAwsSecret struct {
	// Type of OCI identifier supplied as the external ID when OCI assumes the IAM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#external_id_type OdbAutonomousDatabase#external_id_type}
	ExternalIdType *string `field:"required" json:"externalIdType" yaml:"externalIdType"`
	// ARN of the customer-managed IAM role OCI assumes to retrieve the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#iam_role_arn OdbAutonomousDatabase#iam_role_arn}
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// ARN of the AWS Secrets Manager secret that contains the ADMIN password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#secret_arn OdbAutonomousDatabase#secret_arn}
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
}

