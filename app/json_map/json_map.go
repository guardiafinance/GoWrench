package json_map

import (
	"encoding/json"
	"strings"
)

func GetValue(jsonValue []byte, propertyName string) (string, error) {
	value := ""

	var jsonMapCurrent map[string]interface{}
	var jsonMap map[string]interface{}
	jsonErr := json.Unmarshal(jsonValue, &jsonMap)

	if jsonErr != nil {
		return "", jsonErr
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
			break
		}
	}

	return value, nil
}
