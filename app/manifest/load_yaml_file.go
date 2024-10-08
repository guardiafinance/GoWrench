package manifest

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadYamlFile() (*ApplicationSetting, error) {
	file, err := os.Open("configApp.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var yamlMap ApplicationSetting
	err = yaml.Unmarshal(data, &yamlMap)
	if err != nil {
		return nil, err
	}

	return &yamlMap, nil
}
