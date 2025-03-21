package startup

import (
	"net/http"
	"os"
	"strings"

	"wrench/app/auth"
	handler "wrench/app/handlers"
	"wrench/app/manifest/api_settings"
	"wrench/app/manifest/application_settings"

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
		shouldConfigureAuthorization := endpoint.ShouldConfigureAuthorization(hasAuthorization)
		var delegate = new(handler.RequestDelegate)
		delegate.SetEndpoint(&endpoint)

		if endpoint.IsProxy {
			method := strings.ToUpper(string(endpoint.Method))
			route := endpoint.Route

			if shouldConfigureAuthorization {
				r.Handle(route, authMiddleware(app.Api.Authorization, endpoint, http.HandlerFunc(delegate.HttpHandler))).Methods(method)
			} else {
				r.HandleFunc(route, delegate.HttpHandler).Methods(method)
			}
			initialPage.Append("Route: <i>" + route + "</i> Method: <i>" + method + "</i> <b>Not is proxy</b>")
		} else {
			for _, proxyRoute := range getProxyEndpoints() {
				route := proxyRoute
				for _, proxyMethod := range getProxyEndpointMethods() {
					method := proxyMethod

					if shouldConfigureAuthorization {
						r.Handle(route, authMiddleware(app.Api.Authorization, endpoint, http.HandlerFunc(delegate.HttpHandler))).Methods(method)
					} else {
						r.HandleFunc(route, delegate.HttpHandler).Methods(method)
					}
					initialPage.Append("Route: <i>" + route + "</i> Method: <i>" + method + "</i> <b> IS PROXY</b>")
				}
			}
		}
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

func authMiddleware(authorizationSettings *api_settings.AuthorizationSettings, endpoint api_settings.EndpointSettings, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		if authorizationSettings.Type == api_settings.JWKSAuthorizationType {
			tokenIsValid := auth.JwksValidationAuthentication(tokenString, authorizationSettings)
			if tokenIsValid {
				tokenIsAuthorized := auth.JwksValidationAuthorization(tokenString, endpoint.Roles, endpoint.Scopes, endpoint.Claims)
				if !tokenIsAuthorized {
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte("Forbidden"))
					return
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}

		}

		next.ServeHTTP(w, r)
	})
}

func getProxyEndpoints() [10]string {
	var endpoints [10]string
	endpoints[0] = "{r1}"
	endpoints[1] = "{r1}/{r2}"
	endpoints[2] = "{r1}/{r2}/{r3}"
	endpoints[3] = "{r1}/{r2}/{r3}/{r4}"
	endpoints[4] = "{r1}/{r2}/{r3}/{r4}/{r5}"
	endpoints[5] = "{r1}/{r2}/{r3}/{r4}/{r5}/{r6}"
	endpoints[6] = "{r1}/{r2}/{r3}/{r4}/{r5}/{r6}/{r7}"
	endpoints[7] = "{r1}/{r2}/{r3}/{r4}/{r5}/{r6}/{r7}/{r8}"
	endpoints[8] = "{r1}/{r2}/{r3}/{r4}/{r5}/{r6}/{r7}/{r8}/{r9}"
	endpoints[9] = "{r1}/{r2}/{r3}/{r4}/{r5}/{r6}/{r7}/{r8}/{r9}/{r10}"
	return endpoints
}

func getProxyEndpointMethods() [5]string {
	var methods [5]string
	methods[0] = "GET"
	methods[1] = "POST"
	methods[2] = "PUT"
	methods[3] = "PATCH"
	methods[4] = "DELETE"
	return methods
}
