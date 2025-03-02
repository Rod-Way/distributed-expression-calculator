package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	// Rest Server
	RestServerPort string `env:"REST_SERVER_PORT" env-default:"localhost:8080"`

	// MongoDB
	MongoDBUri      string `env:"MONGODB_URI" env-default:"mongodb://localhost:27017"`
	MongoDBDatabase string `env:"MONGODB_DB" env-default:"calculator"`

	// Duration of arithmetic operations
	TimeAddition        int `env:"TIME_ADDITION_MS" env-default:"100"`
	TimeSubtraction     int `env:"TIME_SUBTRACTION_MS" env-default:"100"`
	TimeMultiplications int `env:"TIME_MULTIPLICATIONS_MS" env-default:"200"`
	TimeDivisions       int `env:"TIME_DIVISIONS_MS" env-default:"200"`
}

func New(configName string) (*Config, error) {
	cfg := &Config{}

	configPath := fmt.Sprintf("./backend/configs/%s", configName)

	if err := readConfig(configPath, cfg); err != nil {
		log.Printf("warning: failed to read config by path: %d", err)
	} else {
		return cfg, nil
	}

	log.Println("message: trying to read environment variables ...")

	if err := readEnvVars(cfg); err != nil {
		log.Printf("warning: failed to read config by path: %d", err)
	} else {
		return cfg, nil
	}

	log.Panicf("error: config not found")
	return nil, fmt.Errorf("error: config not found")
}

func readConfig(configPath string, cfg *Config) error {
	return cleanenv.ReadConfig(configPath, &cfg)
}

func readEnvVars(cfg *Config) error {
	return cleanenv.ReadEnv(cfg)
}
