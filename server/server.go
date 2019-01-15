package server

import (
	"context"
	"encoding/json"
	"github.com/ftob/ta/health"
	"github.com/ftob/ta/index"
	"github.com/ftob/ta/notallow"
	"github.com/ftob/ta/notfound"
	"github.com/go-chi/chi"
	kitlog "github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// Server holds the dependencies for a HTTP server.
type Server struct {
	Index    index.Service
	Health   health.Service
	Logger   kitlog.Logger
	NotFound notfound.Service
	NotAllow notallow.Service
	router   chi.Router
}

// New returns a new HTTP server.
func New(ix index.Service, hlt health.Service, nf notfound.Service, mna notallow.Service, logger kitlog.Logger) *Server {
	s := &Server{
		Index:    ix,
		Health:   hlt,
		NotFound: nf,
		NotAllow: mna,
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

	nah := notallowHandler{s.NotAllow, s.Logger}
	r.MethodNotAllowed(nah.methodNotAllow)

	nfh := notfoundHandler{s.NotFound, s.Logger}
	r.NotFound(nfh.notFound)

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

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case notallow.ErrorMethodNotAllow:
		w.WriteHeader(http.StatusMethodNotAllowed)
	// add custom error
	case notfound.ErrorNotFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	ctx = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
