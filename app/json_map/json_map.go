package json_map

import (
	"strings"
	"time"
	"wrench/app/contexts"
	"wrench/app/manifest/contract_settings/maps"

	"github.com/google/uuid"
)

func GetValue(jsonMap map[string]interface{}, propertyName string, deleteProperty bool) (string, map[string]interface{}) {
	value := ""

	var jsonMapCurrent map[string]interface{}
	jsonMapCurrent = jsonMap
	propertyNameSplitted := strings.Split(propertyName, ".")

	for _, property := range propertyNameSplitted {
		valueTemp, ok := jsonMapCurrent[property].(map[string]interface{})
		if ok {
			jsonMapCurrent = valueTemp
			continue
		}

		valueTempString, ok := jsonMapCurrent[property].(string)
		if ok {
			value = valueTempString

			if deleteProperty {
				delete(jsonMapCurrent, property)
			}

			break
		}
	}
	return value, jsonMap
}

func SetValue(jsonMap map[string]interface{}, propertyName string, newValue string) map[string]interface{} {
	var jsonMapCurrent map[string]interface{}
	jsonMapCurrent = jsonMap
	propertyNameSplitted := strings.Split(propertyName, ".")
	total := len(propertyNameSplitted)

	for i, property := range propertyNameSplitted {
		valueTemp, ok := jsonMapCurrent[property].(map[string]interface{})
		if ok {
			jsonMapCurrent = valueTemp
			continue
		}

		if i+1 == total {
			jsonMapCurrent[property] = newValue
		}
	}

	return jsonMap
}

func CreateProperty(jsonMap map[string]interface{}, propertyName string, value string) map[string]interface{} {

	var jsonMapCurrent map[string]interface{}
	jsonMapCurrent = jsonMap
	propertyNameSplitted := strings.Split(propertyName, ".")
	total := len(propertyNameSplitted)

	for i, property := range propertyNameSplitted {
		valueTemp, ok := jsonMapCurrent[property].(map[string]interface{})
		if ok {
			jsonMapCurrent = valueTemp
		} else {
			if i+1 < total {
				jsonMapNew := make(map[string]interface{})
				jsonMapCurrent[property] = jsonMapNew
				jsonMapCurrent = jsonMapNew
			}
		}

		if i+1 == total {
			jsonMapCurrent[property] = value
		}
	}
	return jsonMap
}

func RenameProperties(jsonMap map[string]interface{}, properties []string) map[string]interface{} {
	jsonValueCurrent := jsonMap
	for _, property := range properties {
		propertyNameSplitted := strings.Split(property, ":")
		propertyNameOld := propertyNameSplitted[0]
		propertyNameNew := propertyNameSplitted[1]
		jsonValueCurrent = RenameProperty(jsonValueCurrent, propertyNameOld, propertyNameNew)
	}
	return jsonValueCurrent
}

func DuplicatePropertiesValue(jsonMap map[string]interface{}, properties []string) map[string]interface{} {
	jsonValueCurrent := jsonMap
	for _, property := range properties {
		propertyNameSplitted := strings.Split(property, ":")
		propertyNameSource := propertyNameSplitted[0]
		propertyNameDestination := propertyNameSplitted[1]
		jsonValueCurrent = DuplicatePropertyValue(jsonValueCurrent, propertyNameSource, propertyNameDestination)
	}
	return jsonValueCurrent
}

func DuplicatePropertyValue(jsonMap map[string]interface{}, propertyNameSource string, propertyNameDestination string) map[string]interface{} {
	value, jsonValue := GetValue(jsonMap, propertyNameSource, false)
	return CreateProperty(jsonValue, propertyNameDestination, value)
}

func RenameProperty(jsonMap map[string]interface{}, propertyNameOld string, propertyNameNew string) map[string]interface{} {
	value, jsonValue := GetValue(jsonMap, propertyNameOld, true)
	return CreateProperty(jsonValue, propertyNameNew, value)
}

func RemoveProperties(jsonMap map[string]interface{}, propertiesName []string) map[string]interface{} {
	if propertiesName == nil {
		return nil
	}

	currentJsonValue := jsonMap
	for _, property := range propertiesName {
		currentJsonValue = RemoveProperty(currentJsonValue, property)
	}

	return currentJsonValue
}

func RemoveProperty(jsonMap map[string]interface{}, propertyName string) map[string]interface{} {
	var jsonMapCurrent map[string]interface{}
	jsonMapCurrent = jsonMap

	propertyNameSplitted := strings.Split(propertyName, ".")
	total := len(propertyNameSplitted)

	for i, property := range propertyNameSplitted {
		valueTemp, ok := jsonMapCurrent[property].(map[string]interface{})

		if ok && i+1 == total {
			delete(jsonMapCurrent, property)
			break
		} else {
			jsonMapCurrent = valueTemp
		}

		if i+1 == total {
			delete(jsonMapCurrent, property)
		}
	}

	return jsonMap
}

func CreatePropertiesInterpolationValue(jsonMap map[string]interface{}, propertiesValues []string, wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) map[string]interface{} {
	jsonValueCurrent := jsonMap
	for _, propertyValue := range propertiesValues {
		propertyValueSplitted := strings.Split(propertyValue, ":")
		propertyName := propertyValueSplitted[0]
		valueArray := propertyValueSplitted[1:]
		value := strings.Join(valueArray, ":")
		jsonValueCurrent = CreatePropertyInterpolationValue(jsonValueCurrent, propertyName, value, wrenchContext, bodyContext)
	}
	return jsonValueCurrent
}

func CreatePropertyInterpolationValue(jsonMap map[string]interface{}, propertyName string, value string, wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) map[string]interface{} {
	valueResult := value

	if contexts.IsCalculatedValue(value) {
		rawValue := contexts.ReplaceCalculatedValue(value)

		if rawValue == "uuid" {
			valueResult = uuid.New().String()
		} else if strings.HasPrefix(rawValue, "time") {
			timeFormat := strings.ReplaceAll(rawValue, "time ", "")
			timeNow := time.Now()

			if len(timeFormat) > 0 {
				valueResult = timeNow.Format(timeFormat)
			} else {
				valueResult = timeNow.String()
			}
		} else if strings.HasPrefix(rawValue, "wrenchContext") {
			valueResult = contexts.GetValueWrenchContext(rawValue, wrenchContext)
		}

	}

	return CreateProperty(jsonMap, propertyName, valueResult)
}

func ParseValues(jsonMap map[string]interface{}, parse *maps.ParseSettings) map[string]interface{} {
	jsonValueCurrent := jsonMap
	if parse.WhenEquals != nil {
		for _, whenEqual := range parse.WhenEquals {
			if contexts.IsCalculatedValue(whenEqual) {
				whenEqual = contexts.ReplacePrefixBodyContext(whenEqual)
				rawWhenEqual := contexts.ReplaceCalculatedValue(whenEqual)

				whenEqualSplitted := strings.Split(rawWhenEqual, ":")
				propertyNameWithEqualValue := whenEqualSplitted[0]
				propertyNameWithEqualValueSplitted := strings.Split(propertyNameWithEqualValue, ".")

				lenWithEqual := len(propertyNameWithEqualValueSplitted)

				valueArray := propertyNameWithEqualValueSplitted[:lenWithEqual-1]

				propertyName := strings.Join(valueArray, ".")
				equalValue := propertyNameWithEqualValueSplitted[lenWithEqual-1] // value to compare

				parseToValue := whenEqualSplitted[1] // value if equals should be used

				valueCurrent, _ := GetValue(jsonMap, propertyName, false)

				if valueCurrent == equalValue {
					jsonValueCurrent = SetValue(jsonValueCurrent, propertyName, parseToValue)
				}
			}
		}
	}

	return jsonValueCurrent
}
