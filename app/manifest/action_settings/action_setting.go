package action_settings

type ActionSetting struct {
	Id   string                  `yaml:"id"`
	Type string                  `yaml:"type"`
	Mock HttpResponseMockSetting `yaml:"mock"`
}
