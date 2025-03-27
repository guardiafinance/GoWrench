package api_settings

import "wrench/app/manifest/validation"

type CorsSettings struct {
	Origins []string `yaml:"origins"`
	Methods []string `yaml:"methods"`
	Headers []string `yaml:"headers"`
}

func (setting CorsSettings) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	return result
}
