package manifest

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Endpoints struct {
	Route    string `yaml:"route"`
	Method   string `yaml:"method"`
	ActionID string `yaml:"actionId"`
}

type API struct {
	Endpoints []Endpoints `yaml:"endpoints"`
}

type HttpClient struct {
	Id string `yaml:"id"`
	BaseUrl string `yaml:"baseUrl"`
}


type Action struct {
	Id string `yaml:"id"`
	Type string `yaml:"type"`
	Path string `yaml:"path"`
	HttpClientId string `yaml:"httpClientId"`
	Mock struct {
		Value string `yaml:"value"`
		ContentType string `yaml:"contentType"`
		Method string `yaml:"method"`
	} `yaml:"mock"`
}

type Service struct {
	Name string `yaml:"name"`
	Version string `yaml:"version"`
}

type Config struct {
	Api           API           `yaml:"api"`
	Actions       []Action      `yaml:"actions"`
	Service 			Service 			`yaml:"service"`
	HttpClients   []HttpClient    `yaml:"httpClients"`
}

func LoadYamlFile() (*Config, error) {
	file, err := os.Open("configApp.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var yamlMap Config
	err = yaml.Unmarshal(data, &yamlMap)
	if err != nil {
		return nil, err
	}

	return &yamlMap, nil
}
