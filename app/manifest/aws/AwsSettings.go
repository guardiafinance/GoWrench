package aws

import (
	aws "wrench/app/manifest/aws/secrets"
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
