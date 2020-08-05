package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	fmt.Println(ctx.Request.Body())
}

func startFastHttp() {
	fasthttp.ListenAndServe("8080", HandleFastHTTP)
}

func main() {
	startFastHttp()
}
