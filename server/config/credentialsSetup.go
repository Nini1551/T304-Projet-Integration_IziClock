package config

import (
	"encoding/json"
	"fmt"
	"os"
	"server/models"
)

func setupCredentials() {
	var c models.Credentials
	file, err := os.ReadFile("./creds.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &c)
}
