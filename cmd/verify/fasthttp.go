package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type MyHandler struct {
	foobar string
}

// request handler in net/http style, i.e. method bound to MyHandler struct.
func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	// notice that we may access MyHandler properties here - see h.foobar.
	body := ctx.Request.Body()
	fmt.Sprintln(string(body))
}

var fastServer fasthttp.Server

func startFastHttp() {
	myHandler := &MyHandler{
		foobar: "foobar",
	}
	fastServer = fasthttp.Server{
		Handler: myHandler.HandleFastHTTP,
	}
	go fastServer.ListenAndServe(":8080")

}
func stopFastHttp() {
	//fastServer.Shutdown()
}
