package auth

import (
	"context"
	"log"
	"sort"
	"strings"
	"wrench/app/manifest/api_settings"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
)

var jwksData *keyfunc.JWKS

func JwksValidationAuthorization(tokenString string, roles []string, scopes []string, claims []string) bool {
	var rolesValid, scopesValid, claimsValid bool = true, true, true

	tokenSplitted := strings.Split(tokenString, ".")
	tokenPayload := tokenSplitted[1]

	tokenPayloadMap := ConvertJwtPayloadBase64ToJwtPaylodData(tokenPayload)
	if tokenPayloadMap == nil {
		return false
	}

	if len(roles) > 0 {
		rolesValid = rolesValidation(tokenPayloadMap, roles)
	}

	if len(scopes) > 0 {
		scopesValid = scopesVadalition(tokenPayloadMap, scopes)
	}

	if len(claims) > 0 {
		claimsValid = claimsValidation(tokenPayloadMap, claims)
	}

	return rolesValid && scopesValid && claimsValid
}

func rolesValidation(tokenPayloadMap map[string]interface{}, roles []string) bool {
	rolesParsed, ok := tokenPayloadMap["roles"].([]interface{})
	result := false

	if ok {
		rolesParsedLen := len(rolesParsed)
		rolesParsedString := make([]string, rolesParsedLen)
		if rolesParsedLen > 0 {
			for index, roleParsed := range rolesParsed {
				roleParsedStringValue, ok := roleParsed.(string)
				if ok {
					rolesParsedString[index] = roleParsedStringValue
				}
			}
		}

		sort.Strings(rolesParsedString)
		sort.Strings(roles)

		rolesToken := strings.Join(rolesParsedString, " ")
		rolesRequired := strings.Join(roles, " ")

		if strings.HasPrefix(rolesToken, rolesRequired) {
			result = true
		} else {
			log.Printf("Roles %s is required", rolesRequired)
		}
	}
	return result
}

func scopesVadalition(tokenPayloadMap map[string]interface{}, scopes []string) bool {
	scopeParsed, ok := tokenPayloadMap["scope"].(string)
	result := true
	if ok {
		for _, scope := range scopes {

			if !strings.Contains(scopeParsed, scope) {
				result = false
				log.Printf("scope %s is required", scope)
				break
			}
		}
	} else {
		result = false
	}

	return result
}

func claimsValidation(tokenPayloadMap map[string]interface{}, claims []string) bool {
	result := true

	for _, claim := range claims {
		claimSplitted := strings.Split(claim, ":")
		claimName := claimSplitted[0]
		claimValue := claimSplitted[1]
		claimTokenValue, ok := tokenPayloadMap[claimName].(string)

		if !ok || claimTokenValue != claimValue {
			result = false
			log.Printf("claim %s with value %s is required", claimName, claimValue)
			break
		}
	}

	return result
}

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
