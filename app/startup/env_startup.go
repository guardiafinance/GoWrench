package startup

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"wrench/app"
)

func LoadEnvsFiles() {
	envFolder := os.Getenv(app.ENV_PATH_FOLDER_ENV_FILES)

	if len(envFolder) == 0 {
		envFolder = "./"
	}

	envPath := fmt.Sprintf("%s.ENV", envFolder)
	setEnvFileToSystemEnv(envPath)

	envValue := os.Getenv(app.ENV_APP_ENV)
	envPathEnvironment := fmt.Sprintf("%s.ENV.%s", envFolder, envValue)
	setEnvFileToSystemEnv(envPathEnvironment)
}

func EnvInterpolation(value []byte) []byte {
	valueString := string(value)

	var envs = os.Environ()
	for _, env := range envs {
		envArray := strings.Split(env, "=")
		envKey := envArray[0]
		envValue := envArray[1]

		toReplace := fmt.Sprintf("{{%s}}", envKey)
		if toReplace != "{{}}" {
			valueString = strings.ReplaceAll(valueString, toReplace, envValue)
		}
	}

	return []byte(valueString)
}

func setEnvFileToSystemEnv(pathEnvFile string) {
	file, err := os.Open(pathEnvFile)
	defer file.Close()

	if err != nil {
		if os.IsNotExist(err) {
			log.Print(fmt.Sprintf("Env file %s not found ", pathEnvFile))
			return
		} else {
			log.Fatal(err)
		}
	}
	r := bufio.NewReader(file)

	log.Print(fmt.Sprintf("Loading file %s", pathEnvFile))
	for {
		line, _, err := r.ReadLine()
		if err != nil && fmt.Sprint(err) != "EOF" {
			log.Fatal(err)
		}

		if len(line) > 0 {
			lineText := string(line)
			if lineText[0] != '#' {
				arrayLineText := strings.Split(lineText, "=")
				envKey := arrayLineText[0]
				envValue := arrayLineText[1]

				if !strings.ContainsAny(envKey, " ") {
					os.Setenv(envKey, envValue)
				}
			}
		} else {
			break
		}
	}
	log.Print(fmt.Sprintf("Done file %s", pathEnvFile))
}
