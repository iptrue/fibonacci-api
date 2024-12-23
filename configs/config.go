package configs

import (
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	//API
	MaxFibonacciN int64         `yaml:"max_fibonacci_n" validate:"required,min=1"`
	GRPCPort      string        `yaml:"grpc_port" validate:"required"`
	GRPCTimeout   time.Duration `yaml:"grpc_timeout" validate:"required,min=1s"`

	//Caching
	RedisAddr       string        `yaml:"redis_addr" validate:"required,hostname_port"`
	CacheExpiration time.Duration `yaml:"cache_expiration" validate:"required,min=1s"`
	LocalCacheSize  int           `yaml:"local_cache_size" validate:"required,min=100"`

	//Logger
	LogLevel  string `yaml:"log_level" validate:"required,oneof=debug info warn error fatal"`
	LogOutput string `yaml:"log_output" validate:"required,oneof=console file"`

	//RateLimiter
	MaxTokens  int `yaml:"max_tokens" validate:"required,min=1"`
	RefillRate int `yaml:"refill_rate" validate:"required,min=1"`

	//Metrics
	PrometheusAddr string `yaml:"prometheus_addr" validate:"required,min=1"`
}

func LoadConfig(filePath string) (*AppConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var config AppConfig
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	if err := validateConfig(&config); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &config, nil
}

// validateConfig выполняет валидацию параметров конфигурации.
func validateConfig(config *AppConfig) error {
	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	return nil
}
