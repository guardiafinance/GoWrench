package maps

import (
	"fmt"
	"slices"
	"strings"
	"wrench/app/manifest/validation"
)

var funcValids = []string{"rename", "new", "remove"}

type ContractMapSetting struct {
	Id       string   `yaml:"id"`
	Rename   []string `yaml:"rename"`
	Remove   []string `yaml:"remove"`
	Sequency []string `yaml:"sequency"`
	New      []string `yaml:"new"`
}

func (setting ContractMapSetting) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if len(setting.Id) <= 0 {
		result.AddError("contract.maps.id is required")
	}

	if len(setting.Rename) > 0 {
		errorSplitted := "contract.maps.rename should be configured looks like 'propertySource:propertyDestination' without space"
		for _, property := range setting.Rename {

			if strings.Contains(property, " ") {
				result.AddError(errorSplitted)
			}

			propertySplitted := strings.Split(property, ":")
			if len(propertySplitted) != 2 {
				result.AddError(errorSplitted)
			}
		}
	}

	if len(setting.Remove) > 0 {
		for _, remove := range setting.Remove {

			if strings.Contains(remove, " ") {
				result.AddError("contract.maps.remove can't contain space")
			}
		}
	}

	if len(setting.Sequency) > 0 {
		for _, s := range setting.Sequency {
			if slices.Contains(funcValids, s) == false {
				result.AddError(fmt.Sprintf("contract.maps.sequency should contain valid values. The value %s is not valid", s))
			}

			if s == "rename" && setting.Rename == nil {
				result.AddError("contract.maps.sequency rename not configured")
			} else if s == "new" && setting.New == nil {
				result.AddError("contract.maps.new rename not configured")
			} else if s == "remove" && setting.Remove == nil {
				result.AddError("contract.maps.sequency remove not configured")
			}
		}
	}

	return result
}
