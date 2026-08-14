// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/arcregionswitchplan/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference interface {
	cdktn.ComplexObject
	Behavior() *string
	SetBehavior(val *string)
	BehaviorInput() *string
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
	CrossAccountRole() *string
	SetCrossAccountRole(val *string)
	CrossAccountRoleInput() *string
	ExternalId() *string
	SetExternalId(val *string)
	ExternalIdInput() *string
	// Experimental.
	Fqn() *string
	GlobalClusterIdentifier() *string
	SetGlobalClusterIdentifier(val *string)
	GlobalClusterIdentifierInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RegionDatabaseClusterArns() *map[string]*string
	SetRegionDatabaseClusterArns(val *map[string]*string)
	RegionDatabaseClusterArnsInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeoutMinutes() *float64
	SetTimeoutMinutes(val *float64)
	TimeoutMinutesInput() *float64
	Ungraceful() ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigUngracefulList
	UngracefulInput() interface{}
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
	PutUngraceful(value interface{})
	ResetCrossAccountRole()
	ResetExternalId()
	ResetTimeoutMinutes()
	ResetUngraceful()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference
type jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) BehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) CrossAccountRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossAccountRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) CrossAccountRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossAccountRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) GlobalClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) GlobalClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) RegionDatabaseClusterArns() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"regionDatabaseClusterArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) RegionDatabaseClusterArnsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"regionDatabaseClusterArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) TimeoutMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) TimeoutMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) Ungraceful() ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigUngracefulList {
	var returns ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigUngracefulList
	_jsii_.Get(
		j,
		"ungraceful",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) UngracefulInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ungracefulInput",
		&returns,
	)
	return returns
}


func NewArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference {
	_init_.Initialize()

	if err := validateNewArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.arcregionswitchPlan.ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference_Override(a ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.arcregionswitchPlan.ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference)SetBehavior(val *string) {
	if err := j.validateSetBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"behavior",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference)SetCrossAccountRole(val *string) {
	if err := j.validateSetCrossAccountRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossAccountRole",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference)SetExternalId(val *string) {
	if err := j.validateSetExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference)SetGlobalClusterIdentifier(val *string) {
	if err := j.validateSetGlobalClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference)SetRegionDatabaseClusterArns(val *map[string]*string) {
	if err := j.validateSetRegionDatabaseClusterArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionDatabaseClusterArns",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference)SetTimeoutMinutes(val *float64) {
	if err := j.validateSetTimeoutMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutMinutes",
		val,
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) PutUngraceful(value interface{}) {
	if err := a.validatePutUngracefulParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUngraceful",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) ResetCrossAccountRole() {
	_jsii_.InvokeVoid(
		a,
		"resetCrossAccountRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) ResetExternalId() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) ResetTimeoutMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) ResetUngraceful() {
	_jsii_.InvokeVoid(
		a,
		"resetUngraceful",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepNeptuneGlobalDatabaseConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

