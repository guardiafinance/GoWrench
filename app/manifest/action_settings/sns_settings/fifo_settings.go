package sns_settings

type FifoSettings struct {
	GroupId         string `yaml:"groupId"`
	DeduplicationId string `yaml:"deduplicationId"`
}
