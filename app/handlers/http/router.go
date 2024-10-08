package handlershttp

import (
	"io"
	"log"
	"net/http"
	"wrench/app/manifest"

	"github.com/gorilla/mux"
)

func CreateAPIRoutes(configApp *manifest.Config) http.Handler {
	r := mux.NewRouter()


	actionsMap := make(map[string]manifest.Action)
	for _, action := range configApp.Actions {
			actionsMap[action.Id] = action
	}
	httpClientMap := make(map[string]manifest.HttpClient)
	for _, httpClient := range configApp.HttpClients {
		httpClientMap[httpClient.Id] = httpClient
	}

	for _, endpoint := range configApp.Api.Endpoints {
		ep := endpoint
		action := actionsMap[ep.ActionID]
		log.Printf("Adding route: %s, method: %s with action: %s", ep.Route, ep.Method, action.Id)

		r.HandleFunc(ep.Route, func(w http.ResponseWriter, r *http.Request) {
				var jsonResponse []byte
				var err error

				log.Printf("Handling request for path: %s, method: %s", ep.Route, ep.Method)
				w.Header().Set("Content-Type", action.Mock.ContentType)

				if action.Type == "httpRequestMock" {
					if action.Mock.Value != "" {
						log.Printf("Mocking response for action: %s with value: %s", action.Id, action.Mock.Value)
						w.WriteHeader(http.StatusOK)
						w.Write([]byte(action.Mock.Value))
					}
				}
				

				if action.Type == "httpRequest" {
					log.Printf("Sending request for action: %s", action.Id)

					httpClient, found := httpClientMap[action.HttpClientId]
					
					if !found {
						log.Printf("Error getting http client: %v", err)
						http.Error(w, "Error getting http client", http.StatusInternalServerError)
						return
					}

					url := httpClient.BaseUrl + action.Path

					log.Printf("URL: %s", url)
					resp, err := http.Get(url)

					if err != nil {
						log.Printf("Error sending request: %v", err)
						http.Error(w, "Error sending request", http.StatusInternalServerError)
						return
					}
					defer resp.Body.Close()

					body, err := io.ReadAll(resp.Body)
					if err != nil {
						log.Fatalf("Erro ao ler a resposta: %v", err)
					}
					w.Header().Set("Content-Type", "application/json")
					w.Write(body)
					return
				}
				

				_, err = w.Write(jsonResponse)
				if err != nil {
						log.Printf("Error writing response body: %v", err)
						http.Error(w, "Error writing response body", http.StatusInternalServerError)
				}
		}).Methods(ep.Method)
	}

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	log.Printf("routes created")

	return r
}
