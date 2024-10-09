package manifest

import (
	"io"
	"os"

	"wrench/app/manifest/application_settings"

	"gopkg.in/yaml.v3"
)

func LoadYamlFile() (*application_settings.ApplicationSetting, error) {
	file, err := os.Open("configApp.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var yamlMap application_settings.ApplicationSetting
	err = yaml.Unmarshal(data, &yamlMap)
	if err != nil {
		return nil, err
	}

	return &yamlMap, nil
}
