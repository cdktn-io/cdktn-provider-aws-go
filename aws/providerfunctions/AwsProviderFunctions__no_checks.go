// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package providerfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsProviderFunctions) validateArnBuildParameters(partition *string, service *string, region *string, accountId *string, resource *string) error {
	return nil
}

func (a *jsiiProxy_AwsProviderFunctions) validateArnParseParameters(arn *string) error {
	return nil
}

func (a *jsiiProxy_AwsProviderFunctions) validateTrimIamRolePathParameters(arn *string) error {
	return nil
}

func (a *jsiiProxy_AwsProviderFunctions) validateUserAgentParameters(productName *string, productVersion *string, comment *string) error {
	return nil
}

func validateNewAwsProviderFunctionsParameters(providerLocalName *string) error {
	return nil
}

