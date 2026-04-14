package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (config *Config) SetUser(userName string) {
	config.CurrentUserName = userName
	write(config)
}

func ReadConfig() (*Config, error) {
	configFile, err := os.Open(getConfigFilePath())
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config Config
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func getConfigFilePath() string {
	userHome, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Couldn't get the $HOME filepath: %v\n", err)
		return ""
	}
	return userHome + "/" + configFileName
}

func write(cfg *Config) error {
	jsonBytes, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(getConfigFilePath(), jsonBytes, 0644)
}
