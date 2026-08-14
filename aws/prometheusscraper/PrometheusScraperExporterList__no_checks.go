// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package prometheusscraper

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrometheusScraperExporterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PrometheusScraperExporterList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrometheusScraperExporterList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrometheusScraperExporterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PrometheusScraperExporterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrometheusScraperExporterList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrometheusScraperExporterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrometheusScraperExporterListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

