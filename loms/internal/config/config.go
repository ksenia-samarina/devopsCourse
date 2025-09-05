package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const configFile = "config.yml"

type Config struct {
	PortGRPC string   `yaml:"port-grpc"`
	PortREST string   `yaml:"port-rest"`
	Storages Storages `yaml:"storages"`
}

type Storages struct {
	Postgres        string `yaml:"postgres"`
	MinConnections  int32  `yaml:"minConnections"`
	MaxConnections  int32  `yaml:"maxConnections"`
	MaxConnIdleTime string `yaml:"maxConnIdleTime"`
	MaxConnLifetime string `yaml:"maxConnLifetime"`
}

type CancelOrders struct {
	Timeout  string `yaml:"timeout"`
	Interval string `yaml:"interval"`
}

func Init() (Config, error) {
	rawYAML, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %s", err.Error())
	}

	var configData Config
	err = yaml.Unmarshal(rawYAML, &configData)
	if err != nil {
		return Config{}, fmt.Errorf("error parsing yaml: %s", err.Error())
	}

	return configData, nil
}
