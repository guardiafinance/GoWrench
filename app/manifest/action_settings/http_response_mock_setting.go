package action_settings

type HttpResponseMockSetting struct {
	Value       string `yaml:"value"`
	ContentType string `yaml:"contentType"`
	Method      string `yaml:"method"`
}
