// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrulefororganization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/observabilityadmintelemetryrulefororganization/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference interface {
	cdktn.ComplexObject
	CloudtrailParameters() ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationCloudtrailParametersList
	CloudtrailParametersInput() interface{}
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
	DestinationPattern() *string
	SetDestinationPattern(val *string)
	DestinationPatternInput() *string
	DestinationType() *string
	SetDestinationType(val *string)
	DestinationTypeInput() *string
	ElbLoadBalancerLoggingParameters() ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationElbLoadBalancerLoggingParametersList
	ElbLoadBalancerLoggingParametersInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogDeliveryParameters() ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationLogDeliveryParametersList
	LogDeliveryParametersInput() interface{}
	MskMonitoringParameters() ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList
	MskMonitoringParametersInput() interface{}
	RetentionInDays() *float64
	SetRetentionInDays(val *float64)
	RetentionInDaysInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VpcFlowLogParameters() ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationVpcFlowLogParametersList
	VpcFlowLogParametersInput() interface{}
	WafLoggingParameters() ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationWafLoggingParametersList
	WafLoggingParametersInput() interface{}
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
	PutCloudtrailParameters(value interface{})
	PutElbLoadBalancerLoggingParameters(value interface{})
	PutLogDeliveryParameters(value interface{})
	PutMskMonitoringParameters(value interface{})
	PutVpcFlowLogParameters(value interface{})
	PutWafLoggingParameters(value interface{})
	ResetCloudtrailParameters()
	ResetDestinationPattern()
	ResetDestinationType()
	ResetElbLoadBalancerLoggingParameters()
	ResetLogDeliveryParameters()
	ResetMskMonitoringParameters()
	ResetRetentionInDays()
	ResetVpcFlowLogParameters()
	ResetWafLoggingParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference
type jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) CloudtrailParameters() ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationCloudtrailParametersList {
	var returns ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationCloudtrailParametersList
	_jsii_.Get(
		j,
		"cloudtrailParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) CloudtrailParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudtrailParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) DestinationPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) DestinationPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) DestinationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) DestinationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) ElbLoadBalancerLoggingParameters() ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationElbLoadBalancerLoggingParametersList {
	var returns ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationElbLoadBalancerLoggingParametersList
	_jsii_.Get(
		j,
		"elbLoadBalancerLoggingParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) ElbLoadBalancerLoggingParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elbLoadBalancerLoggingParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) LogDeliveryParameters() ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationLogDeliveryParametersList {
	var returns ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationLogDeliveryParametersList
	_jsii_.Get(
		j,
		"logDeliveryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) LogDeliveryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDeliveryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) MskMonitoringParameters() ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList {
	var returns ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationMskMonitoringParametersList
	_jsii_.Get(
		j,
		"mskMonitoringParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) MskMonitoringParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mskMonitoringParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) RetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) RetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) VpcFlowLogParameters() ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationVpcFlowLogParametersList {
	var returns ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationVpcFlowLogParametersList
	_jsii_.Get(
		j,
		"vpcFlowLogParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) VpcFlowLogParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcFlowLogParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) WafLoggingParameters() ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationWafLoggingParametersList {
	var returns ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationWafLoggingParametersList
	_jsii_.Get(
		j,
		"wafLoggingParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) WafLoggingParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wafLoggingParametersInput",
		&returns,
	)
	return returns
}


func NewObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.observabilityadminTelemetryRuleForOrganization.ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference_Override(o ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.observabilityadminTelemetryRuleForOrganization.ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference)SetDestinationPattern(val *string) {
	if err := j.validateSetDestinationPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPattern",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference)SetDestinationType(val *string) {
	if err := j.validateSetDestinationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationType",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference)SetRetentionInDays(val *float64) {
	if err := j.validateSetRetentionInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionInDays",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) PutCloudtrailParameters(value interface{}) {
	if err := o.validatePutCloudtrailParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCloudtrailParameters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) PutElbLoadBalancerLoggingParameters(value interface{}) {
	if err := o.validatePutElbLoadBalancerLoggingParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putElbLoadBalancerLoggingParameters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) PutLogDeliveryParameters(value interface{}) {
	if err := o.validatePutLogDeliveryParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLogDeliveryParameters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) PutMskMonitoringParameters(value interface{}) {
	if err := o.validatePutMskMonitoringParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMskMonitoringParameters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) PutVpcFlowLogParameters(value interface{}) {
	if err := o.validatePutVpcFlowLogParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putVpcFlowLogParameters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) PutWafLoggingParameters(value interface{}) {
	if err := o.validatePutWafLoggingParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWafLoggingParameters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) ResetCloudtrailParameters() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudtrailParameters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) ResetDestinationPattern() {
	_jsii_.InvokeVoid(
		o,
		"resetDestinationPattern",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) ResetDestinationType() {
	_jsii_.InvokeVoid(
		o,
		"resetDestinationType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) ResetElbLoadBalancerLoggingParameters() {
	_jsii_.InvokeVoid(
		o,
		"resetElbLoadBalancerLoggingParameters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) ResetLogDeliveryParameters() {
	_jsii_.InvokeVoid(
		o,
		"resetLogDeliveryParameters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) ResetMskMonitoringParameters() {
	_jsii_.InvokeVoid(
		o,
		"resetMskMonitoringParameters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) ResetRetentionInDays() {
	_jsii_.InvokeVoid(
		o,
		"resetRetentionInDays",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) ResetVpcFlowLogParameters() {
	_jsii_.InvokeVoid(
		o,
		"resetVpcFlowLogParameters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) ResetWafLoggingParameters() {
	_jsii_.InvokeVoid(
		o,
		"resetWafLoggingParameters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

