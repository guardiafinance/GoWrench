package startup

import (
	handlers "wrench/app/handlers"
	settings "wrench/app/manifest/application_settings"

	"github.com/gorilla/mux"
)

func LoadApplicationSettings(settings *settings.ApplicationSettings) *mux.Router {
	var chain = handlers.ChainStatic.GetStatic()
	chain.BuildChain(settings)
	return LoadApiEndpoint()
}
