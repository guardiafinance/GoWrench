package json_map

import (
	"encoding/json"
	"errors"
	"fmt"
)

func GetValue(jsonValue []byte, propertyName string) (string, error) {
	value := ""

	var jsonMap map[string]interface{}
	jsonErr := json.Unmarshal(jsonValue, &jsonMap)

	if jsonErr != nil {
		return "", jsonErr
	}

	value, ok := jsonMap[propertyName].(string)
	if !ok {
		return value, errors.New(fmt.Sprintf("Property %s not found", propertyName))
	}

	return value, nil
}
