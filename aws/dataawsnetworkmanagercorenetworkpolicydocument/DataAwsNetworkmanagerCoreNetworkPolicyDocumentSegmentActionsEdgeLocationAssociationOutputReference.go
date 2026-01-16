// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkmanagercorenetworkpolicydocument

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/dataawsnetworkmanagercorenetworkpolicydocument/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference interface {
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
	EdgeLocation() *string
	SetEdgeLocation(val *string)
	EdgeLocationInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociation
	SetInternalValue(val *DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociation)
	PeerEdgeLocation() *string
	SetPeerEdgeLocation(val *string)
	PeerEdgeLocationInput() *string
	RoutingPolicyNames() *[]*string
	SetRoutingPolicyNames(val *[]*string)
	RoutingPolicyNamesInput() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference
type jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) EdgeLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) EdgeLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) InternalValue() *DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociation {
	var returns *DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) PeerEdgeLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerEdgeLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) PeerEdgeLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerEdgeLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) RoutingPolicyNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routingPolicyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) RoutingPolicyNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routingPolicyNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsNetworkmanagerCoreNetworkPolicyDocument.DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference_Override(d DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsNetworkmanagerCoreNetworkPolicyDocument.DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference)SetEdgeLocation(val *string) {
	if err := j.validateSetEdgeLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeLocation",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference)SetInternalValue(val *DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference)SetPeerEdgeLocation(val *string) {
	if err := j.validateSetPeerEdgeLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerEdgeLocation",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference)SetRoutingPolicyNames(val *[]*string) {
	if err := j.validateSetRoutingPolicyNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingPolicyNames",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

