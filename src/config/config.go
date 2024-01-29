package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type StorageConfig struct {
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name"`
}

func NewStorageConfig() (*StorageConfig, error) {
	file, err := os.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}
	s := &StorageConfig{}
	yaml.Unmarshal(file, s)
	return s, nil
}
