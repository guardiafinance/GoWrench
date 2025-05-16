package nats_settings

type NatsSettings struct {
	ConnectionId string `yaml:"connectionId"`
	IsStream     bool   `yaml:"isStream"`
	SubjectName  string `yaml:"subjectName"`
}
