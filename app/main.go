package main

import (
	"log"
	"wrench/app/manifest"
)
func main()  {
	configApp, err := manifest.LoadYamlFile()

	if err != nil {
		log.Fatalf("Error loading YAML: %v", err)
	}

	log.Printf("Service: %v", configApp.Service.Name)
	log.Printf("Version: %v", configApp.Service.Version)
}
