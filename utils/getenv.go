package utils

import (
	"fmt"
	"log"
	"os"
)

func GetEnv(envName string) string {
	env := os.Getenv(envName)

	if env == "" {
		log.Fatal(fmt.Sprintf("%s is not defined", envName))
	}

	return env
}
