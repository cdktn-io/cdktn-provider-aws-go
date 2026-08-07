// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtableoptimizer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/gluecatalogtableoptimizer/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList
type jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList {
	_init_.Initialize()

	if err := validateNewGlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList{}

	_jsii_.Create(
		"@cdktn/provider-aws.glueCatalogTableOptimizer.GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList_Override(g GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.glueCatalogTableOptimizer.GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := g.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		g,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList) Get(index *float64) GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfigurationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

