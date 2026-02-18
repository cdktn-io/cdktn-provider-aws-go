// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package prometheusscraper

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrometheusScraperSourceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PrometheusScraperSourceList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrometheusScraperSourceList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrometheusScraperSourceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PrometheusScraperSourceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrometheusScraperSourceList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrometheusScraperSourceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrometheusScraperSourceListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

