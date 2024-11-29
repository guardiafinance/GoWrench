package startup

import (
	"net/http"
	"strings"
	handler "wrench/app/handlers"
	settings "wrench/app/manifest/api_settings"

	"github.com/gorilla/mux"
)

func LoadApiEndpoint(endpoints []settings.EndpointSettings) {
	r := mux.NewRouter()
	initialPage := new(InitialPage)

	for _, endpoint := range endpoints {
		method := strings.ToUpper(string(endpoint.Method))
		route := endpoint.Route
		if route[0] != '/' {
			route = "/" + route
		}

		r.HandleFunc(route, handler.FirstHttp).Methods(method)
		initialPage.Append("Route: <i>" + route + "</i> Method: <i>" + method + "</i>")
	}

	r.HandleFunc("/", initialPage.WriteInitialPage).Methods("GET")

	http.ListenAndServe(":8085", r)
}
