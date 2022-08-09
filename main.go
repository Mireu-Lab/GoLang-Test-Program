package main

import (
	"github.com/gin-gonic/gin"
)

func json_read(app *gin.Context) {
	app.JSON(200, test.read.json_read())
}

func json_write(email string, text string) string {

	return
}

func main() {
	app := gin.Default()

	// app.GET("/")
	app.GET("/", json_read())

	app.Run(":8080")
}
