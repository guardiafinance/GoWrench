package startup

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"
	client "wrench/app/clients/http"
	"wrench/app/manifest/application_settings"
	"wrench/app/manifest/idp"
)

var idpTokens map[string]*JwtData

func GetIdpToken(idpId string) *JwtData {
	if idpTokens == nil {
		return nil
	}

	return idpTokens[idpId]
}

func LoadIdpAuthentication() {
	app_settings := application_settings.ApplicationSettingsStatic

	if len(app_settings.IDP) > 0 {
		if idpTokens == nil {
			idpTokens = make(map[string]*JwtData)
		}

		for {
			for _, idpSetting := range app_settings.IDP {

				jwtData := GetIdpToken(idpSetting.Id)

				if jwtData != nil {
					if jwtData.IsExpired(5) == false {
						continue
					}
				}

				jwtData, err := authenticateClientCredentials(idpSetting)
				if err != nil {
					continue
					// TODO setting error to unhealthy api
				}
				jwtData.LoadJwtPayload()
				idpTokens[idpSetting.Id] = jwtData
			}

			time.Sleep(2 * time.Minute)
		}
	}
}

func authenticateClientCredentials(idpSetting *idp.IdpAuthenticationSetting) (*JwtData, error) {
	request := new(client.HttpClientRequestData)
	data := url.Values{}
	data.Set("client_id", idpSetting.ClientId)
	data.Set("client_secret", idpSetting.ClientSecret)
	data.Set("grant_type", "client_credentials")

	request.Body = []byte(data.Encode())
	request.Method = "POST"
	request.Url = idpSetting.AuthEndpoint

	request.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	ctx := context.Background()

	response, err := client.HttpClientDo(ctx, request)

	if err != nil {
		// TODO add error flow
	}

	if response.StatusCodeSuccess() {
		token := string(response.Body)
		tokenArray := []byte(token)
		jwtData := new(JwtData)

		jsonErr := json.Unmarshal(tokenArray, &jwtData)
		if jsonErr != nil {
			return nil, jsonErr
		}

		return jwtData, nil
	}

	return nil, errors.New(fmt.Sprintf("Can't get jwtToken response_status_code: %v", response.StatusCode))
}
