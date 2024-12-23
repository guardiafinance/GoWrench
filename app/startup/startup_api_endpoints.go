package startup

import (
	"net/http"
	"os"
	"strings"
	handler "wrench/app/handlers"
	"wrench/app/manifest/application_settings"

	"github.com/gorilla/mux"
)

func LoadApiEndpoint() *mux.Router {
	app := application_settings.ApplicationSettingsStatic
	endpoints := app.Api.Endpoints
	if len(endpoints) == 0 {
		return nil
	}

	r := mux.NewRouter()
	initialPage := new(InitialPage)
	initialPage.Append("<h2>Service: " + app.Service.Name + "version: " + app.Service.Version + "</h2>")

	initialPage.Append("<h2>Endpoints</h2>")
	for _, endpoint := range endpoints {
		method := strings.ToUpper(string(endpoint.Method))
		route := endpoint.Route

		var delegate = new(handler.RequestDelegate)
		delegate.SetEndpoint(&endpoint)
		r.HandleFunc(route, delegate.FirstHttp).Methods(method)
		initialPage.Append("Route: <i>" + route + "</i> Method: <i>" + method + "</i>")
	}
	initialPage.Append("</br></br>")

	initialPage.Append("<h2>Envs</h2>")
	for _, env := range os.Environ() {
		envSplitted := strings.Split(env, "=")
		envName := envSplitted[0]
		initialPage.Append("Env: <i>" + envName + "</i>")
	}

	r.HandleFunc("/", initialPage.WriteInitialPage).Methods("GET")
	r.HandleFunc("/hc", initialPage.HealthCheckEndpoint).Methods("GET")
	return r
}

func LoadEndpointHc(w http.ResponseWriter, r *http.Request) {
	app := application_settings.ApplicationSettingsStatic
	validationResult := app.Valid()

	if len(validationResult.GetErrors()) > 0 {

	}
}
