package main

import (
	"log"
	"net/http"
	handlershttp "wrench/app/handlers/http"
	"wrench/app/manifest"
)
func main()  {
	configApp, err := manifest.LoadYamlFile()

	if err != nil {
		log.Fatalf("Error loading YAML: %v", err)
	}


	r := handlershttp.CreateAPIRoutes(configApp)	

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
