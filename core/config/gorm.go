package config

import (
	"github.com/fathoor/simkes-api/core/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func ProvideDB(cfg Config) *gorm.DB {
	var (
		host     = cfg.Get("DB_HOST")
		port     = cfg.Get("DB_PORT")
		user     = cfg.Get("DB_USER")
		password = cfg.Get("DB_PASSWORD")
		dbname   = cfg.Get("DB_NAME")
		maxOpen  = cfg.GetInt("DB_MAX_OPEN")
		maxIdle  = cfg.GetInt("DB_MAX_IDLE")
		maxLife  = cfg.GetInt("DB_MAX_LIFE")
	)

	dsn := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable TimeZone=Asia/Jakarta"
	dialect := postgres.Open(dsn)

	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	exception.PanicIfError(err)

	postgre, err := db.DB()
	exception.PanicIfError(err)

	postgre.SetMaxOpenConns(maxOpen)
	postgre.SetMaxIdleConns(maxIdle)
	postgre.SetConnMaxLifetime(time.Minute * time.Duration(maxLife))

	return db
}
