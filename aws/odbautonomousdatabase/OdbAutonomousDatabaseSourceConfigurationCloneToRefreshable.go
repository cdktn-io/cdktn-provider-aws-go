// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase


type OdbAutonomousDatabaseSourceConfigurationCloneToRefreshable struct {
	// ID of the source Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#source_autonomous_database_id OdbAutonomousDatabase#source_autonomous_database_id}
	SourceAutonomousDatabaseId *string `field:"required" json:"sourceAutonomousDatabaseId" yaml:"sourceAutonomousDatabaseId"`
	// Frequency at which the refreshable clone is automatically refreshed, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#auto_refresh_frequency_in_seconds OdbAutonomousDatabase#auto_refresh_frequency_in_seconds}
	AutoRefreshFrequencyInSeconds *float64 `field:"optional" json:"autoRefreshFrequencyInSeconds" yaml:"autoRefreshFrequencyInSeconds"`
	// Time lag between the refreshable clone and its source, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#auto_refresh_point_lag_in_seconds OdbAutonomousDatabase#auto_refresh_point_lag_in_seconds}
	AutoRefreshPointLagInSeconds *float64 `field:"optional" json:"autoRefreshPointLagInSeconds" yaml:"autoRefreshPointLagInSeconds"`
	// Type of clone to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#clone_type OdbAutonomousDatabase#clone_type}
	CloneType *string `field:"optional" json:"cloneType" yaml:"cloneType"`
	// Open mode of the refreshable clone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#open_mode OdbAutonomousDatabase#open_mode}
	OpenMode *string `field:"optional" json:"openMode" yaml:"openMode"`
	// Refresh mode of the clone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#refreshable_mode OdbAutonomousDatabase#refreshable_mode}
	RefreshableMode *string `field:"optional" json:"refreshableMode" yaml:"refreshableMode"`
	// Date and time when automatic refresh starts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#time_of_auto_refresh_start OdbAutonomousDatabase#time_of_auto_refresh_start}
	TimeOfAutoRefreshStart *string `field:"optional" json:"timeOfAutoRefreshStart" yaml:"timeOfAutoRefreshStart"`
}

