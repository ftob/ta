package server

import (
	"context"
	"encoding/json"
	"github.com/ftob/ta/index"
	"github.com/go-chi/chi"
	kitlog "github.com/go-kit/kit/log"
	"net/http"
)

type indexHandler struct {
	s      index.Service
	logger kitlog.Logger
}

func (h *indexHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.sayHello)

	r.Method("GET", "/docs", http.StripPrefix("/v1/docs", http.FileServer(http.Dir("index/docs"))))

	return r
}


func (h *indexHandler) sayHello(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	message, _ := h.s.SayHello()

	var response = struct {
		Message []byte `json:"message"`
	}{
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Log("error", err)
		encodeError(ctx, err, w)
		return
	}
}
