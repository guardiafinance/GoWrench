package application_settings

import (
	"wrench/app/manifest/action_settings"
	"wrench/app/manifest/api_settings"
	"wrench/app/manifest/service_settings"
	"wrench/app/manifest/validation"
)

type ApplicationSettings struct {
	Api     *api_settings.ApiSettings         `yaml:"api"`
	Actions []action_settings.ActionSettings  `yaml:"actions"`
	Service *service_settings.ServiceSettings `yaml:"service"`
}

func (settings ApplicationSettings) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if settings.Api != nil {
		result.AppendValidable(settings.Api)
	}

	if settings.Actions != nil {
		for _, validable := range settings.Actions {
			result.AppendValidable(validable)
		}
	}

	if settings.Api != nil {
		result.AppendValidable(settings.Api)
	}

	return result
}
