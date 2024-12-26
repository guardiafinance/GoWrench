package sns_settings

import (
	"strings"
	"wrench/app/manifest/validation"
)

type SnsSettings struct {
	TopicArn string        `yaml:"topicArn"`
	Fifo     *FifoSettings `yaml:"fifo"`
	Filters  []string      `yaml:"filters"`
}

func (settings *SnsSettings) IsFifo() bool {
	return strings.HasSuffix(settings.TopicArn, ".fifo")
}

func (setting SnsSettings) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if setting.Fifo != nil && setting.IsFifo() {
		result.AddError("actions.sns.fifo can't be configured when topic isn't fifo")
	}

	if setting.IsFifo() && setting.Fifo == nil {
		result.AddError("actions.sns.fifo should be configured when topic is fifo")
	}

	return result
}
