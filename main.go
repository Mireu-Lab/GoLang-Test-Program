package main

import "github.com/gin-gonic/gin"

func json_read() []byte {
	test_read()
	return
}

func json_write(email string, text string) string {

	return
}

func main() {
	r := gin.Default()

	r.GET("/")

	r.Run(":8080")
}
