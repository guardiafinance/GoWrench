package api_settings

type EndpointSetting struct {
	Route        string `yaml:"route"`
	Method       string `yaml:"method"`
	ActionID     string `yaml:"actionId"`
	FlowActionID string `yaml:"flowActionId"`
}
