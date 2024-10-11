package manifest

import (
	"io"
	"os"

	"wrench/app/manifest/application_settings"

	"gopkg.in/yaml.v3"
)

func LoadYamlFile(pathFile string) (*application_settings.ApplicationSettings, error) {
	file, err := os.Open(pathFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	applicationSettings := new(application_settings.ApplicationSettings)
	err = yaml.Unmarshal(data, applicationSettings)
	if err != nil {
		return nil, err
	}

	return applicationSettings, nil
}
