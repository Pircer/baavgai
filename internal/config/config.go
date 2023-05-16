package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

const (
	defaultEnvironment = "dev"
	defaultLogLevel    = "error"
	defaultHost        = "localhost"
	defaultPort        = "8000"
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

func Load(configPath string) (Config, error) {
	fileName, err := filepath.Abs(configPath)
	if err != nil {
		return Config{}, err
	}
	yamlData, err := os.ReadFile(fileName)
	if err != nil {
		return Config{}, err
	}
	config := Config{
		App: ApplicationParams{
			Environment: defaultEnvironment,
			Host:        defaultHost,
			Port:        defaultPort,
			LogLevel:    defaultLogLevel,
		},
	}
	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
