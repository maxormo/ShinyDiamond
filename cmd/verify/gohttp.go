package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "hello mister")
}
func main() {
	http.HandleFunc("/api/v1", handler)

	http.ListenAndServe("0.0.0.0:8080", nil)
}
