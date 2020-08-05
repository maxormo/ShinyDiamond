package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
}
func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}
