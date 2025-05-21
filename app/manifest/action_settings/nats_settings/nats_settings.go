package nats_settings

import "wrench/app/manifest/validation"

type NatsSettings struct {
	ConnectionId string `yaml:"connectionId"`
	IsStream     bool   `yaml:"isStream"`
	SubjectName  string `yaml:"subjectName"`
}

func (settings NatsSettings) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if len(settings.ConnectionId) == 0 {
		result.AddError("actions.nats.connectionId is required")
	}

	if len(settings.SubjectName) == 0 {
		result.AddError("actions.nats.subjectName is required")
	}

	return result
}
