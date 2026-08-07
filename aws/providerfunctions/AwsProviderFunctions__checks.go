// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package providerfunctions

import (
	"fmt"
)

func (a *jsiiProxy_AwsProviderFunctions) validateArnBuildParameters(partition *string, service *string, region *string, accountId *string, resource *string) error {
	if partition == nil {
		return fmt.Errorf("parameter partition is required, but nil was provided")
	}

	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	if region == nil {
		return fmt.Errorf("parameter region is required, but nil was provided")
	}

	if accountId == nil {
		return fmt.Errorf("parameter accountId is required, but nil was provided")
	}

	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsProviderFunctions) validateArnParseParameters(arn *string) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsProviderFunctions) validateTrimIamRolePathParameters(arn *string) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsProviderFunctions) validateUserAgentParameters(productName *string, productVersion *string, comment *string) error {
	if productName == nil {
		return fmt.Errorf("parameter productName is required, but nil was provided")
	}

	if productVersion == nil {
		return fmt.Errorf("parameter productVersion is required, but nil was provided")
	}

	if comment == nil {
		return fmt.Errorf("parameter comment is required, but nil was provided")
	}

	return nil
}

func validateNewAwsProviderFunctionsParameters(providerLocalName *string) error {
	if providerLocalName == nil {
		return fmt.Errorf("parameter providerLocalName is required, but nil was provided")
	}

	return nil
}

