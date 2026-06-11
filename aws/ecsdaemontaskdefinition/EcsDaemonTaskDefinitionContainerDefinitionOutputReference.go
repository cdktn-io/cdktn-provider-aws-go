// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/ecsdaemontaskdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsDaemonTaskDefinitionContainerDefinitionOutputReference interface {
	cdktn.ComplexObject
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	Cpu() *float64
	SetCpu(val *float64)
	CpuInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DependsOn() EcsDaemonTaskDefinitionContainerDefinitionDependsOnList
	DependsOnInput() interface{}
	EntryPoint() *[]*string
	SetEntryPoint(val *[]*string)
	EntryPointInput() *[]*string
	Environment() EcsDaemonTaskDefinitionContainerDefinitionEnvironmentList
	EnvironmentFile() EcsDaemonTaskDefinitionContainerDefinitionEnvironmentFileList
	EnvironmentFileInput() interface{}
	EnvironmentInput() interface{}
	Essential() interface{}
	SetEssential(val interface{})
	EssentialInput() interface{}
	FirelensConfiguration() EcsDaemonTaskDefinitionContainerDefinitionFirelensConfigurationList
	FirelensConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	HealthCheck() EcsDaemonTaskDefinitionContainerDefinitionHealthCheckList
	HealthCheckInput() interface{}
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	Interactive() interface{}
	SetInteractive(val interface{})
	InteractiveInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LinuxParameters() EcsDaemonTaskDefinitionContainerDefinitionLinuxParametersList
	LinuxParametersInput() interface{}
	LogConfiguration() EcsDaemonTaskDefinitionContainerDefinitionLogConfigurationList
	LogConfigurationInput() interface{}
	Memory() *float64
	SetMemory(val *float64)
	MemoryInput() *float64
	MemoryReservation() *float64
	SetMemoryReservation(val *float64)
	MemoryReservationInput() *float64
	MountPoint() EcsDaemonTaskDefinitionContainerDefinitionMountPointList
	MountPointInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Privileged() interface{}
	SetPrivileged(val interface{})
	PrivilegedInput() interface{}
	PseudoTerminal() interface{}
	SetPseudoTerminal(val interface{})
	PseudoTerminalInput() interface{}
	ReadonlyRootFilesystem() interface{}
	SetReadonlyRootFilesystem(val interface{})
	ReadonlyRootFilesystemInput() interface{}
	RepositoryCredentials() EcsDaemonTaskDefinitionContainerDefinitionRepositoryCredentialsList
	RepositoryCredentialsInput() interface{}
	RestartPolicy() EcsDaemonTaskDefinitionContainerDefinitionRestartPolicyList
	RestartPolicyInput() interface{}
	Secret() EcsDaemonTaskDefinitionContainerDefinitionSecretList
	SecretInput() interface{}
	StartTimeout() *float64
	SetStartTimeout(val *float64)
	StartTimeoutInput() *float64
	StopTimeout() *float64
	SetStopTimeout(val *float64)
	StopTimeoutInput() *float64
	SystemControl() EcsDaemonTaskDefinitionContainerDefinitionSystemControlList
	SystemControlInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Ulimit() EcsDaemonTaskDefinitionContainerDefinitionUlimitList
	UlimitInput() interface{}
	User() *string
	SetUser(val *string)
	UserInput() *string
	WorkingDirectory() *string
	SetWorkingDirectory(val *string)
	WorkingDirectoryInput() *string
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
	PutDependsOn(value interface{})
	PutEnvironment(value interface{})
	PutEnvironmentFile(value interface{})
	PutFirelensConfiguration(value interface{})
	PutHealthCheck(value interface{})
	PutLinuxParameters(value interface{})
	PutLogConfiguration(value interface{})
	PutMountPoint(value interface{})
	PutRepositoryCredentials(value interface{})
	PutRestartPolicy(value interface{})
	PutSecret(value interface{})
	PutSystemControl(value interface{})
	PutUlimit(value interface{})
	ResetCommand()
	ResetCpu()
	ResetDependsOn()
	ResetEntryPoint()
	ResetEnvironment()
	ResetEnvironmentFile()
	ResetEssential()
	ResetFirelensConfiguration()
	ResetHealthCheck()
	ResetInteractive()
	ResetLinuxParameters()
	ResetLogConfiguration()
	ResetMemory()
	ResetMemoryReservation()
	ResetMountPoint()
	ResetName()
	ResetPrivileged()
	ResetPseudoTerminal()
	ResetReadonlyRootFilesystem()
	ResetRepositoryCredentials()
	ResetRestartPolicy()
	ResetSecret()
	ResetStartTimeout()
	ResetStopTimeout()
	ResetSystemControl()
	ResetUlimit()
	ResetUser()
	ResetWorkingDirectory()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsDaemonTaskDefinitionContainerDefinitionOutputReference
type jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) CpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) DependsOn() EcsDaemonTaskDefinitionContainerDefinitionDependsOnList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionDependsOnList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) DependsOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) EntryPoint() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) EntryPointInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) Environment() EcsDaemonTaskDefinitionContainerDefinitionEnvironmentList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) EnvironmentFile() EcsDaemonTaskDefinitionContainerDefinitionEnvironmentFileList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionEnvironmentFileList
	_jsii_.Get(
		j,
		"environmentFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) EnvironmentFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) Essential() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"essential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) EssentialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"essentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) FirelensConfiguration() EcsDaemonTaskDefinitionContainerDefinitionFirelensConfigurationList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionFirelensConfigurationList
	_jsii_.Get(
		j,
		"firelensConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) FirelensConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firelensConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) HealthCheck() EcsDaemonTaskDefinitionContainerDefinitionHealthCheckList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionHealthCheckList
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) HealthCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) Interactive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"interactive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) InteractiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"interactiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) LinuxParameters() EcsDaemonTaskDefinitionContainerDefinitionLinuxParametersList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionLinuxParametersList
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) LinuxParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linuxParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) LogConfiguration() EcsDaemonTaskDefinitionContainerDefinitionLogConfigurationList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionLogConfigurationList
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) LogConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) MemoryReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) MemoryReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) MountPoint() EcsDaemonTaskDefinitionContainerDefinitionMountPointList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionMountPointList
	_jsii_.Get(
		j,
		"mountPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) MountPointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mountPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) Privileged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) PrivilegedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) PseudoTerminal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pseudoTerminal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) PseudoTerminalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pseudoTerminalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ReadonlyRootFilesystem() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ReadonlyRootFilesystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyRootFilesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) RepositoryCredentials() EcsDaemonTaskDefinitionContainerDefinitionRepositoryCredentialsList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionRepositoryCredentialsList
	_jsii_.Get(
		j,
		"repositoryCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) RepositoryCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) RestartPolicy() EcsDaemonTaskDefinitionContainerDefinitionRestartPolicyList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionRestartPolicyList
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) RestartPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) Secret() EcsDaemonTaskDefinitionContainerDefinitionSecretList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) StartTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) StartTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) StopTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stopTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) StopTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stopTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) SystemControl() EcsDaemonTaskDefinitionContainerDefinitionSystemControlList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionSystemControlList
	_jsii_.Get(
		j,
		"systemControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) SystemControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) Ulimit() EcsDaemonTaskDefinitionContainerDefinitionUlimitList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionUlimitList
	_jsii_.Get(
		j,
		"ulimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) UlimitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ulimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) WorkingDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) WorkingDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectoryInput",
		&returns,
	)
	return returns
}


func NewEcsDaemonTaskDefinitionContainerDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EcsDaemonTaskDefinitionContainerDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewEcsDaemonTaskDefinitionContainerDefinitionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEcsDaemonTaskDefinitionContainerDefinitionOutputReference_Override(e EcsDaemonTaskDefinitionContainerDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetCpu(val *float64) {
	if err := j.validateSetCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetEntryPoint(val *[]*string) {
	if err := j.validateSetEntryPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryPoint",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetEssential(val interface{}) {
	if err := j.validateSetEssentialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"essential",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetInteractive(val interface{}) {
	if err := j.validateSetInteractiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interactive",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetMemory(val *float64) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetMemoryReservation(val *float64) {
	if err := j.validateSetMemoryReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryReservation",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetPrivileged(val interface{}) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetPseudoTerminal(val interface{}) {
	if err := j.validateSetPseudoTerminalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pseudoTerminal",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetReadonlyRootFilesystem(val interface{}) {
	if err := j.validateSetReadonlyRootFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readonlyRootFilesystem",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetStartTimeout(val *float64) {
	if err := j.validateSetStartTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTimeout",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetStopTimeout(val *float64) {
	if err := j.validateSetStopTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stopTimeout",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference)SetWorkingDirectory(val *string) {
	if err := j.validateSetWorkingDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDirectory",
		val,
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) PutDependsOn(value interface{}) {
	if err := e.validatePutDependsOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDependsOn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) PutEnvironment(value interface{}) {
	if err := e.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) PutEnvironmentFile(value interface{}) {
	if err := e.validatePutEnvironmentFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEnvironmentFile",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) PutFirelensConfiguration(value interface{}) {
	if err := e.validatePutFirelensConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putFirelensConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) PutHealthCheck(value interface{}) {
	if err := e.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) PutLinuxParameters(value interface{}) {
	if err := e.validatePutLinuxParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLinuxParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) PutLogConfiguration(value interface{}) {
	if err := e.validatePutLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLogConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) PutMountPoint(value interface{}) {
	if err := e.validatePutMountPointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMountPoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) PutRepositoryCredentials(value interface{}) {
	if err := e.validatePutRepositoryCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRepositoryCredentials",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) PutRestartPolicy(value interface{}) {
	if err := e.validatePutRestartPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRestartPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) PutSecret(value interface{}) {
	if err := e.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSecret",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) PutSystemControl(value interface{}) {
	if err := e.validatePutSystemControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSystemControl",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) PutUlimit(value interface{}) {
	if err := e.validatePutUlimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putUlimit",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		e,
		"resetCommand",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetCpu() {
	_jsii_.InvokeVoid(
		e,
		"resetCpu",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		e,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetEntryPoint() {
	_jsii_.InvokeVoid(
		e,
		"resetEntryPoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		e,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetEnvironmentFile() {
	_jsii_.InvokeVoid(
		e,
		"resetEnvironmentFile",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetEssential() {
	_jsii_.InvokeVoid(
		e,
		"resetEssential",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetFirelensConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetFirelensConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetInteractive() {
	_jsii_.InvokeVoid(
		e,
		"resetInteractive",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetLinuxParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetLinuxParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetLogConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetLogConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetMemory() {
	_jsii_.InvokeVoid(
		e,
		"resetMemory",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetMemoryReservation() {
	_jsii_.InvokeVoid(
		e,
		"resetMemoryReservation",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetMountPoint() {
	_jsii_.InvokeVoid(
		e,
		"resetMountPoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetPseudoTerminal() {
	_jsii_.InvokeVoid(
		e,
		"resetPseudoTerminal",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetReadonlyRootFilesystem() {
	_jsii_.InvokeVoid(
		e,
		"resetReadonlyRootFilesystem",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetRepositoryCredentials() {
	_jsii_.InvokeVoid(
		e,
		"resetRepositoryCredentials",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetRestartPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetRestartPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		e,
		"resetSecret",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetStartTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetStartTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetStopTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetStopTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetSystemControl() {
	_jsii_.InvokeVoid(
		e,
		"resetSystemControl",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetUlimit() {
	_jsii_.InvokeVoid(
		e,
		"resetUlimit",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		e,
		"resetUser",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ResetWorkingDirectory() {
	_jsii_.InvokeVoid(
		e,
		"resetWorkingDirectory",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

