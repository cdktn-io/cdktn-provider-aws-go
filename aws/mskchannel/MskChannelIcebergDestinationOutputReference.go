// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/mskchannel/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MskChannelIcebergDestinationOutputReference interface {
	cdktn.ComplexObject
	AppendOnly() interface{}
	SetAppendOnly(val interface{})
	AppendOnlyInput() interface{}
	Catalog() MskChannelIcebergDestinationCatalogList
	CatalogInput() interface{}
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
	CompressionType() *string
	SetCompressionType(val *string)
	CompressionTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataFreshnessInSeconds() *float64
	SetDataFreshnessInSeconds(val *float64)
	DataFreshnessInSecondsInput() *float64
	DeadLetterQueueS3() MskChannelIcebergDestinationDeadLetterQueueS3List
	DeadLetterQueueS3Input() interface{}
	DestinationTable() MskChannelIcebergDestinationDestinationTableList
	DestinationTableInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SchemaEvolution() MskChannelIcebergDestinationSchemaEvolutionList
	SchemaEvolutionInput() interface{}
	ServiceExecutionRoleArn() *string
	SetServiceExecutionRoleArn(val *string)
	ServiceExecutionRoleArnInput() *string
	TableCreation() MskChannelIcebergDestinationTableCreationList
	TableCreationInput() interface{}
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
	PutCatalog(value interface{})
	PutDeadLetterQueueS3(value interface{})
	PutDestinationTable(value interface{})
	PutSchemaEvolution(value interface{})
	PutTableCreation(value interface{})
	ResetCatalog()
	ResetCompressionType()
	ResetDataFreshnessInSeconds()
	ResetDeadLetterQueueS3()
	ResetDestinationTable()
	ResetSchemaEvolution()
	ResetTableCreation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MskChannelIcebergDestinationOutputReference
type jsiiProxy_MskChannelIcebergDestinationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) AppendOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appendOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) AppendOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appendOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) Catalog() MskChannelIcebergDestinationCatalogList {
	var returns MskChannelIcebergDestinationCatalogList
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) CatalogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) DataFreshnessInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataFreshnessInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) DataFreshnessInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataFreshnessInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) DeadLetterQueueS3() MskChannelIcebergDestinationDeadLetterQueueS3List {
	var returns MskChannelIcebergDestinationDeadLetterQueueS3List
	_jsii_.Get(
		j,
		"deadLetterQueueS3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) DeadLetterQueueS3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterQueueS3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) DestinationTable() MskChannelIcebergDestinationDestinationTableList {
	var returns MskChannelIcebergDestinationDestinationTableList
	_jsii_.Get(
		j,
		"destinationTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) DestinationTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) SchemaEvolution() MskChannelIcebergDestinationSchemaEvolutionList {
	var returns MskChannelIcebergDestinationSchemaEvolutionList
	_jsii_.Get(
		j,
		"schemaEvolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) SchemaEvolutionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaEvolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) ServiceExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) ServiceExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) TableCreation() MskChannelIcebergDestinationTableCreationList {
	var returns MskChannelIcebergDestinationTableCreationList
	_jsii_.Get(
		j,
		"tableCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) TableCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMskChannelIcebergDestinationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MskChannelIcebergDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewMskChannelIcebergDestinationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MskChannelIcebergDestinationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.mskChannel.MskChannelIcebergDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMskChannelIcebergDestinationOutputReference_Override(m MskChannelIcebergDestinationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.mskChannel.MskChannelIcebergDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference)SetAppendOnly(val interface{}) {
	if err := j.validateSetAppendOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appendOnly",
		val,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference)SetCompressionType(val *string) {
	if err := j.validateSetCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference)SetDataFreshnessInSeconds(val *float64) {
	if err := j.validateSetDataFreshnessInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFreshnessInSeconds",
		val,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference)SetServiceExecutionRoleArn(val *string) {
	if err := j.validateSetServiceExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MskChannelIcebergDestinationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) PutCatalog(value interface{}) {
	if err := m.validatePutCatalogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCatalog",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) PutDeadLetterQueueS3(value interface{}) {
	if err := m.validatePutDeadLetterQueueS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDeadLetterQueueS3",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) PutDestinationTable(value interface{}) {
	if err := m.validatePutDestinationTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDestinationTable",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) PutSchemaEvolution(value interface{}) {
	if err := m.validatePutSchemaEvolutionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSchemaEvolution",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) PutTableCreation(value interface{}) {
	if err := m.validatePutTableCreationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTableCreation",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) ResetCatalog() {
	_jsii_.InvokeVoid(
		m,
		"resetCatalog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		m,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) ResetDataFreshnessInSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetDataFreshnessInSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) ResetDeadLetterQueueS3() {
	_jsii_.InvokeVoid(
		m,
		"resetDeadLetterQueueS3",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) ResetDestinationTable() {
	_jsii_.InvokeVoid(
		m,
		"resetDestinationTable",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) ResetSchemaEvolution() {
	_jsii_.InvokeVoid(
		m,
		"resetSchemaEvolution",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) ResetTableCreation() {
	_jsii_.InvokeVoid(
		m,
		"resetTableCreation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskChannelIcebergDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

