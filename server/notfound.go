package server

import (
	"context"
	"encoding/json"
	"github.com/ftob/ta/notfound"
	"github.com/go-chi/chi"
	kitlog "github.com/go-kit/kit/log"
	"net/http"
)

type notfoundHandler struct {
	s      notfound.Service
	logger kitlog.Logger
}

func (h *notfoundHandler) router() chi.Router {
	r := chi.NewRouter()

	r.NotFound(h.notFound)

	return r
}

func (h *notfoundHandler) notFound(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	message, _ := h.s.NotFound()

	var response = struct {
		Message string `json:"message"`
	}{
		Message: string(message),
	}

	_ = h.logger.Log("code", 404, "addr", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		_ = h.logger.Log("error", err)
		encodeError(ctx, err, w)
		return
	}
}
