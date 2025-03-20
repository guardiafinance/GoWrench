package file_settings

import "wrench/app/manifest/validation"

type FileResponseSettings struct {
	ContentType string            `default:"application/json" yaml:"contentType"`
	Headers     map[string]string `yaml:"headers"`
	StatusCode  int               `default:"200" yaml:"statusCode"`
}

func (setting FileResponseSettings) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	return result
}
