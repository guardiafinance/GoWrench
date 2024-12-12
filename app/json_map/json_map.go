package json_map

import (
	"encoding/json"
	"strings"
)

func GetValue(jsonValue []byte, propertyName string, deleteProperty bool) string {
	value := ""

	var jsonMapCurrent map[string]interface{}
	var jsonMap map[string]interface{}
	jsonErr := json.Unmarshal(jsonValue, &jsonMap)

	if jsonErr != nil {
		return ""
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

	return value
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
		}

		if i+1 == total {
			jsonMapCurrent[property] = value
		}
	}

	jsonArray, _ := json.Marshal(jsonMapCurrent)
	return jsonArray
}

func RenameProperties(jsonValue []byte, properties []string) []byte {
	jsonValueCurrent := jsonValue
	for _, property := range properties {
		propertyNameSplitted := strings.Split(property, ":")
		propertyNameOld := propertyNameSplitted[0]
		propertyNameNew := propertyNameSplitted[1]
		jsonValueCurrent = RenameProperty(jsonValue, propertyNameOld, propertyNameNew)
	}
	return jsonValueCurrent
}

func ClonePropertyValue(jsonValue []byte, propertyNameSource string, propertyNameDestination string) []byte {
	value := GetValue(jsonValue, propertyNameSource, false)
	return CreateProperty(jsonValue, propertyNameDestination, value)
}

func RenameProperty(jsonValue []byte, propertyNameOld string, propertyNameNew string) []byte {
	value := GetValue(jsonValue, propertyNameOld, true)
	return CreateProperty(jsonValue, propertyNameNew, value)
}

func RemoveProperties(jsonValue []byte, propertiesName []string) []byte {
	if propertiesName == nil {
		return nil
	}

	currentJsonValue := jsonValue
	for _, property := range propertiesName {
		currentJsonValue = RemoveProperty(jsonValue, property)
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
		if ok {
			jsonMapCurrent = valueTemp
		}

		if i+1 == total {
			delete(jsonMapCurrent, property)
		}
	}

	jsonArray, _ := json.Marshal(jsonMapCurrent)
	return jsonArray
}
