package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	LogFilePath string `yaml:"logfilepath"`
	MaxLogSize  int    `yaml:"maxlogsize"`
}

// LoadConfig 从 YAML 配置文件加载配置
func LoadConfig(filePath string) (*Config, error) {
	var config Config
	if err := loadYAML(filePath, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// LoadyamlConfig 从 YAML 配置文件加载配置
func loadYAML(filePath string, out interface{}) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, out)
}
