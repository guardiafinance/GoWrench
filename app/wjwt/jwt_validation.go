package wjwt

import (
	"context"
	"log"
	"wrench/app/manifest/api_settings"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
)

var jwksData *keyfunc.JWKS

func JwksValidationAuthentication(tokenString string, authorizationSettings *api_settings.AuthorizationSettings) bool {
	LoadCertificates(authorizationSettings.JwksUrl)

	token, err := jwt.Parse(tokenString, jwksData.Keyfunc)
	if err != nil {
		log.Printf("Failed to parse the JWT.\nError: %s", err.Error())
		return false
	}

	// Check if the token is valid.
	if !token.Valid {
		log.Println("The token is not valid.")
	}
	log.Println("The token is valid.")
	return token.Valid
}

func LoadCertificates(jwksUrl string) {

	if jwksData == nil {
		ctx := context.Background()

		options := keyfunc.Options{
			Ctx: ctx,
			RefreshErrorHandler: func(err error) {
				log.Printf("There was an error with the jwt.Keyfunc\nError: %s", err.Error())
			},
		}
		jwks, err := keyfunc.Get(jwksUrl, options)
		if err != nil {
			log.Fatalf("Failed to create JWKS from resource at the given URL.\nError: %s", err.Error())
		}

		jwksData = jwks
	}
}
