package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DBURL 				string `json:"db_url"`
	CurrentUserName 	string `json:"current_user_name"`
}

func Read () (Config, error) {
	filePath := getConfigFilePath()

	content, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("Error - Read: %v", err)
	}

	var cfg Config
	if err = json.Unmarshal(content, &cfg); err != nil {
		return Config{}, fmt.Errorf("Error - Read: %v", err)
	}

	return cfg, nil 
}

func (cfg *Config) SetUser(currentUserName string) error {
	cfg.CurrentUserName = currentUserName
	if err := write(*cfg); err != nil {
		return fmt.Errorf("Error - SetUser: %v", err)
	}
	return nil
}

