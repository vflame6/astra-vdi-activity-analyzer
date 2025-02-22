package utils

import (
	"encoding/json"
	"os"
)

type Config struct {
	Processing        bool   `json:"processing"`
	ProcessingAddress string `json:"processing_address"`
	Password          string `json:"password"`
}

func LoadConfig(filepath string) (*Config, error) {
	configFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(configFile, config)
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
