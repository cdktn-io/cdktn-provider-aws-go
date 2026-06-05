// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynctask


type DatasyncTaskSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/datasync_task#schedule_expression DatasyncTask#schedule_expression}.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/datasync_task#status DatasyncTask#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

