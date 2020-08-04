package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"token-verify/pkg/handlers"
)

var server http.Server

func startHttp() {
	handler := MyHand{}

	server = http.Server{Addr: ":8080", Handler: handler}
	go server.ListenAndServe()

}

func startHttpWithSequeenceHandler() {
	handler := handlers.SequenceHandler{}
	handler.AddRequestHandler(func(ctx context.Context, writer http.ResponseWriter, header http.Header, body []byte) {
		fmt.Sprintln(string(body))
	})
	server = http.Server{Addr: ":8080", Handler: handler}
	go server.ListenAndServe()

}

type MyHand struct {
}

func (h MyHand) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Sprintln(string(body))
}

func main() {

	startFastHttp()
	time.Sleep(1 * time.Second)
	client := &http.Client{}
	payload := getPayload()
	for n := 0; n < 2; n++ {
		sendRequest(client, payload)
	}
	stopFastHttp()
}
