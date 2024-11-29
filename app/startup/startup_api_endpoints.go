package startup

import (
	"net/http"
	handler "wrench/app/handlers"
	settings "wrench/app/manifest/api_settings"

	"github.com/gorilla/mux"
)

func LoadApiEndpoint(endpoints []settings.EndpointSettings) {
	r := mux.NewRouter()

	for _, endpoint := range endpoints {
		r.HandleFunc(endpoint.Route, handler.FirstHttp)
	}

	http.ListenAndServe(":8085", r)
}
