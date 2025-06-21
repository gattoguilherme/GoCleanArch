package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config holds the application configuration.
type Config struct {
	Env    string `yaml:"env"`
	Server ServerConfig `yaml:"server"`
	Dev    DevConfig    `yaml:"dev"`
	Prod   ProdConfig   `yaml:"prod"`
}

// ServerConfig holds the server configuration.
type ServerConfig struct {
	Port string `yaml:"port"`
}

// DevConfig holds the development environment configuration.
type DevConfig struct{}

// ProdConfig holds the production environment configuration.
type ProdConfig struct {
	AWS AWSConfig `yaml:"aws"`
	DB  DBConfig  `yaml:"db"`
}

// AWSConfig holds the AWS configuration.
type AWSConfig struct {
	Region      string `yaml:"region"`
	SQSQueueURL string `yaml:"sqs_queue_url"`
}

// DBConfig holds the database configuration.
type DBConfig struct {
	Driver string `yaml:"driver"`
	DSN    string `yaml:"dsn"`
}

var cfg *Config

// LoadConfig loads the configuration from the given file path.
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
