// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/gluecatalog/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference interface {
	cdktn.ComplexObject
	Compaction() *map[string]*string
	SetCompaction(val *map[string]*string)
	CompactionInput() *map[string]*string
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
	OrphanFileDeletion() *map[string]*string
	SetOrphanFileDeletion(val *map[string]*string)
	OrphanFileDeletionInput() *map[string]*string
	Retention() *map[string]*string
	SetRetention(val *map[string]*string)
	RetentionInput() *map[string]*string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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
	ResetCompaction()
	ResetOrphanFileDeletion()
	ResetRetention()
	ResetRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference
type jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) Compaction() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"compaction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) CompactionInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"compactionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) OrphanFileDeletion() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"orphanFileDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) OrphanFileDeletionInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"orphanFileDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) Retention() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) RetentionInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"retentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.glueCatalog.GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference_Override(g GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.glueCatalog.GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference)SetCompaction(val *map[string]*string) {
	if err := j.validateSetCompactionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compaction",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference)SetOrphanFileDeletion(val *map[string]*string) {
	if err := j.validateSetOrphanFileDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orphanFileDeletion",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference)SetRetention(val *map[string]*string) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) ResetCompaction() {
	_jsii_.InvokeVoid(
		g,
		"resetCompaction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) ResetOrphanFileDeletion() {
	_jsii_.InvokeVoid(
		g,
		"resetOrphanFileDeletion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) ResetRetention() {
	_jsii_.InvokeVoid(
		g,
		"resetRetention",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		g,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesIcebergOptimizationPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

