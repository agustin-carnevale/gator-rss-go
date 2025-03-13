package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/agustin-carnevale/gator-rss-go/internal/database"
)

type State struct {
	Config    *Config
	DBQueries *database.Queries
}

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFilename = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configFilePath := homePath + "/" + configFilename
	return configFilePath, nil
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return Config{}, err
	}

	configFileBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(configFileBytes, &config)
	if err != nil {
		fmt.Println("Error parsing config json:", err)
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName

	configFilePath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return err
	}

	configFileBytes, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		fmt.Println("Error parsing Config struct to json:", err)
		return err
	}

	err = os.WriteFile(configFilePath, configFileBytes, 0644)
	if err != nil {
		fmt.Println("Error writing config file:", err)
		return err
	}

	return nil
}
