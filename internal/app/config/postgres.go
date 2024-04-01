package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func NewPostgres(cfg *Config) *gorm.DB {
	var (
		dsn          = cfg.Get("PG_DSN")
		connOpen     = cfg.GetInt("PG_CONN_OPEN", 100)
		connIdle     = cfg.GetInt("PG_CONN_IDLE", 10)
		connLifeTime = cfg.GetInt("PG_CONN_LIFETIME", 15)
	)

	dialect := postgres.Open(dsn)

	loggerConfig := logger.Default.LogMode(logger.Info) // Info on development, Error on production

	gormConfig := &gorm.Config{
		Logger:         loggerConfig,
		TranslateError: true,
	}

	db, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	psql, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to create postgres instance: %v", err)
	}

	psql.SetMaxOpenConns(connOpen)
	psql.SetMaxIdleConns(connIdle)
	psql.SetConnMaxLifetime(time.Minute * time.Duration(connLifeTime))

	return db
}
