package startup

import (
	settings "wrench/app/manifest/application_settings"
)

func LoadApplicationSettings(settings *settings.ApplicationSettings) {
	var endpoints = settings.Api.Endpoints
	if len(endpoints) > 0 {
		LoadApiEndpoint(endpoints)
	}
}
