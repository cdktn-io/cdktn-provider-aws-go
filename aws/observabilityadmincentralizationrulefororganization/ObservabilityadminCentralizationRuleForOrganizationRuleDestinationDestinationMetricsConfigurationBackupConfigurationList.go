// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmincentralizationrulefororganization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/observabilityadmincentralizationrulefororganization/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList interface {
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
	Get(index *float64) ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList
type jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList {
	_init_.Initialize()

	if err := validateNewObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList{}

	_jsii_.Create(
		"@cdktn/provider-aws.observabilityadminCentralizationRuleForOrganization.ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList_Override(o ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.observabilityadminCentralizationRuleForOrganization.ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList) Get(index *float64) ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationOutputReference {
	if err := o.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationOutputReference

	_jsii_.Invoke(
		o,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationMetricsConfigurationBackupConfigurationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

