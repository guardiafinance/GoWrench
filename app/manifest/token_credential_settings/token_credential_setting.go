package token_credential_settings

import (
	"strings"
	"wrench/app/manifest/validation"
)

type TokenCredentialSetting struct {
	Id               string                   `yaml:"id"`
	Type             TokenCredentialType      `yaml:"type"`
	AuthEndpoint     string                   `yaml:"authEndpoint"`
	IsOpaque         bool                     `yaml:"isOpaque"`
	ClientCredential *ClientCredentialSetting `yaml:"clientCredential"`
	Basic            *BasicSetting            `yaml:"basic"`
}

type TokenCredentialType string

const (
	TokenCredentialClientCredential TokenCredentialType = "client_credentials"
	TokenCredentialBasicCredential  TokenCredentialType = "basic"
)

func (setting TokenCredentialSetting) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if len(setting.Id) == 0 {
		result.AddError("tokenCredentials.id is required")
	} else {
		if strings.ContainsAny(setting.Id, " ") {
			result.AddError("tokenCredentials.id can't contains space")
		}
	}

	if len(setting.AuthEndpoint) == 0 {
		result.AddError("tokenCredentials.authEndpoint is required")
	}

	if setting.Type == TokenCredentialClientCredential {

		if setting.ClientCredential == nil {
			result.AddError("tokenCredentials.ClientCredential is required")
		} else {
			result.AppendValidable(setting.ClientCredential)
		}
	}

	if setting.Type == TokenCredentialBasicCredential {

		if setting.Basic == nil {
			result.AddError("tokenCredentials.Basic is required")
		} else {
			result.AppendValidable(setting.Basic)
		}
	}

	return result
}
