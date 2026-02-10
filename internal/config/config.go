package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL	  	 string `json:"db_url"`
	CurrentUser  string `json:"current_user_name"`
}


func getConfigFilePath() (string, error) {
    home, err := os.UserHomeDir()
	if err != nil {
		return "", err	
	}

    fullPath := filepath.Join(home, configFileName)
    return fullPath, nil
}

func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonEncoder := json.NewEncoder(file)
	if err := jsonEncoder.Encode(cfg);  err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}


func (c *Config) SetUser(userName string) error {
	c.CurrentUser = userName
	return write(*c)
}