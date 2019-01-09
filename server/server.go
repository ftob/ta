package server

import (
	"context"
	"encoding/json"
	"github.com/ftob/ta/health"
	"github.com/ftob/ta/index"
	"github.com/go-chi/chi"
	kitlog "github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// Server holds the dependencies for a HTTP server.
type Server struct {
	Index  index.Service
	Health	health.Service
	Logger kitlog.Logger

	router chi.Router
}

// New returns a new HTTP server.
func New(ix index.Service, hlt health.Service, logger kitlog.Logger) *Server {
	s := &Server{
		Index:  ix,
		Health: hlt,
		Logger:   logger,
	}

	r := chi.NewRouter()

	r.Use(accessControl)

	r.Route("/", func(r chi.Router) {
		h := indexHandler{s.Index, s.Logger}
		r.Mount("/", h.router())
	})

	r.Route("/service/", func(r chi.Router) {
		h := healthHandler{s.Health, s.Logger}
		r.Mount("/v1", h.router())
	})


	r.Method("GET", "/metrics", promhttp.Handler())

	s.router = r

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	// add custom error
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
