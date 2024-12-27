package startup

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"wrench/app/manifest/application_settings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func LoadAwsSecrets(fileConfig []byte) {

	applicationSetting, err := application_settings.ParseToApplicationSetting(fileConfig)
	if err != nil {
		return
	}

	setting := applicationSetting.Aws
	if setting != nil &&
		setting.AwsSecretSettings != nil &&
		len(setting.AwsSecretSettings.SecretsName) > 0 {
		for _, secretName := range setting.AwsSecretSettings.SecretsName {
			log.Print("Loading secret " + secretName)
			secretValue := getSecretValue(setting.Region, secretName)
			if secretValue == "" {
				continue
			}
			secretMap, err := parseSecretToMap(secretValue)
			if err == nil {
				setMapToEnv(secretName, secretMap)
			}
			log.Print("Loaded secret " + secretName)
		}
	}
}

func getSecretValue(region string, secretName string) string {
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Print(err)
	}

	svc := secretsmanager.NewFromConfig(config)
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Print(err.Error())
		return ""
	}

	var secretString string = *result.SecretString
	return secretString
}

func parseSecretToMap(secretValue string) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	jsonErr := json.Unmarshal([]byte(secretValue), &jsonMap)

	if jsonErr != nil {
		return nil, jsonErr
	}

	return jsonMap, nil
}

func setMapToEnv(secretName string, jsonMap map[string]interface{}) {
	for k, v := range jsonMap {
		jsonMapOther, ok := v.(map[string]interface{})
		name := secretName + "__" + k

		if ok {
			setMapToEnv(name, jsonMapOther)
		}

		value, ok := v.(string)
		if ok {
			os.Setenv(name, value)
		}
	}
}
