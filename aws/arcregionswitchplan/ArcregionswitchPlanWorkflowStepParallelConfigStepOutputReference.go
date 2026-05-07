// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/arcregionswitchplan/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference interface {
	cdktn.ComplexObject
	ArcRoutingControlConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepArcRoutingControlConfigList
	ArcRoutingControlConfigInput() interface{}
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
	CustomActionLambdaConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigList
	CustomActionLambdaConfigInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DocumentDbConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepDocumentDbConfigList
	DocumentDbConfigInput() interface{}
	Ec2AsgCapacityIncreaseConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigList
	Ec2AsgCapacityIncreaseConfigInput() interface{}
	EcsCapacityIncreaseConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepEcsCapacityIncreaseConfigList
	EcsCapacityIncreaseConfigInput() interface{}
	EksResourceScalingConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigList
	EksResourceScalingConfigInput() interface{}
	ExecutionApprovalConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepExecutionApprovalConfigList
	ExecutionApprovalConfigInput() interface{}
	ExecutionBlockType() *string
	SetExecutionBlockType(val *string)
	ExecutionBlockTypeInput() *string
	// Experimental.
	Fqn() *string
	GlobalAuroraConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepGlobalAuroraConfigList
	GlobalAuroraConfigInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	RegionSwitchPlanConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepRegionSwitchPlanConfigList
	RegionSwitchPlanConfigInput() interface{}
	Route53HealthCheckConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepRoute53HealthCheckConfigList
	Route53HealthCheckConfigInput() interface{}
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
	PutArcRoutingControlConfig(value interface{})
	PutCustomActionLambdaConfig(value interface{})
	PutDocumentDbConfig(value interface{})
	PutEc2AsgCapacityIncreaseConfig(value interface{})
	PutEcsCapacityIncreaseConfig(value interface{})
	PutEksResourceScalingConfig(value interface{})
	PutExecutionApprovalConfig(value interface{})
	PutGlobalAuroraConfig(value interface{})
	PutRegionSwitchPlanConfig(value interface{})
	PutRoute53HealthCheckConfig(value interface{})
	ResetArcRoutingControlConfig()
	ResetCustomActionLambdaConfig()
	ResetDescription()
	ResetDocumentDbConfig()
	ResetEc2AsgCapacityIncreaseConfig()
	ResetEcsCapacityIncreaseConfig()
	ResetEksResourceScalingConfig()
	ResetExecutionApprovalConfig()
	ResetGlobalAuroraConfig()
	ResetRegionSwitchPlanConfig()
	ResetRoute53HealthCheckConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference
type jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ArcRoutingControlConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepArcRoutingControlConfigList {
	var returns ArcregionswitchPlanWorkflowStepParallelConfigStepArcRoutingControlConfigList
	_jsii_.Get(
		j,
		"arcRoutingControlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ArcRoutingControlConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"arcRoutingControlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) CustomActionLambdaConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigList {
	var returns ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigList
	_jsii_.Get(
		j,
		"customActionLambdaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) CustomActionLambdaConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customActionLambdaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) DocumentDbConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepDocumentDbConfigList {
	var returns ArcregionswitchPlanWorkflowStepParallelConfigStepDocumentDbConfigList
	_jsii_.Get(
		j,
		"documentDbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) DocumentDbConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentDbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) Ec2AsgCapacityIncreaseConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigList {
	var returns ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigList
	_jsii_.Get(
		j,
		"ec2AsgCapacityIncreaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) Ec2AsgCapacityIncreaseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2AsgCapacityIncreaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) EcsCapacityIncreaseConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepEcsCapacityIncreaseConfigList {
	var returns ArcregionswitchPlanWorkflowStepParallelConfigStepEcsCapacityIncreaseConfigList
	_jsii_.Get(
		j,
		"ecsCapacityIncreaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) EcsCapacityIncreaseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecsCapacityIncreaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) EksResourceScalingConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigList {
	var returns ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigList
	_jsii_.Get(
		j,
		"eksResourceScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) EksResourceScalingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eksResourceScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ExecutionApprovalConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepExecutionApprovalConfigList {
	var returns ArcregionswitchPlanWorkflowStepParallelConfigStepExecutionApprovalConfigList
	_jsii_.Get(
		j,
		"executionApprovalConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ExecutionApprovalConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executionApprovalConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ExecutionBlockType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionBlockType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ExecutionBlockTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionBlockTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) GlobalAuroraConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepGlobalAuroraConfigList {
	var returns ArcregionswitchPlanWorkflowStepParallelConfigStepGlobalAuroraConfigList
	_jsii_.Get(
		j,
		"globalAuroraConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) GlobalAuroraConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalAuroraConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) RegionSwitchPlanConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepRegionSwitchPlanConfigList {
	var returns ArcregionswitchPlanWorkflowStepParallelConfigStepRegionSwitchPlanConfigList
	_jsii_.Get(
		j,
		"regionSwitchPlanConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) RegionSwitchPlanConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionSwitchPlanConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) Route53HealthCheckConfig() ArcregionswitchPlanWorkflowStepParallelConfigStepRoute53HealthCheckConfigList {
	var returns ArcregionswitchPlanWorkflowStepParallelConfigStepRoute53HealthCheckConfigList
	_jsii_.Get(
		j,
		"route53HealthCheckConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) Route53HealthCheckConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"route53HealthCheckConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference {
	_init_.Initialize()

	if err := validateNewArcregionswitchPlanWorkflowStepParallelConfigStepOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.arcregionswitchPlan.ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference_Override(a ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.arcregionswitchPlan.ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference)SetExecutionBlockType(val *string) {
	if err := j.validateSetExecutionBlockTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionBlockType",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) PutArcRoutingControlConfig(value interface{}) {
	if err := a.validatePutArcRoutingControlConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putArcRoutingControlConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) PutCustomActionLambdaConfig(value interface{}) {
	if err := a.validatePutCustomActionLambdaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomActionLambdaConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) PutDocumentDbConfig(value interface{}) {
	if err := a.validatePutDocumentDbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDocumentDbConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) PutEc2AsgCapacityIncreaseConfig(value interface{}) {
	if err := a.validatePutEc2AsgCapacityIncreaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEc2AsgCapacityIncreaseConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) PutEcsCapacityIncreaseConfig(value interface{}) {
	if err := a.validatePutEcsCapacityIncreaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcsCapacityIncreaseConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) PutEksResourceScalingConfig(value interface{}) {
	if err := a.validatePutEksResourceScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEksResourceScalingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) PutExecutionApprovalConfig(value interface{}) {
	if err := a.validatePutExecutionApprovalConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExecutionApprovalConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) PutGlobalAuroraConfig(value interface{}) {
	if err := a.validatePutGlobalAuroraConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGlobalAuroraConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) PutRegionSwitchPlanConfig(value interface{}) {
	if err := a.validatePutRegionSwitchPlanConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRegionSwitchPlanConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) PutRoute53HealthCheckConfig(value interface{}) {
	if err := a.validatePutRoute53HealthCheckConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRoute53HealthCheckConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ResetArcRoutingControlConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetArcRoutingControlConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ResetCustomActionLambdaConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomActionLambdaConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ResetDocumentDbConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDocumentDbConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ResetEc2AsgCapacityIncreaseConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2AsgCapacityIncreaseConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ResetEcsCapacityIncreaseConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEcsCapacityIncreaseConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ResetEksResourceScalingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEksResourceScalingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ResetExecutionApprovalConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionApprovalConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ResetGlobalAuroraConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetGlobalAuroraConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ResetRegionSwitchPlanConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRegionSwitchPlanConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ResetRoute53HealthCheckConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRoute53HealthCheckConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

