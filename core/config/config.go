package config

import (
	"github.com/fathoor/simkes-api/core/exception"
	"os"
	"strconv"
)

type Config interface {
	Get(key string) string
	GetInt(key string) int
	GetBool(key string) bool
}

type configImpl struct {
}

func (c *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func (c *configImpl) GetInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	exception.PanicIfError(err)

	return value
}

func (c *configImpl) GetBool(key string) bool {
	value, err := strconv.ParseBool(os.Getenv(key))
	exception.PanicIfError(err)

	return value
}

func ProvideConfig() Config {
	return &configImpl{}
}
