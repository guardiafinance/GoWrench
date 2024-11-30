package handlers

import (
	settings "wrench/app/manifest/application_settings"
)

type Chain struct {
	MapHandle map[string]Handler
}

func (chain *Chain) BuildChain(settings *settings.ApplicationSettings) Handler {
	var firstHandler = new(HttpFirstHandler)

	// for _, endpoint := range settings.Api.Endpoints {
	// 	if endpoint.
	// }

	return firstHandler
}
