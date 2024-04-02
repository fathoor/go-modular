package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
}

func (c *Config) Get(key, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	log.Printf("Failed to get %s, using default value: %s", key, def)
	return def
}

func (c *Config) GetInt(key string, def int) int {
	if value, err := strconv.Atoi(os.Getenv(key)); err != nil {
		return value
	}

	log.Printf("Failed to parse %s to int, using default value: %d", key, def)
	return def
}

func NewConfig() *Config {
	return &Config{}
}
