package test

import (
	"encoding/json"
	"io/ioutil"
)

type Message struct {
	ID    int    `json:ID`
	Email string `json:"Email"`
	Text  string `json:"Text"`
}

var num = 0

func json_write(email string, text string) {
	bytes, _ := ioutil.ReadFile("1.json")

	var Messages []Message
	json.Unmarshal([]byte(bytes), &Messages)

	messages := append(Messages, Message{num, email, text})
	data, _ := json.Marshal(messages)

	_ = ioutil.WriteFile("1.json", data, 0)
	num += 1

}
