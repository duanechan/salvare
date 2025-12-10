package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	cwd, _   = os.Getwd()
	filename = "salvare.config.json"
)

const (
	EmptyConnString string = "://:@:/"
)

type Config struct {
	Conn struct {
		Scheme   string `json:"scheme"`
		Username string `json:"username"`
		Password string `json:"password"`
		Hostname string `json:"host"`
		Port     string `json:"port"`
		Database string `json:"databaseName"`
		Query    string `json:"query"`
	} `json:"connection"`
	BackupDirectory string `json:"backupDirectory"`
}

func (c Config) ConnectionString() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s%s",
		c.Conn.Scheme,
		c.Conn.Username,
		c.Conn.Password,
		c.Conn.Hostname,
		c.Conn.Port,
		c.Conn.Database,
		c.Conn.Query,
	)
}

func (c Config) IsEmpty() bool {
	return c.Conn.Scheme == "" &&
		c.Conn.Username == "" &&
		c.Conn.Password == "" &&
		c.Conn.Hostname == "" &&
		c.Conn.Port == "" &&
		c.Conn.Database == ""
}

var (
	ConfigFileNotExists = errors.New("configuration file does not exist")
	ConfigFileEOF       = errors.New("configuration file EOF")
)

func LoadConfig() (*Config, error) {
	path := filepath.Join(cwd, filename)
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, ConfigFileNotExists
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
		if errors.Is(err, io.EOF) {
			return nil, ConfigFileEOF
		}
		return nil, err
	}

	return &config, nil
}

func WriteConfig(config *Config) error {
	path := filepath.Join(cwd, filename)

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(config); err != nil {
		return err
	}

	return nil
}
