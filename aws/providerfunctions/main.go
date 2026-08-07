// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package providerfunctions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.providerFunctions.AwsProviderFunctions",
		reflect.TypeOf((*AwsProviderFunctions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "arnBuild", GoMethod: "ArnBuild"},
			_jsii_.MemberMethod{JsiiMethod: "arnParse", GoMethod: "ArnParse"},
			_jsii_.MemberMethod{JsiiMethod: "trimIamRolePath", GoMethod: "TrimIamRolePath"},
			_jsii_.MemberMethod{JsiiMethod: "userAgent", GoMethod: "UserAgent"},
		},
		func() interface{} {
			return &jsiiProxy_AwsProviderFunctions{}
		},
	)
}
