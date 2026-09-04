// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxwindowsfilesystem


type FsxWindowsFileSystemTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/fsx_windows_file_system#create FsxWindowsFileSystem#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/fsx_windows_file_system#delete FsxWindowsFileSystem#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/fsx_windows_file_system#update FsxWindowsFileSystem#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

