package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
)

const payloadPath = "payload.json"

func getPayload() []byte {
	file, err := os.Open(payloadPath)
	if err != nil {
		panic(err)
	}

	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return all
}

func sendRequest(client *http.Client, b []byte) {
	reader := bytes.NewReader(b)
	res, err := client.Post("http://127.0.0.1:8080", "content-type: application.json", reader)

	if err != nil {
		panic(err)
	}

	if res.StatusCode != 200 {
		panic("request failed")
	}
}
