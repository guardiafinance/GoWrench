package contexts

import (
	"strings"
)

const prefixWrenchContextRequestHeaders = "wrenchContext.request.headers."
const prefixBodyContext = "bodyContext."

func IsCalculatedValue(value string) bool {
	return strings.HasPrefix(value, "{{") && strings.HasSuffix(value, "}}")
}

func ReplaceCalculatedValue(command string) string {
	return strings.ReplaceAll(strings.ReplaceAll(command, "{{", ""), "}}", "")
}

func IsWrenchContextCommand(command string) bool {
	return strings.HasPrefix(command, prefixWrenchContextRequestHeaders)
}

func IsBodyContextCommand(command string) bool {
	return strings.HasPrefix(command, prefixBodyContext)
}

func GetValueWrenchContext(command string, wrenchContext *WrenchContext) string {

	if IsCalculatedValue(command) {
		command = ReplaceCalculatedValue(command)
	}

	if strings.HasPrefix(command, prefixWrenchContextRequestHeaders) {
		headerName := strings.ReplaceAll(command, prefixWrenchContextRequestHeaders, "")
		return wrenchContext.Request.Header.Get(headerName)
	}

	return ""
}

func ReplacePrefixBodyContext(command string) string {
	if strings.HasPrefix(command, prefixBodyContext) {
		command = strings.ReplaceAll(command, prefixBodyContext, "")
	}
	return command
}

func GetValueBodyContext(command string, bodyContext *BodyContext) string {

	if IsCalculatedValue(command) {
		command = ReplaceCalculatedValue(command)
	}

	if strings.HasPrefix(command, prefixBodyContext) {
		propertyName := strings.ReplaceAll(command, prefixBodyContext, "")
		jsonMap := bodyContext.ParseBodyToMapObject()
		return getBodyValue(jsonMap, propertyName)
	}

	return ""
}

func getBodyValue(jsonMap map[string]interface{}, propertyName string) string {
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
			break
		}
	}
	return value
}
