package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const (
	DEFAULT_ENVIRONMENT = "dev"
	DEFAULT_LOG_LEVEL   = "Error"
	DEFAULT_HOST        = "localhost"
	DEFAULT_PORT        = "8000"
)

type Config struct {
	App ApplicationParams `yaml:"app"`
}

type ApplicationParams struct {
		Environment string `yaml:"environment"`
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		LogLevel    string `yaml:"log-level"`
}

func Load(configPath string) (*Config, error) {
	fileName, err := filepath.Abs(configPath)
	if err != nil {
		return nil, err
	}
	yamlData, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	config := &Config{
		App: ApplicationParams {
			App.Environment: DEFAULT_ENVIRONMENT,
			App.Host:        DEFAULT_HOST,
			App.Port:        DEFAULT_PORT,
			App.LogLevel:    DEFAULT_LOG_LEVEL,
		}
	}
	err = yaml.Unmarshal(yamlData, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
