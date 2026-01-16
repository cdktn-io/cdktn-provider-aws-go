// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/athenaworkgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference interface {
	cdktf.ComplexObject
	CloudWatchLoggingConfiguration() AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference
	CloudWatchLoggingConfigurationInput() *AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfiguration
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
	// Experimental.
	Fqn() *string
	InternalValue() *AthenaWorkgroupConfigurationMonitoringConfiguration
	SetInternalValue(val *AthenaWorkgroupConfigurationMonitoringConfiguration)
	ManagedLoggingConfiguration() AthenaWorkgroupConfigurationMonitoringConfigurationManagedLoggingConfigurationOutputReference
	ManagedLoggingConfigurationInput() *AthenaWorkgroupConfigurationMonitoringConfigurationManagedLoggingConfiguration
	S3LoggingConfiguration() AthenaWorkgroupConfigurationMonitoringConfigurationS3LoggingConfigurationOutputReference
	S3LoggingConfigurationInput() *AthenaWorkgroupConfigurationMonitoringConfigurationS3LoggingConfiguration
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
	PutCloudWatchLoggingConfiguration(value *AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfiguration)
	PutManagedLoggingConfiguration(value *AthenaWorkgroupConfigurationMonitoringConfigurationManagedLoggingConfiguration)
	PutS3LoggingConfiguration(value *AthenaWorkgroupConfigurationMonitoringConfigurationS3LoggingConfiguration)
	ResetCloudWatchLoggingConfiguration()
	ResetManagedLoggingConfiguration()
	ResetS3LoggingConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference
type jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) CloudWatchLoggingConfiguration() AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference {
	var returns AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudWatchLoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) CloudWatchLoggingConfigurationInput() *AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfiguration {
	var returns *AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfiguration
	_jsii_.Get(
		j,
		"cloudWatchLoggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) InternalValue() *AthenaWorkgroupConfigurationMonitoringConfiguration {
	var returns *AthenaWorkgroupConfigurationMonitoringConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) ManagedLoggingConfiguration() AthenaWorkgroupConfigurationMonitoringConfigurationManagedLoggingConfigurationOutputReference {
	var returns AthenaWorkgroupConfigurationMonitoringConfigurationManagedLoggingConfigurationOutputReference
	_jsii_.Get(
		j,
		"managedLoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) ManagedLoggingConfigurationInput() *AthenaWorkgroupConfigurationMonitoringConfigurationManagedLoggingConfiguration {
	var returns *AthenaWorkgroupConfigurationMonitoringConfigurationManagedLoggingConfiguration
	_jsii_.Get(
		j,
		"managedLoggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) S3LoggingConfiguration() AthenaWorkgroupConfigurationMonitoringConfigurationS3LoggingConfigurationOutputReference {
	var returns AthenaWorkgroupConfigurationMonitoringConfigurationS3LoggingConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3LoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) S3LoggingConfigurationInput() *AthenaWorkgroupConfigurationMonitoringConfigurationS3LoggingConfiguration {
	var returns *AthenaWorkgroupConfigurationMonitoringConfigurationS3LoggingConfiguration
	_jsii_.Get(
		j,
		"s3LoggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAthenaWorkgroupConfigurationMonitoringConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewAthenaWorkgroupConfigurationMonitoringConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.athenaWorkgroup.AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAthenaWorkgroupConfigurationMonitoringConfigurationOutputReference_Override(a AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.athenaWorkgroup.AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference)SetInternalValue(val *AthenaWorkgroupConfigurationMonitoringConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) PutCloudWatchLoggingConfiguration(value *AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfiguration) {
	if err := a.validatePutCloudWatchLoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudWatchLoggingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) PutManagedLoggingConfiguration(value *AthenaWorkgroupConfigurationMonitoringConfigurationManagedLoggingConfiguration) {
	if err := a.validatePutManagedLoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedLoggingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) PutS3LoggingConfiguration(value *AthenaWorkgroupConfigurationMonitoringConfigurationS3LoggingConfiguration) {
	if err := a.validatePutS3LoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3LoggingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) ResetCloudWatchLoggingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudWatchLoggingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) ResetManagedLoggingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedLoggingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) ResetS3LoggingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetS3LoggingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationMonitoringConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

