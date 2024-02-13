package config

import (
	"fmt"
	"github.com/fathoor/simkes-api/core/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func ProvideDB(cfg Config) *gorm.DB {
	var (
		host        = cfg.Get("PG_HOST")
		port        = cfg.Get("PG_PORT")
		user        = cfg.Get("PG_USER")
		password    = cfg.Get("PG_PASSWORD")
		dbname      = cfg.Get("PG_NAME")
		maxOpen     = cfg.GetInt("PG_MAX_OPEN")
		maxIdle     = cfg.GetInt("PG_MAX_IDLE")
		maxLifeTime = cfg.GetInt("PG_MAX_LIFETIME")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, user, password, dbname)
	dialect := postgres.Open(dsn)

	db, err := gorm.Open(dialect, &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})
	exception.PanicIfError(err)

	postgre, err := db.DB()
	exception.PanicIfError(err)

	postgre.SetMaxOpenConns(maxOpen)
	postgre.SetMaxIdleConns(maxIdle)
	postgre.SetConnMaxLifetime(time.Minute * time.Duration(maxLifeTime))

	return db
}
