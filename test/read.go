package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Read_Message struct {
	ID   int    `json:id`
	Text string `json:"text"`
}

func Json_read(app *gin.Context) {
	var messages []Read_Message
	bytes, _ := ioutil.ReadFile("1.json")

	json.Unmarshal([]byte(bytes), &messages)

	app.IndentedJSON(http.StatusOK, messages)
}
