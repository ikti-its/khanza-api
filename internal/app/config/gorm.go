package config

import (
	"fmt"
	"github.com/fathoor/simkes-api/internal/app/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func ProvideDB(cfg Config) *gorm.DB {
	var (
		host         = cfg.Get("PG_HOST")
		port         = cfg.Get("PG_PORT")
		user         = cfg.Get("PG_USER")
		password     = cfg.Get("PG_PASSWORD")
		dbname       = cfg.Get("PG_NAME")
		connOpen     = cfg.GetInt("PG_CONN_OPEN")
		connIdle     = cfg.GetInt("PG_CONN_IDLE")
		connLifeTime = cfg.GetInt("PG_CONN_LIFETIME")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, user, password, dbname)
	dialect := postgres.Open(dsn)

	db, err := gorm.Open(dialect, &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})
	exception.PanicIfError(err)

	psql, err := db.DB()
	exception.PanicIfError(err)

	psql.SetMaxOpenConns(connOpen)
	psql.SetMaxIdleConns(connIdle)
	psql.SetConnMaxLifetime(time.Minute * time.Duration(connLifeTime))

	return db
}
