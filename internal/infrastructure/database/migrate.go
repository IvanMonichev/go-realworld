package database

import (
	"go-realworld/internal/users"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&users.UserModel{})
}
