package application_settings

import "wrench/app/manifest/validation"

type BeforeStartupSettings struct {
	PathBashFiles []string `yaml:"pathBashFiles"`
}

func (settings BeforeStartupSettings) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	return result
}
