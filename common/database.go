package common

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./../gorm.db"), &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Fatal("db err: (Init)")
	}

	sqlDB, err := db.DB()
	if err != nil {
		logrus.WithError(err).Fatal("failed to get *sql.DB from gorm.DB")
	}

	sqlDB.SetMaxIdleConns(10)

	DB = db
	return DB
}

func Close() {
	if DB == nil {
		return
	}

	sqlDB, err := DB.DB()
	if err != nil {
		logrus.WithError(err).Error("failed to get *sql.DB during Close()")
		return
	}

	if err := sqlDB.Close(); err != nil {
		logrus.WithError(err).Error("failed to close DB connection")
	}
}
