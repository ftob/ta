package server

import (
	"context"
	"encoding/json"
	"github.com/ftob/ta/notallow"
	"github.com/go-chi/chi"
	kitlog "github.com/go-kit/kit/log"
	"net/http"
)

type notallowHandler struct {
	s      notallow.Service
	logger kitlog.Logger
}

func (h *notallowHandler) router() chi.Router {
	r := chi.NewRouter()

	r.MethodNotAllowed(h.methodNotAllow)

	return r
}

func (h *notallowHandler) methodNotAllow(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	message, _ := h.s.MethodNotAllow()

	var response = struct {
		Message string `json:"message"`
	}{
		Message: string(message),
	}

	_ = h.logger.Log("code", 405, "addr", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		_ = h.logger.Log("error", err)
		encodeError(ctx, err, w)
		return
	}
}
