package utils

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	ClientName string `json:"client_name"`
	Address    string `json:"address"`
	UseTLS     bool   `json:"use_tls"`
	Key        string `json:"key"`
}

func LoadConfig(filepath string) (*Config, error) {
	configFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func SaveConfig(filepath string, config *Config) error {
	configFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer configFile.Close()
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = configFile.Write(data)
	if err != nil {
		return err
	}
	return nil
}
