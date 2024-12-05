package action_settings

import (
	"fmt"
	"strings"
	"wrench/app/manifest/types"
	"wrench/app/manifest/validation"
)

type HttpRequestSetting struct {
	Method            types.HttpMethod  `yaml:"method"`
	Url               string            `yaml:"url"`
	MapFixedHeaders   map[string]string `yaml:"mapFixedHeaders"`
	MapResponseHeader []string          `yaml:"mapResponseHeader"`
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

	if setting.MapResponseHeader != nil {
		for _, mapHeader := range setting.MapResponseHeader {
			mapSplitted := strings.Split(mapHeader, ":")
			if len(mapSplitted) > 2 {
				result.AddError("api.actions.request.mapResponseHeader should contains only one splitter ':'")
			}

			if len(mapHeader) == 0 {
				result.AddError("api.actions.request.mapResponseHeader itens can't contains empty values")
			}
		}
	}

	return result
}
