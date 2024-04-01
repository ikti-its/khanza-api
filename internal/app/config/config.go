package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
}

func (c *Config) Get(key string) string {
	return os.Getenv(key)
}

func (c *Config) GetInt(key string, def int) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Printf("Failed to parse %s to int: %v", key, err)
		return def
	}

	return value
}

func NewConfig() *Config {
	return &Config{}
}
