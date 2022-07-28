package commands

import (
	"encoding/json"
	"fmt"
	"os"
)

func getConfig() (string, string, error) {
	token := os.Getenv("TOKEN")
	domain := os.Getenv("DOMAIN")

	if token == "" || domain == "" {
		return "", "", fmt.Errorf("TOKEN and DOMAIN environment variables must be set")
	}

	return token, domain, nil
}

func printJSON(d interface{}) (string, error) {
	val, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
