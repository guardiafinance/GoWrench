package wjwt

import (
	"context"
	"log"
	client "wrench/app/clients/http"
	"wrench/app/manifest/api_settings"

	"github.com/lestrrat-go/jwx/jwk"
)

var jwksData jwk.Set

func JwksValidation(tokenString string, authorizationSettings *api_settings.AuthorizationSettings) {
	LoadCertificates(authorizationSettings.JwksUrl)

	_, err = jws.VerifyWithJWKSet([]byte(tokenString), jwksData, nil)
}

func LoadCertificates(jwksUrl string) {
	if jwksData == nil {

		ctx := context.Background()
		request := new(client.HttpClientRequestData)
		request.Method = "GET"
		request.Url = jwksUrl
		response, err := client.HttpClientDo(ctx, request)
		if err != nil {
			log.Print(err)
		}

		jwks, err := jwk.Parse(response.Body)
		if err != nil {
			panic(err)
		}

		jwksData = jwks
	}
}

// func (jwks *JwksData) GetCertificateByKid(kid string) *CertificateData {
// 	if len(jwks.Certificates) == 0 {
// 		return nil
// 	}

// 	var certificate *CertificateData

// 	for _, cert := range jwks.Certificates {
// 		if cert.Kid == kid {
// 			certificate = cert
// 			break
// 		}
// 	}

// 	return certificate
// }
