package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

var (
	root, _  = os.UserConfigDir()
	filename = ".salvare.config.json"
)

type Config struct {
	Conn struct {
		Protocol string `json:"protocol"`
		Username string `json:"username"`
		Password string `json:"password"`
		Hostname string `json:"host"`
		Port     string `json:"port"`
		Database string `json:"databaseName"`
		Query    string `json:"query"`
	} `json:"connection"`
}

func LoadConfig() (*Config, error) {
	path := filepath.Join(root, filename)
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			config := &Config{}
			WriteConfig(config)
			return config, nil
		}
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func WriteConfig(config *Config) error {
	path := filepath.Join(root, filename)

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(config); err != nil {
		return err
	}

	return nil
}
