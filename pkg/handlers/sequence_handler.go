package handlers

import (
	"context"
	"io/ioutil"
	"net/http"
)

type HeadersOnlyHandler func(context.Context, http.ResponseWriter, http.Header)
type FullRequestHandler func(context.Context, http.ResponseWriter, http.Header, []byte)

var _ http.Handler = SequenceHandler{}

type SequenceHandler struct {
	headersHandlers []HeadersOnlyHandler
	requestHandlers []FullRequestHandler
}

func (h *SequenceHandler) AddHaderHandler(handler HeadersOnlyHandler) {
	h.headersHandlers = append(h.headersHandlers, handler)
}

func (h *SequenceHandler) AddRequestHandler(handler FullRequestHandler) {
	h.requestHandlers = append(h.requestHandlers, handler)
}

func (h SequenceHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	for i := 0; i < len(h.headersHandlers); i++ {
		f := h.headersHandlers[i]
		f(req.Context(), w, req.Header)
	}

	if len(h.requestHandlers) == 0 {
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		panic(err)
	}

	for i := 0; i < len(h.requestHandlers); i++ {
		f := h.requestHandlers[i]
		f(req.Context(), w, req.Header, body)
	}
}