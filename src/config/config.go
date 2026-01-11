package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
}
type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
	Host string `mapstructure:"host"`
	Env  string `mapstructure:"env"`
}

var cfg *Config

// LoadConfig loads configuration based on environment
func LoadConfig() (*Config, error) {
	// Set config file name based on environment
	configName := "config"

	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./src/config")

	// Enable reading from environment variables
	viper.AutomaticEnv()
	//viper.SetEnvPrefix("MH") // Optional: prefix for env vars like MH_SERVER_PORT

	// Bind specific environment variables to config keys
	viper.BindEnv("server.mode", "SERVER_MODE")
	viper.BindEnv("server.port", "SERVER_PORT")

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Get environment from ENV variable, default to "dev"
	env := os.Getenv("ENV")
	if env != "" {
		// Set config file name based on environment
		viper.SetConfigName(fmt.Sprintf("config.%s", env))
		if err := viper.MergeInConfig(); err != nil {
			return nil, fmt.Errorf("error reading env config: %w", err)
		}
	}

	log.Printf("Loading configuration from: %s", viper.ConfigFileUsed())

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	// Validate configuration
	if err := validateConfig(&config); err != nil {
		return nil, err
	}
	cfg = &config
	return &config, nil
}
func validateConfig(cfg *Config) error {
	if cfg.Server.Port == "" {
		return fmt.Errorf("server port is required")
	}
	return nil
}

func GetConfig() *Config {
	return cfg
}
