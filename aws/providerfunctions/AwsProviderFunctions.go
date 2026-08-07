// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package providerfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Provider-defined functions of the aws provider.
type AwsProviderFunctions interface {
	// Builds an ARN from its constituent parts.
	ArnBuild(partition *string, service *string, region *string, accountId *string, resource *string) *string
	// Parses an ARN into its constituent parts.
	ArnParse(arn *string) cdktn.IResolvable
	// Trims the path prefix from an IAM role Amazon Resource Name (ARN).
	//
	// This function can be used when services require role ARNs to be passed without a path.
	TrimIamRolePath(arn *string) *string
	// Formats a User-Agent product for use with the user_agent argument in the provider or provider_meta block.
	UserAgent(productName *string, productVersion *string, comment *string) *string
}

// The jsii proxy struct for AwsProviderFunctions
type jsiiProxy_AwsProviderFunctions struct {
	_ byte // padding
}

func NewAwsProviderFunctions(providerLocalName *string) AwsProviderFunctions {
	_init_.Initialize()

	if err := validateNewAwsProviderFunctionsParameters(providerLocalName); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsProviderFunctions{}

	_jsii_.Create(
		"@cdktn/provider-aws.providerFunctions.AwsProviderFunctions",
		[]interface{}{providerLocalName},
		&j,
	)

	return &j
}

func NewAwsProviderFunctions_Override(a AwsProviderFunctions, providerLocalName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.providerFunctions.AwsProviderFunctions",
		[]interface{}{providerLocalName},
		a,
	)
}

func (a *jsiiProxy_AwsProviderFunctions) ArnBuild(partition *string, service *string, region *string, accountId *string, resource *string) *string {
	if err := a.validateArnBuildParameters(partition, service, region, accountId, resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"arnBuild",
		[]interface{}{partition, service, region, accountId, resource},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProviderFunctions) ArnParse(arn *string) cdktn.IResolvable {
	if err := a.validateArnParseParameters(arn); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"arnParse",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProviderFunctions) TrimIamRolePath(arn *string) *string {
	if err := a.validateTrimIamRolePathParameters(arn); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"trimIamRolePath",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProviderFunctions) UserAgent(productName *string, productVersion *string, comment *string) *string {
	if err := a.validateUserAgentParameters(productName, productVersion, comment); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"userAgent",
		[]interface{}{productName, productVersion, comment},
		&returns,
	)

	return returns
}

