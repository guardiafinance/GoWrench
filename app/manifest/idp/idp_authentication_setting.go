package idp

import (
	"fmt"
	"strings"
	"wrench/app/manifest/validation"
)

type IdpAuthenticationSetting struct {
	Id           string                `yaml:"id"`
	Type         IdpAuthenticationType `yaml:"type"`
	AuthEndpoint string                `yaml:"authEndpoint"`
	ClientId     string                `yaml:"clientId"`
	ClientSecret string                `yaml:"clientSecret"`
}

type IdpAuthenticationType string

const (
	IdpAuthenticationClientCredential IdpAuthenticationType = "clientCredentials"
)

func (setting IdpAuthenticationSetting) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if len(setting.Id) == 0 {
		result.AddError("idpAuthentication.id is required")
	} else {
		if strings.ContainsAny(setting.Id, " ") {
			result.AddError("idpAuthentication.id can't contains space")
		}
	}

	if len(setting.AuthEndpoint) == 0 {
		result.AddError("idpAuthentication.authEndpoint is required")
	}

	if len(setting.ClientId) == 0 {
		result.AddError("idpAuthentication.clientId is required")
	}

	if len(setting.ClientSecret) == 0 {
		result.AddError("idpAuthentication.clientSecret is required")
	}

	if setting.Type != IdpAuthenticationClientCredential {
		result.AddError(fmt.Sprintf("idpAuthentication.type should be %s", IdpAuthenticationClientCredential))
	}

	return result
}
