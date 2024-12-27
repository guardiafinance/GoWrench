package aws_settings

import (
	aws "wrench/app/manifest/aws_settings/secrets_settings"
	"wrench/app/manifest/validation"
)

type AwsSettings struct {
	Region            string                 `default:"us-east-1" yaml:"region"`
	AwsSecretSettings *aws.AwsSecretSettings `yaml:"secret"`
}

func (setting AwsSettings) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if setting.AwsSecretSettings != nil {
		result.AppendValidable(setting.AwsSecretSettings)
	}

	return result
}
