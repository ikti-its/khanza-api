package test

import (
	"github.com/fathoor/simkes-api/core/config"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	cfg = config.ProvideConfig()
	db  = config.ProvideDB(cfg)
	app = config.ProvideApp(cfg)
)

func TestConfig(t *testing.T) {
	assert.NotNil(t, cfg)
}

func TestGORM(t *testing.T) {
	assert.NotNil(t, db)
}

func TestFiber(t *testing.T) {
	assert.NotNil(t, app)
}

func TestConfig_Get(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		t.Run("When key is valid", func(t *testing.T) {
			assert.NotEmpty(t, cfg.Get("APP_NAME"))
			assert.Equal(t, "simkes-api", cfg.Get("APP_NAME"))
		})

		t.Run("When key is invalid", func(t *testing.T) {
			assert.Empty(t, cfg.Get("INVALID_KEY"))
		})
	})

	t.Run("GetInt", func(t *testing.T) {
		t.Run("When key is valid", func(t *testing.T) {
			assert.NotEmpty(t, cfg.Get("APP_PORT"))
			assert.Equal(t, 1337, cfg.GetInt("APP_PORT"))
		})

		t.Run("When key is valid but not int", func(t *testing.T) {
			assert.NotEmpty(t, cfg.Get("APP_NAME"))
			assert.Panics(t, func() {
				cfg.GetInt("APP_NAME")
			})
		})
	})

	t.Run("GetBool", func(t *testing.T) {
		t.Run("When key is valid", func(t *testing.T) {
			assert.NotEmpty(t, cfg.Get("FIBER_PREFORK"))
			assert.Equal(t, false, cfg.GetBool("FIBER_PREFORK"))
		})

		t.Run("When key is valid but not bool", func(t *testing.T) {
			assert.NotEmpty(t, cfg.Get("APP_NAME"))
			assert.Panics(t, func() {
				cfg.GetBool("APP_NAME")
			})
		})
	})
}
