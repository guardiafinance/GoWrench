package startup

import (
	"net/http"
	"os"
	"strings"

	handler "wrench/app/handlers"
	"wrench/app/manifest/api_settings"
	"wrench/app/manifest/application_settings"
	"wrench/app/wjwt"

	"github.com/gorilla/mux"
)

func LoadApiEndpoint() *mux.Router {
	app := application_settings.ApplicationSettingsStatic

	if app.Api == nil || app.Api.Endpoints == nil {
		return nil
	}

	hasAuthorization := app.Api.HasAuthorization()
	endpoints := app.Api.Endpoints
	r := mux.NewRouter()
	initialPage := new(InitialPage)
	initialPage.Append("<h2>Service: " + app.Service.Name + "version: " + app.Service.Version + "</h2>")
	initialPage.Append("<h2>Endpoints</h2>")

	for _, endpoint := range endpoints {
		method := strings.ToUpper(string(endpoint.Method))
		route := endpoint.Route

		var delegate = new(handler.RequestDelegate)
		delegate.SetEndpoint(&endpoint)
		shouldConfigureAuthorization := endpoint.ShouldConfigureAuthorization(hasAuthorization)

		if shouldConfigureAuthorization {
			r.Handle(route, authMiddleware(app.Api.Authorization, http.HandlerFunc(delegate.HttpHandler))).Methods(method)
		} else {
			r.HandleFunc(route, delegate.HttpHandler).Methods(method)
		}

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

func authMiddleware(authorizationSettings *api_settings.AuthorizationSettings, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		if authorizationSettings.Type == api_settings.JWKSAuthorizationType {
			tokenIsValid := wjwt.JwksValidationAuthentication(tokenString, authorizationSettings)
			if tokenIsValid == false {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorization"))
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
