package secrets_settings

import "wrench/app/manifest/validation"

type AwsSecretSettings struct {
	SecretsName []string `yaml:"secretsName"`
}

func (setting AwsSecretSettings) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	return result
}
