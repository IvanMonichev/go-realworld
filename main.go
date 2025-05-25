package main

import (
	"github.com/gin-gonic/gin"
	"go-realworld/common"
)

func main() {
	common.Init()
	defer common.Close()

	r := gin.Default()

	r.Run(":4000")
}
