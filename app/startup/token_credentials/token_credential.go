package token_credentials

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
	"wrench/app/auth"
	client "wrench/app/clients/http"
	"wrench/app/json_map"
	"wrench/app/manifest/application_settings"
	credential "wrench/app/manifest/token_credential_settings"
)

var tokenCredentials map[string]*auth.JwtData

func GetTokenCredentialById(tokenCredentialId string) *auth.JwtData {
	if tokenCredentials == nil {
		return nil
	}

	return tokenCredentials[tokenCredentialId]
}

func LoadTokenCredentialAuthentication() {
	app_settings := application_settings.ApplicationSettingsStatic

	if len(app_settings.TokenCredentials) > 0 {
		if tokenCredentials == nil {
			tokenCredentials = make(map[string]*auth.JwtData)
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
				} else if setting.Type == credential.TokenCredentialBasicCredential {
					jwtData, err = basicCredentials(setting)
				} else if setting.Type == credential.TokenCredentialCustomAuthentication {
					jwtData, err = customAuthentication(setting)
					continue
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

func authenticateClientCredentials(setting *credential.TokenCredentialSetting) (*auth.JwtData, error) {
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
			jwtData := new(auth.JwtData)

			jsonErr := json.Unmarshal(tokenArray, &jwtData)
			if jsonErr != nil {
				return nil, jsonErr
			}

			return jwtData, nil
		}

		return nil, fmt.Errorf("Can't get jwtToken response_status_code: %v", response.StatusCode)
	}
}

func basicCredentials(setting *credential.TokenCredentialSetting) (*auth.JwtData, error) {
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
			jwtData := new(auth.JwtData)

			jsonErr := json.Unmarshal(tokenArray, &jwtData)
			if jsonErr != nil {
				return nil, jsonErr
			}

			return jwtData, nil
		}

		return nil, fmt.Errorf("Can't get jwtToken response_status_code: %v", response.StatusCode)
	}
}

func customAuthentication(setting *credential.TokenCredentialSetting) (*auth.JwtData, error) {
	request := new(client.HttpClientRequestData)

	if len(setting.Custom.RequestBody) > 0 {
		request.Body = getBodyCustomAuthentication(setting.Custom.RequestBody)
	}
	request.Method = string(setting.Custom.Method)
	request.Url = setting.AuthEndpoint

	if len(setting.Custom.RequestHeaders) > 0 {
		for key, value := range setting.Custom.RequestHeaders {
			request.SetHeader(key, value)
		}
	}

	ctx := context.Background()

	response, err := client.HttpClientDo(ctx, request)

	if err != nil {
		// TODO add error flow
		return nil, err
	} else {

		if response.StatusCodeSuccess() {
			token := string(response.Body)
			tokenArray := []byte(token)
			jwtData := new(auth.JwtData)

			jsonErr := json.Unmarshal(tokenArray, &jwtData)
			if jsonErr != nil {
				return nil, jsonErr
			}

			return jwtData, nil
		}

		return nil, fmt.Errorf("Can't get jwtToken response_status_code: %v", response.StatusCode)
	}
}

func getBodyCustomAuthentication(customBody map[string]string) []byte {
	jsonMap := make(map[string]interface{})
	for key, value := range customBody {
		json_map.CreateProperty(jsonMap, key, value)
	}

	jsonArray, _ := json.Marshal(jsonMap)
	return jsonArray
}
