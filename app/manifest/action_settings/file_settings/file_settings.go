package file_settings

import "wrench/app/manifest/validation"

type FileSettings struct {
	Path     string                `yaml:"path"`
	Response *FileResponseSettings `yaml:"response"`
}

func (setting FileSettings) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if len(setting.Path) == 0 {
		result.AddError("actions.file.path is required")
	}

	if setting.Response != nil {
		result.AppendValidable(setting.Response)
	}

	return result
}
