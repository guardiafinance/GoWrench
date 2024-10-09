package application_settings

import (
	"wrench/app/manifest/action_settings"
	"wrench/app/manifest/api_settings"
	"wrench/app/manifest/service_settings"
)

type ApplicationSetting struct {
	Api     api_settings.ApiSetting         `yaml:"api"`
	Actions []action_settings.ActionSetting `yaml:"actions"`
	Service service_settings.ServiceSetting `yaml:"service"`
}
