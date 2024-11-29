package main

import (
	"fmt"
	"log"
	"wrench/app/manifest"
	startup "wrench/app/startup"
)

func main() {
	applicationSetting, err := manifest.LoadYamlFile("../../configApp.yaml")

	if err != nil {
		log.Fatalf("Error loading YAML: %v", err)
	}

	var result = applicationSetting.Valid()

	if result.HasError() == true {
		var errors = result.GetError()
		for _, error := range errors {
			fmt.Println(error)
		}
	} else {
		startup.LoadApplicationSettings(applicationSetting)
	}
}
