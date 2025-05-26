package main

import (
	"github.com/gin-gonic/gin"
	"go-realworld/internal/infrastructure/database"
)

func main() {
	db := database.Init("real_world.db")
	database.Migrate(db)

	r := gin.Default()

	r.Run(":4000")
}
