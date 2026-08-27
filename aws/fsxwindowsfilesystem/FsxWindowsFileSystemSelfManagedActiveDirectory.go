// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxwindowsfilesystem


type FsxWindowsFileSystemSelfManagedActiveDirectory struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#dns_ips FsxWindowsFileSystem#dns_ips}.
	DnsIps *[]*string `field:"required" json:"dnsIps" yaml:"dnsIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#domain_name FsxWindowsFileSystem#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#domain_join_service_account_secret FsxWindowsFileSystem#domain_join_service_account_secret}.
	DomainJoinServiceAccountSecret *string `field:"optional" json:"domainJoinServiceAccountSecret" yaml:"domainJoinServiceAccountSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#file_system_administrators_group FsxWindowsFileSystem#file_system_administrators_group}.
	FileSystemAdministratorsGroup *string `field:"optional" json:"fileSystemAdministratorsGroup" yaml:"fileSystemAdministratorsGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#organizational_unit_distinguished_name FsxWindowsFileSystem#organizational_unit_distinguished_name}.
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#password FsxWindowsFileSystem#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#password_wo FsxWindowsFileSystem#password_wo}.
	PasswordWo *string `field:"optional" json:"passwordWo" yaml:"passwordWo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#password_wo_version FsxWindowsFileSystem#password_wo_version}.
	PasswordWoVersion *float64 `field:"optional" json:"passwordWoVersion" yaml:"passwordWoVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#username FsxWindowsFileSystem#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

