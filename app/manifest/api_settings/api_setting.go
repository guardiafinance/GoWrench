package api_settings

type ApiSetting struct {
	Endpoints []EndpointSetting `yaml:"endpoints"`
}
