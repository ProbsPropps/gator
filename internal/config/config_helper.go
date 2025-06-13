package config

import (
	"encoding/json"
	"fmt"
	"os"
)

//File at end of path
const configFileName = ".gatorconfig.json"

//File should live at "~/.gatorconfig.json"
func getConfigFilePath() (string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	filePath := homeDir + "/" + configFileName
	return filePath
}

func write(cfg Config) error{
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Error - write: %v", err)
	}
	filePath := getConfigFilePath()

	if err = os.WriteFile(filePath, jsonData, 0o600); err != nil {
		return fmt.Errorf("Error - write: %v", err)
	}
	return nil
}
