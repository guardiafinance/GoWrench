package application_settings

import "wrench/app/manifest/validation"

type TriggerSettings struct {
	BeforeStartup *BeforeStartupSettings `yaml:"beforeStartup"`
}

func (settings TriggerSettings) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if settings.BeforeStartup != nil {
		result.AppendValidable(settings.BeforeStartup)
	}

	return result
}
