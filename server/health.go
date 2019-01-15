package server

import (
	"context"
	"encoding/json"
	"github.com/ftob/ta/health"
	"github.com/go-chi/chi"
	kitlog "github.com/go-kit/kit/log"
	"net/http"
)

type healthHandler struct {
	s      health.Service
	logger kitlog.Logger
}

func (h *healthHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Get("/health", h.health)

	r.Method("GET", "/docs", http.StripPrefix("/v1/docs", http.FileServer(http.Dir("index/docs"))))

	return r
}

func (h *healthHandler) health(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	response := h.s.Health()

	w.Header().Set("Content-Type", "application/health+json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		_ = h.logger.Log("error", err)
		encodeError(ctx, err, w)
		return
	}
}
