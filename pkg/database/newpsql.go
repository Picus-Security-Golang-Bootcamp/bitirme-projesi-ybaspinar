package database

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func Connect(cfg *config.Config) *gorm.DB {

	db, err := gorm.Open(postgres.Open(cfg.DBConfig.DataSourceName), &gorm.Config{})
	if err != nil {
		zap.L().Fatal("Database connection error", zap.Error(err))
	}

	origin, err := db.DB()
	if err != nil {
		zap.L().Fatal("Database connection error", zap.Error(err))
	}

	err = origin.Ping()
	if err != nil {
		zap.L().Fatal("Database connection error", zap.Error(err))
	}
	origin.SetMaxOpenConns(cfg.DBConfig.MaxOpen)
	origin.SetMaxIdleConns(cfg.DBConfig.MaxIdle)
	origin.SetConnMaxLifetime(time.Duration(cfg.DBConfig.MaxLifetime) * time.Second)
	return db
}
