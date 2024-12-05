package action_settings

import (
	"fmt"
	"wrench/app/manifest/validation"
)

type ActionSettings struct {
	Id      string                  `yaml:"id"`
	Type    ActionType              `yaml:"type"`
	Request *HttpRequestSetting     `yaml:"request"`
	Mock    *HttpRequestMockSetting `yaml:"mock"`
}

type ActionType string

const (
	ActionTypeHttpRequest     ActionType = "httpRequest"
	ActionTypeHttpRequestMock ActionType = "httpRequestMock"
)

func (setting ActionSettings) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if len(setting.Id) == 0 {
		result.AddError("actions.id is required")
	}

	if len(setting.Type) == 0 {
		var msg = fmt.Sprintf("actions[%s].type is required", setting.Id)
		result.AddError(msg)
	} else {
		if (setting.Type == ActionTypeHttpRequest ||
			setting.Type == ActionTypeHttpRequestMock) == false {

			var msg = fmt.Sprintf("actions[%s].type should contain valid value", setting.Id)
			result.AddError(msg)
		}
	}

	if setting.Type == ActionTypeHttpRequest {
		setting.ValidTypeActionTypeHttpRequest(&result)
	}

	if setting.Type == ActionTypeHttpRequestMock {
		setting.ValidTypeActionTypeHttpRequestMock(&result)
	}

	return result
}

func (setting ActionSettings) ValidTypeActionTypeHttpRequest(result *validation.ValidateResult) {

	if setting.Request == nil {
		var msg = fmt.Sprintf("actions[%s].request should be nil", setting.Id)
		result.AddError(msg)
	} else {
		result.AppendValidable(setting.Request)
	}

	if setting.Mock != nil {
		var msg = fmt.Sprintf("actions[%s].mock should be nil", setting.Id)
		result.AddError(msg)
	}
}

func (setting ActionSettings) ValidTypeActionTypeHttpRequestMock(result *validation.ValidateResult) {

	if setting.Mock != nil {
		result.AppendValidable(setting.Mock)
	} else {
		var msg = fmt.Sprintf("actions[%s].mock is required", setting.Id)
		result.AddError(msg)
	}
}
