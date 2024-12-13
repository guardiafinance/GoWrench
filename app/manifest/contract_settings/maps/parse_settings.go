package maps

import (
	"wrench/app/manifest/validation"
)

type ParseSettings struct {
	WhenEquals []string `yaml:"whenEquals"`
}

func (setting ParseSettings) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	return result
}
