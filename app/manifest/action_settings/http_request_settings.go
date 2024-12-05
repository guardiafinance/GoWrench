package action_settings

import (
	"fmt"
	"wrench/app/manifest/types"
	"wrench/app/manifest/validation"
)

type HttpRequestSetting struct {
	Method       types.HttpMethod  `yaml:"method"`
	Url          string            `yaml:"url"`
	FixedHeaders map[string]string `yaml:"fixedHeaders"`
}

func (setting HttpRequestSetting) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if len(setting.Method) == 0 {
		var msg = fmt.Sprintf("api.actions.request.method is required")
		result.AddError(msg)
	} else {
		if (setting.Method == types.HttpMethodGet ||
			setting.Method == types.HttpMethodPost ||
			setting.Method == types.HttpMethodPut ||
			setting.Method == types.HttpMethodPatch ||
			setting.Method == types.HttpMethodDelete) == false {

			var msg = fmt.Sprintf("api.actions.request.method should contain valid value (get, post, put, patch or delete)")
			result.AddError(msg)
		}
	}

	if len(setting.Url) == 0 {
		result.AddError("api.actions.request.url is required")
	}

	return result
}
