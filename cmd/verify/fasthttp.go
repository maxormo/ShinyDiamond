package main

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	fmt.Fprintln(ctx.Response.BodyWriter(), "hello mister")
}

func startFastHttp() {
	r := router.New()
	r.GET("/api/fasthttp", HandleFastHTTP)
	fasthttp.ListenAndServe("0.0.0.0:8080", r.Handler)
}

func main() {
	startFastHttp()
}
