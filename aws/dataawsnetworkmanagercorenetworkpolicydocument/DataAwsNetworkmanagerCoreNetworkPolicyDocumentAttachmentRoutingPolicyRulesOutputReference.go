// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkmanagercorenetworkpolicydocument

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/dataawsnetworkmanagercorenetworkpolicydocument/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference interface {
	cdktf.ComplexObject
	Action() DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesActionOutputReference
	ActionInput() *DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesAction
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
	Conditions() DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesConditionsList
	ConditionsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EdgeLocations() *[]*string
	SetEdgeLocations(val *[]*string)
	EdgeLocationsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RuleNumber() *float64
	SetRuleNumber(val *float64)
	RuleNumberInput() *float64
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
	PutAction(value *DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesAction)
	PutConditions(value interface{})
	ResetDescription()
	ResetEdgeLocations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference
type jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) Action() DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesActionOutputReference {
	var returns DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesActionOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) ActionInput() *DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesAction {
	var returns *DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesAction
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) Conditions() DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesConditionsList {
	var returns DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesConditionsList
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) ConditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) EdgeLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"edgeLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) EdgeLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"edgeLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) RuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) RuleNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsNetworkmanagerCoreNetworkPolicyDocument.DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference_Override(d DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsNetworkmanagerCoreNetworkPolicyDocument.DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference)SetEdgeLocations(val *[]*string) {
	if err := j.validateSetEdgeLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeLocations",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference)SetRuleNumber(val *float64) {
	if err := j.validateSetRuleNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleNumber",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) PutAction(value *DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesAction) {
	if err := d.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAction",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) PutConditions(value interface{}) {
	if err := d.validatePutConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConditions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) ResetEdgeLocations() {
	_jsii_.InvokeVoid(
		d,
		"resetEdgeLocations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

