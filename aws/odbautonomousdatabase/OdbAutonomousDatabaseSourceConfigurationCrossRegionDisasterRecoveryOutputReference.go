// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/odbautonomousdatabase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsReplicateAutomaticBackups() interface{}
	SetIsReplicateAutomaticBackups(val interface{})
	IsReplicateAutomaticBackupsInput() interface{}
	RemoteDisasterRecoveryType() *string
	SetRemoteDisasterRecoveryType(val *string)
	RemoteDisasterRecoveryTypeInput() *string
	SourceAutonomousDatabaseArn() *string
	SetSourceAutonomousDatabaseArn(val *string)
	SourceAutonomousDatabaseArnInput() *string
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
	ResetIsReplicateAutomaticBackups()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference
type jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) IsReplicateAutomaticBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isReplicateAutomaticBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) IsReplicateAutomaticBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isReplicateAutomaticBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) RemoteDisasterRecoveryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteDisasterRecoveryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) RemoteDisasterRecoveryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteDisasterRecoveryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) SourceAutonomousDatabaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAutonomousDatabaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) SourceAutonomousDatabaseArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAutonomousDatabaseArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference {
	_init_.Initialize()

	if err := validateNewOdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference_Override(o OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference)SetIsReplicateAutomaticBackups(val interface{}) {
	if err := j.validateSetIsReplicateAutomaticBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isReplicateAutomaticBackups",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference)SetRemoteDisasterRecoveryType(val *string) {
	if err := j.validateSetRemoteDisasterRecoveryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteDisasterRecoveryType",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference)SetSourceAutonomousDatabaseArn(val *string) {
	if err := j.validateSetSourceAutonomousDatabaseArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAutonomousDatabaseArn",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) ResetIsReplicateAutomaticBackups() {
	_jsii_.InvokeVoid(
		o,
		"resetIsReplicateAutomaticBackups",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

