package sns_settings

import (
	"strings"
	"wrench/app/manifest/validation"
)

type SnsSettings struct {
	TopicArn        string            `yaml:"topicArn"`
	GroupId         string            `yaml:"groupId"`
	MapHeaders      []string          `yaml:"mapHeaders"`
	MapFixedHeaders map[string]string `yaml:"mapFixedHeaders"`
}

func (setting SnsSettings) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if setting.MapFixedHeaders != nil {
		for _, mapHeader := range setting.MapFixedHeaders {
			mapSplitted := strings.Split(mapHeader, ":")
			if len(mapSplitted) != 2 {
				result.AddError("actions.sns.mapFixedHeaders invalid")
			}
			if len(mapSplitted[0]) == 0 {
				result.AddError("actions.sns.mapFixedHeaders header key is required")
			}
		}
	}

	if setting.MapHeaders != nil {
		for _, mapHeader := range setting.MapHeaders {
			mapSplitted := strings.Split(mapHeader, ":")
			if len(mapSplitted) > 2 {
				result.AddError("actions.sns.MapHeaders should contains only one splitter ':'")
			}

			if len(mapHeader) == 0 {
				result.AddError("actions.sns.mapHeaders itens can't contains empty values")
			}
		}
	}

	return result
}
