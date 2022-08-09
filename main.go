package main

import (
	"github.com/Mireu-Lab/GoLang-Test-Program/test"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	// app.GET("/")
	app.GET("/", test.Json_read)
	app.POST("/", test.Json_write)

	app.Run(":8080")
}
