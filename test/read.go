package test

import "io/ioutil"

func json_read() []byte {
	bytes, _ := ioutil.ReadFile("1.json")

	return bytes
}
