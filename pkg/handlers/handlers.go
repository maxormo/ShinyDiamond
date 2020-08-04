package handlers

import (
	"context"
	"fmt"
	"net/http"
)

func Print_headers(ctx context.Context, w http.ResponseWriter, h http.Header) {
	for n, h := range h {
		for _, head := range h {
			fmt.Fprintf(w, "%v=%v\n", n, head)
		}
	}
}
