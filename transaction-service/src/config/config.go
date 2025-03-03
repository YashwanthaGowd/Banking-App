package config

import (
	"os"

	"go.uber.org/fx"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server  Server  `yaml:"server"`
	MongoDB MongoDB `yaml:"mongodb"`
	Kafka   Kafka   `yaml:"kafka"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type MongoDB struct {
	URI string `yaml:"uri"`
}

type Kafka struct {
	Brokers       []string `yaml:"brokers"`
	Topic         string   `yaml:"topic"`
	ConsumerGroup string   `yaml:"consumer_group"`
}

func LoadFromFile() (*Config, error) {
	file, err := os.ReadFile(os.Getenv("CONFIG_FILE"))
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil

}

func Module(configPath string) interface{} {
	return fx.Options(
		
		fx.Provide(configPath),
	)
}

