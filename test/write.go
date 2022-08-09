package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type file_Message struct {
	Id   int    `json:id`
	Text string `json:"text"`
}

var num = 0

func Json_write(app *gin.Context) {
	var file_Messages []file_Message
	var post_Messages file_Message

	if err := app.BindJSON(&post_Messages); err != nil {
		app.IndentedJSON(http.StatusCreated, gin.H{"message": "not json type "})
	}

	bytes, _ := ioutil.ReadFile("1.json")

	json.Unmarshal([]byte(bytes), &file_Messages)

	messages := append(file_Messages, file_Message{num, string(post_Messages.Text)})

	data, _ := json.Marshal(messages)

	_ = ioutil.WriteFile("1.json", data, 0)

	num += 1
	app.IndentedJSON(http.StatusCreated, gin.H{"message": "clear"})
}
