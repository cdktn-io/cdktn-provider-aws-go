// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsdataprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/dmsdataprovider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsDataProviderSettingsOutputReference interface {
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
	DocDbSettings() DmsDataProviderSettingsDocDbSettingsList
	DocDbSettingsInput() interface{}
	// Experimental.
	Fqn() *string
	IbmDb2LuwSettings() DmsDataProviderSettingsIbmDb2LuwSettingsList
	IbmDb2LuwSettingsInput() interface{}
	IbmDb2ZosSettings() DmsDataProviderSettingsIbmDb2ZosSettingsList
	IbmDb2ZosSettingsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MariaDbSettings() DmsDataProviderSettingsMariaDbSettingsList
	MariaDbSettingsInput() interface{}
	MicrosoftSqlServerSettings() DmsDataProviderSettingsMicrosoftSqlServerSettingsList
	MicrosoftSqlServerSettingsInput() interface{}
	MongoDbSettings() DmsDataProviderSettingsMongoDbSettingsList
	MongoDbSettingsInput() interface{}
	MysqlSettings() DmsDataProviderSettingsMysqlSettingsList
	MysqlSettingsInput() interface{}
	OracleSettings() DmsDataProviderSettingsOracleSettingsList
	OracleSettingsInput() interface{}
	PostgresqlSettings() DmsDataProviderSettingsPostgresqlSettingsList
	PostgresqlSettingsInput() interface{}
	RedshiftSettings() DmsDataProviderSettingsRedshiftSettingsList
	RedshiftSettingsInput() interface{}
	SybaseAseSettings() DmsDataProviderSettingsSybaseAseSettingsList
	SybaseAseSettingsInput() interface{}
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
	PutDocDbSettings(value interface{})
	PutIbmDb2LuwSettings(value interface{})
	PutIbmDb2ZosSettings(value interface{})
	PutMariaDbSettings(value interface{})
	PutMicrosoftSqlServerSettings(value interface{})
	PutMongoDbSettings(value interface{})
	PutMysqlSettings(value interface{})
	PutOracleSettings(value interface{})
	PutPostgresqlSettings(value interface{})
	PutRedshiftSettings(value interface{})
	PutSybaseAseSettings(value interface{})
	ResetDocDbSettings()
	ResetIbmDb2LuwSettings()
	ResetIbmDb2ZosSettings()
	ResetMariaDbSettings()
	ResetMicrosoftSqlServerSettings()
	ResetMongoDbSettings()
	ResetMysqlSettings()
	ResetOracleSettings()
	ResetPostgresqlSettings()
	ResetRedshiftSettings()
	ResetSybaseAseSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsDataProviderSettingsOutputReference
type jsiiProxy_DmsDataProviderSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) DocDbSettings() DmsDataProviderSettingsDocDbSettingsList {
	var returns DmsDataProviderSettingsDocDbSettingsList
	_jsii_.Get(
		j,
		"docDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) DocDbSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"docDbSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) IbmDb2LuwSettings() DmsDataProviderSettingsIbmDb2LuwSettingsList {
	var returns DmsDataProviderSettingsIbmDb2LuwSettingsList
	_jsii_.Get(
		j,
		"ibmDb2LuwSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) IbmDb2LuwSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ibmDb2LuwSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) IbmDb2ZosSettings() DmsDataProviderSettingsIbmDb2ZosSettingsList {
	var returns DmsDataProviderSettingsIbmDb2ZosSettingsList
	_jsii_.Get(
		j,
		"ibmDb2ZosSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) IbmDb2ZosSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ibmDb2ZosSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) MariaDbSettings() DmsDataProviderSettingsMariaDbSettingsList {
	var returns DmsDataProviderSettingsMariaDbSettingsList
	_jsii_.Get(
		j,
		"mariaDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) MariaDbSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mariaDbSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) MicrosoftSqlServerSettings() DmsDataProviderSettingsMicrosoftSqlServerSettingsList {
	var returns DmsDataProviderSettingsMicrosoftSqlServerSettingsList
	_jsii_.Get(
		j,
		"microsoftSqlServerSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) MicrosoftSqlServerSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"microsoftSqlServerSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) MongoDbSettings() DmsDataProviderSettingsMongoDbSettingsList {
	var returns DmsDataProviderSettingsMongoDbSettingsList
	_jsii_.Get(
		j,
		"mongoDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) MongoDbSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongoDbSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) MysqlSettings() DmsDataProviderSettingsMysqlSettingsList {
	var returns DmsDataProviderSettingsMysqlSettingsList
	_jsii_.Get(
		j,
		"mysqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) MysqlSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mysqlSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) OracleSettings() DmsDataProviderSettingsOracleSettingsList {
	var returns DmsDataProviderSettingsOracleSettingsList
	_jsii_.Get(
		j,
		"oracleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) OracleSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oracleSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) PostgresqlSettings() DmsDataProviderSettingsPostgresqlSettingsList {
	var returns DmsDataProviderSettingsPostgresqlSettingsList
	_jsii_.Get(
		j,
		"postgresqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) PostgresqlSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postgresqlSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) RedshiftSettings() DmsDataProviderSettingsRedshiftSettingsList {
	var returns DmsDataProviderSettingsRedshiftSettingsList
	_jsii_.Get(
		j,
		"redshiftSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) RedshiftSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) SybaseAseSettings() DmsDataProviderSettingsSybaseAseSettingsList {
	var returns DmsDataProviderSettingsSybaseAseSettingsList
	_jsii_.Get(
		j,
		"sybaseAseSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) SybaseAseSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sybaseAseSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDmsDataProviderSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DmsDataProviderSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsDataProviderSettingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsDataProviderSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.dmsDataProvider.DmsDataProviderSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDmsDataProviderSettingsOutputReference_Override(d DmsDataProviderSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.dmsDataProvider.DmsDataProviderSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) PutDocDbSettings(value interface{}) {
	if err := d.validatePutDocDbSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDocDbSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) PutIbmDb2LuwSettings(value interface{}) {
	if err := d.validatePutIbmDb2LuwSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIbmDb2LuwSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) PutIbmDb2ZosSettings(value interface{}) {
	if err := d.validatePutIbmDb2ZosSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIbmDb2ZosSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) PutMariaDbSettings(value interface{}) {
	if err := d.validatePutMariaDbSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMariaDbSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) PutMicrosoftSqlServerSettings(value interface{}) {
	if err := d.validatePutMicrosoftSqlServerSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMicrosoftSqlServerSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) PutMongoDbSettings(value interface{}) {
	if err := d.validatePutMongoDbSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMongoDbSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) PutMysqlSettings(value interface{}) {
	if err := d.validatePutMysqlSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMysqlSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) PutOracleSettings(value interface{}) {
	if err := d.validatePutOracleSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOracleSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) PutPostgresqlSettings(value interface{}) {
	if err := d.validatePutPostgresqlSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPostgresqlSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) PutRedshiftSettings(value interface{}) {
	if err := d.validatePutRedshiftSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRedshiftSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) PutSybaseAseSettings(value interface{}) {
	if err := d.validatePutSybaseAseSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSybaseAseSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ResetDocDbSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetDocDbSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ResetIbmDb2LuwSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetIbmDb2LuwSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ResetIbmDb2ZosSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetIbmDb2ZosSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ResetMariaDbSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetMariaDbSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ResetMicrosoftSqlServerSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetMicrosoftSqlServerSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ResetMongoDbSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetMongoDbSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ResetMysqlSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetMysqlSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ResetOracleSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetOracleSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ResetPostgresqlSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgresqlSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ResetRedshiftSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetRedshiftSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ResetSybaseAseSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetSybaseAseSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

