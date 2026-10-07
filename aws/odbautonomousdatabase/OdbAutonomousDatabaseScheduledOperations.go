// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase


type OdbAutonomousDatabaseScheduledOperations struct {
	// Day of the week.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#day_of_week OdbAutonomousDatabase#day_of_week}
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Scheduled start time in UTC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#scheduled_start_time OdbAutonomousDatabase#scheduled_start_time}
	ScheduledStartTime *string `field:"optional" json:"scheduledStartTime" yaml:"scheduledStartTime"`
	// Scheduled stop time in UTC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#scheduled_stop_time OdbAutonomousDatabase#scheduled_stop_time}
	ScheduledStopTime *string `field:"optional" json:"scheduledStopTime" yaml:"scheduledStopTime"`
}

