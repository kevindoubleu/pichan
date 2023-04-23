package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

func NewConfig(path string) (*Config, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	config := &Config{}
	err = yaml.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
