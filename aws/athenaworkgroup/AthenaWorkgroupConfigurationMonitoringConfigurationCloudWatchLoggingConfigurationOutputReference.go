// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/athenaworkgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfiguration
	SetInternalValue(val *AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfiguration)
	LogGroup() *string
	SetLogGroup(val *string)
	LogGroupInput() *string
	LogStreamNamePrefix() *string
	SetLogStreamNamePrefix(val *string)
	LogStreamNamePrefixInput() *string
	LogType() AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationLogTypeList
	LogTypeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutLogType(value interface{})
	ResetLogGroup()
	ResetLogStreamNamePrefix()
	ResetLogType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference
type jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) InternalValue() *AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfiguration {
	var returns *AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) LogGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) LogGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) LogStreamNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) LogStreamNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) LogType() AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationLogTypeList {
	var returns AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationLogTypeList
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) LogTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewAthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.athenaWorkgroup.AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference_Override(a AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.athenaWorkgroup.AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference)SetInternalValue(val *AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference)SetLogGroup(val *string) {
	if err := j.validateSetLogGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroup",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference)SetLogStreamNamePrefix(val *string) {
	if err := j.validateSetLogStreamNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStreamNamePrefix",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) PutLogType(value interface{}) {
	if err := a.validatePutLogTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) ResetLogGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetLogGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) ResetLogStreamNamePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetLogStreamNamePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) ResetLogType() {
	_jsii_.InvokeVoid(
		a,
		"resetLogType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

