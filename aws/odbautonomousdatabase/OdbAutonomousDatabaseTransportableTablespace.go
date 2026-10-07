// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase


type OdbAutonomousDatabaseTransportableTablespace struct {
	// URL of the transportable tablespace bundle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#tts_bundle_url OdbAutonomousDatabase#tts_bundle_url}
	TtsBundleUrl *string `field:"optional" json:"ttsBundleUrl" yaml:"ttsBundleUrl"`
}

