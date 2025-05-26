package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(dbPath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"dbPath": dbPath,
			"error":  err,
		}).Fatal("failed to connect to SQLite database")
	}

	logrus.Info("connected to SQLite database")

	return db
}
