package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

var run = true
var client = &http.Client{}
var payload = getPayload()

func runBootstrapOnce(bootstrapCode func()) {
	if run {
		bootstrapCode()
		run = false
	}
}

func Benchmark_HTTP(b *testing.B) {

	runBootstrapOnce(func() {
		startHttp()
		time.Sleep(1 * time.Second)
		fmt.Println("hello from http")
	})

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		sendRequest(client, payload)
	}
	b.StopTimer()
}
func Benchmark_SequenceHTTPHandler(b *testing.B) {

	runBootstrapOnce(func() {
		startHttpWithSequeenceHandler()
		time.Sleep(1 * time.Second)
		fmt.Println("hello from sequence handler http")
	})

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		sendRequest(client, payload)
	}
	b.StopTimer()
}

func Benchmark_FastHTTP(b *testing.B) {
	runBootstrapOnce(func() {
		startFastHttp()
		time.Sleep(1 * time.Second)
		fmt.Println("hello from fast http")
	})

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		sendRequest(client, payload)
	}
	b.StopTimer()
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintln("hello")
	}
}
