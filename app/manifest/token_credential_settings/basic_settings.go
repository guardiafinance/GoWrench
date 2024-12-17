package token_credential_settings

import "wrench/app/manifest/validation"

type BasicSetting struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (setting BasicSetting) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if len(setting.Username) == 0 {
		result.AddError("tokenCredentials.basic.username is required")
	}

	if len(setting.Password) == 0 {
		result.AddError("tokenCredentials.basic.password is required")
	}

	return result
}
