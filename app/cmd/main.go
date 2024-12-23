package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"wrench/app"
	"wrench/app/manifest/application_settings"
	"wrench/app/startup"
	"wrench/app/startup/token_credentials"
)

func main() {
	loadBashFiles()

	startup.LoadEnvsFiles()

	byteArray, err := startup.LoadYamlFile(getFileConfigPath())
	startup.LoadAwsSecrets(byteArray)
	if err != nil {
		log.Printf("Error loading YAML: %v", err)
	}

	byteArray = startup.EnvInterpolation(byteArray)
	applicationSetting, err := application_settings.ParseToApplicationSetting(byteArray)

	if err != nil {
		log.Printf("Error parse yaml: %v", err)
	}

	application_settings.ApplicationSettingsStatic = applicationSetting

	go token_credentials.LoadTokenCredentialAuthentication()
	var router = startup.LoadApplicationSettings(applicationSetting)
	port := getPort()
	log.Print(fmt.Sprintf("Server listen in port %s", port))
	http.ListenAndServe(port, router)
}

func loadBashFiles() {
	envbashFiles := os.Getenv(app.ENV_RUN_BASH_FILES_BEFORE_STARTUP)

	if len(envbashFiles) == 0 {
		envbashFiles = "wrench/bash/startup.sh"
	}

	bashFiles := strings.Split(envbashFiles, ",")
	bashRun(bashFiles)
}

func bashRun(paths []string) {
	for _, path := range paths {
		path = strings.TrimSpace(path)
		if _, err := os.Stat(path); err != nil {
			log.Print(fmt.Sprintf("file bash %s not found", path))
			continue
		}

		log.Print(fmt.Sprintf("Will process file bash %s", path))
		cmd := exec.Command("/bin/sh", "./"+path)

		output, err := cmd.Output()
		if err != nil {
			log.Print("Error: ", err)
			return
		} else {
			log.Print(output)
		}
	}
}

func getFileConfigPath() string {
	configPath := os.Getenv(app.ENV_PATH_FILE_CONFIG)
	if len(configPath) == 0 {
		configPath = "../../configApp.yaml"
	}
	return configPath
}

func getPort() string {
	port := os.Getenv(app.ENV_PORT)
	if len(port) == 0 {
		port = ":9090"
	}

	if port[0] != ':' {
		port = ":" + port
	}

	return port
}
