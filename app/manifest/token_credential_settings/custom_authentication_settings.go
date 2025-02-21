package token_credential_settings

import (
	"wrench/app/manifest/validation"
)

type CustomAuthenticationConfigurations struct {
	Token string `yaml:"token"`
}

type CustomAuthentication struct {
	RequestBody map[string]string                  `yaml:"requestBody"`
	Configs     CustomAuthenticationConfigurations `yaml:"configs"`
}

func (setting CustomAuthentication) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if len(setting.RequestBody) <= 0 {
		result.AddError("tokenCredentials.custom_authentication.custom.requestBody is required")
	}

	if len(setting.Configs.Token) == 0 {
		result.AddError("tokenCredentials.custom_authentication.custom.configs.token is required")
	}

	return result
}
