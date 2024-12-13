package json_map

import (
	"encoding/json"
	"strings"
	"time"
	"wrench/app/contexts"
	"wrench/app/manifest/contract_settings/maps"

	"github.com/google/uuid"
)

func GetValue(jsonValue []byte, propertyName string, deleteProperty bool) (string, []byte) {
	value := ""

	var jsonMapCurrent map[string]interface{}
	var jsonMap map[string]interface{}
	jsonErr := json.Unmarshal(jsonValue, &jsonMap)

	if jsonErr != nil {
		return "", jsonValue
	}

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

	jsonArray, _ := json.Marshal(jsonMap)
	return value, jsonArray
}

func SetValue(jsonValue []byte, propertyName string, newValue string) []byte {

	var jsonMapCurrent map[string]interface{}
	var jsonMap map[string]interface{}
	jsonErr := json.Unmarshal(jsonValue, &jsonMap)

	if jsonErr != nil {
		return jsonValue
	}

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

	jsonArray, _ := json.Marshal(jsonMap)
	return jsonArray
}

func CreateProperty(jsonValue []byte, propertyName string, value string) []byte {

	var jsonMapCurrent map[string]interface{}
	var jsonMap map[string]interface{}
	json.Unmarshal(jsonValue, &jsonMap)

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

	jsonArray, _ := json.Marshal(jsonMap)
	return jsonArray
}

func RenameProperties(jsonValue []byte, properties []string) []byte {
	jsonValueCurrent := jsonValue
	for _, property := range properties {
		propertyNameSplitted := strings.Split(property, ":")
		propertyNameOld := propertyNameSplitted[0]
		propertyNameNew := propertyNameSplitted[1]
		jsonValueCurrent = RenameProperty(jsonValueCurrent, propertyNameOld, propertyNameNew)
	}
	return jsonValueCurrent
}

func DuplicatePropertiesValue(jsonValue []byte, properties []string) []byte {
	jsonValueCurrent := jsonValue
	for _, property := range properties {
		propertyNameSplitted := strings.Split(property, ":")
		propertyNameSource := propertyNameSplitted[0]
		propertyNameDestination := propertyNameSplitted[1]
		jsonValueCurrent = DuplicatePropertyValue(jsonValueCurrent, propertyNameSource, propertyNameDestination)
	}
	return jsonValueCurrent
}

func DuplicatePropertyValue(jsonValue []byte, propertyNameSource string, propertyNameDestination string) []byte {
	value, jsonValue := GetValue(jsonValue, propertyNameSource, false)
	return CreateProperty(jsonValue, propertyNameDestination, value)
}

func RenameProperty(jsonValue []byte, propertyNameOld string, propertyNameNew string) []byte {
	value, jsonValue := GetValue(jsonValue, propertyNameOld, true)
	return CreateProperty(jsonValue, propertyNameNew, value)
}

func RemoveProperties(jsonValue []byte, propertiesName []string) []byte {
	if propertiesName == nil {
		return nil
	}

	currentJsonValue := jsonValue
	for _, property := range propertiesName {
		currentJsonValue = RemoveProperty(currentJsonValue, property)
	}

	return currentJsonValue
}

func RemoveProperty(jsonValue []byte, propertyName string) []byte {
	var jsonMapCurrent map[string]interface{}
	var jsonMap map[string]interface{}
	json.Unmarshal(jsonValue, &jsonMap)

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

	jsonArray, _ := json.Marshal(jsonMap)
	return jsonArray
}

func CreatePropertiesInterpolationValue(jsonValue []byte, propertiesValues []string, wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) []byte {
	jsonValueCurrent := jsonValue
	for _, propertyValue := range propertiesValues {
		propertyValueSplitted := strings.Split(propertyValue, ":")
		propertyName := propertyValueSplitted[0]
		valueArray := propertyValueSplitted[1:]
		value := strings.Join(valueArray, ":")
		jsonValueCurrent = CreatePropertyInterpolationValue(jsonValueCurrent, propertyName, value, wrenchContext, bodyContext)
	}
	return jsonValueCurrent
}

func calculatedValue(value string) bool {
	return strings.HasPrefix(value, "{{") && strings.HasSuffix(value, "}}")
}

func replaceCalculatedValue(command string) string {
	return strings.ReplaceAll(strings.ReplaceAll(command, "{{", ""), "}}", "")
}

func CreatePropertyInterpolationValue(jsonValue []byte, propertyName string, value string, wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) []byte {
	valueResult := value

	if calculatedValue(value) {
		rawValue := replaceCalculatedValue(value)

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
			valueResult = getValueWrenchContext(rawValue, wrenchContext)
		}

	}

	return CreateProperty(jsonValue, propertyName, valueResult)
}

const wrenchContextRequestHeaders = "wrenchContext.request.headers."

func getValueWrenchContext(command string, wrenchContext *contexts.WrenchContext) string {

	if strings.HasPrefix(command, wrenchContextRequestHeaders) {
		headerName := strings.ReplaceAll(command, wrenchContextRequestHeaders, "")
		return wrenchContext.Request.Header.Get(headerName)
	}

	return ""
}

const bodyContext = "bodyContext."

func ParseValues(jsonValue []byte, parse *maps.ParseSettings) []byte {
	jsonValueCurrent := jsonValue
	if parse.WhenEquals != nil {
		for _, whenEqual := range parse.WhenEquals {
			if calculatedValue(whenEqual) {
				whenEqual = strings.ReplaceAll(whenEqual, bodyContext, "")
				rawWhenEqual := replaceCalculatedValue(whenEqual)

				whenEqualSplitted := strings.Split(rawWhenEqual, ":")
				propertyNameWithEqualValue := whenEqualSplitted[0]
				propertyNameWithEqualValueSplitted := strings.Split(propertyNameWithEqualValue, ".")

				lenWithEqual := len(propertyNameWithEqualValueSplitted)

				valueArray := propertyNameWithEqualValueSplitted[:lenWithEqual-1]

				propertyName := strings.Join(valueArray, ".")
				equalValue := propertyNameWithEqualValueSplitted[lenWithEqual-1] // value to compare

				parseToValue := whenEqualSplitted[1] // value if equals should be used

				valueCurrent, _ := GetValue(jsonValue, propertyName, false)

				if valueCurrent == equalValue {
					jsonValueCurrent = SetValue(jsonValueCurrent, propertyName, parseToValue)
				}
			}
		}
	}

	return jsonValueCurrent
}
