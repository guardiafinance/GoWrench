package trigger_settings

import (
	"wrench/app/manifest/validation"
)

type TriggerSetting struct {
	Before *BeforeSetting `yaml:"before"`
	After  *AfterSetting  `yaml:"after"`
}

func (setting TriggerSetting) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if setting.After != nil {
		result.AppendValidable(setting.After)
	}

	if setting.Before != nil {
		result.AppendValidable(setting.Before)
	}

	return result
}
