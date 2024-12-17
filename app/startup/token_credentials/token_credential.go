package token_credentials

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"
	client "wrench/app/clients/http"
	"wrench/app/manifest/application_settings"
	credential "wrench/app/manifest/token_credential_settings"
)

var tokenCredentials map[string]*JwtData

func GetTokenCredentialById(tokenCredentialId string) *JwtData {
	if tokenCredentials == nil {
		return nil
	}

	return tokenCredentials[tokenCredentialId]
}

func LoadTokenCredentialAuthentication() {
	app_settings := application_settings.ApplicationSettingsStatic

	if len(app_settings.TokenCredentials) > 0 {
		if tokenCredentials == nil {
			tokenCredentials = make(map[string]*JwtData)
		}

		for {
			for _, setting := range app_settings.TokenCredentials {

				jwtData := GetTokenCredentialById(setting.Id)
				if jwtData != nil {
					if jwtData.IsExpired(5, setting.IsOpaque) == false {
						continue
					}
				}

				var err error
				if setting.Type == credential.TokenCredentialClientCredential {
					jwtData, err = authenticateClientCredentials(setting)
				} else {
					jwtData, err = basicCredentials(setting)
				}

				if err != nil {
					continue
					// TODO setting error to unhealthy api
				}

				if setting.IsOpaque == false {
					jwtData.LoadJwtPayload()
				}

				tokenCredentials[setting.Id] = jwtData
			}

			time.Sleep(2 * time.Minute)
		}
	}
}

func authenticateClientCredentials(setting *credential.TokenCredentialSetting) (*JwtData, error) {
	request := new(client.HttpClientRequestData)
	data := url.Values{}
	data.Set("client_id", setting.ClientCredential.ClientId)
	data.Set("client_secret", setting.ClientCredential.ClientSecret)
	data.Set("grant_type", "client_credentials")

	request.Body = []byte(data.Encode())
	request.Method = "POST"
	request.Url = setting.AuthEndpoint

	request.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	ctx := context.Background()

	response, err := client.HttpClientDo(ctx, request)

	if err != nil {
		// TODO add error flow
		return nil, err
	} else {

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
}

func basicCredentials(setting *credential.TokenCredentialSetting) (*JwtData, error) {
	request := new(client.HttpClientRequestData)
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	request.Body = []byte(data.Encode())
	request.Method = "POST"
	request.Url = setting.AuthEndpoint

	credential := fmt.Sprintf("%s:%s", setting.Basic.Username, setting.Basic.Password)
	credentialEncoded := base64.StdEncoding.EncodeToString([]byte(credential))

	request.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	request.SetHeader("Authorization", fmt.Sprintf("Basic %s", credentialEncoded))

	ctx := context.Background()

	response, err := client.HttpClientDo(ctx, request)

	if err != nil {
		// TODO add error flow
		return nil, err
	} else {

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
}
