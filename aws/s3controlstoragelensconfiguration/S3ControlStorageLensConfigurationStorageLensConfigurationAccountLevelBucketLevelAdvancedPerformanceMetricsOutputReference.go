// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3controlstoragelensconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/s3controlstoragelensconfiguration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetrics
	SetInternalValue(val *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetrics)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference
type jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) InternalValue() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetrics {
	var returns *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetrics
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference {
	_init_.Initialize()

	if err := validateNewS3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.s3ControlStorageLensConfiguration.S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference_Override(s S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.s3ControlStorageLensConfiguration.S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference)SetInternalValue(val *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetrics) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

