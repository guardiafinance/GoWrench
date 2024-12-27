package application_settings

import (
	"errors"
	"wrench/app/manifest/action_settings"
	"wrench/app/manifest/api_settings"
	aws "wrench/app/manifest/aws_settings"
	"wrench/app/manifest/contract_settings"
	"wrench/app/manifest/service_settings"
	credential "wrench/app/manifest/token_credential_settings"
	"wrench/app/manifest/validation"

	"gopkg.in/yaml.v3"
)

var ApplicationSettingsStatic *ApplicationSettings

type ApplicationSettings struct {
	Api              *api_settings.ApiSettings            `yaml:"api"`
	Actions          []action_settings.ActionSettings     `yaml:"actions"`
	Service          *service_settings.ServiceSettings    `yaml:"service"`
	TokenCredentials []*credential.TokenCredentialSetting `yaml:"tokenCredentials"`
	Contract         *contract_settings.ContractSetting   `yaml:"contract"`
	Aws              *aws.AwsSettings                     `yaml:"aws"`
}

func (settings ApplicationSettings) GetActionById(actionId string) (*action_settings.ActionSettings, error) {
	for _, action := range settings.Actions {
		if action.Id == actionId {
			return &action, nil
		}
	}

	return nil, errors.New("action not found")
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

	if settings.TokenCredentials != nil {
		for _, validable := range settings.TokenCredentials {
			result.AppendValidable(validable)
		}
	}

	if settings.Contract != nil {
		result.AppendValidable(settings.Contract)
	}

	return result
}

func ParseToApplicationSetting(data []byte) (*ApplicationSettings, error) {

	applicationSettings := new(ApplicationSettings)
	err := yaml.Unmarshal(data, applicationSettings)
	if err != nil {
		return nil, err
	}
	return applicationSettings, nil
}
