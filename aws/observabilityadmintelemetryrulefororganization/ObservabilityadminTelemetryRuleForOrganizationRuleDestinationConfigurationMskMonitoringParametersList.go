// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrulefororganization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/observabilityadmintelemetryrulefororganization/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList
type jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList {
	_init_.Initialize()

	if err := validateNewObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList{}

	_jsii_.Create(
		"@cdktn/provider-aws.observabilityadminTelemetryRuleForOrganization.ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList_Override(o ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.observabilityadminTelemetryRuleForOrganization.ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := o.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		o,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList) Get(index *float64) ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersOutputReference {
	if err := o.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersOutputReference

	_jsii_.Invoke(
		o,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

