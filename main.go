package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ftob/ta/health"
	"github.com/ftob/ta/index"
	"github.com/ftob/ta/server"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	defaultPort = "8080"
	defaultVersion = "0.1.0"
	serviceID = "say_hello"
	componentID = "http_say_hello"
	componentType = "backend"
)

// PCHP - program code of a healthy person
func main() {

	startTime := time.Now()

	var (
		addr     = envString("APP_PORT", defaultPort)
		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
		ctx = context.Background()
	)

	flag.Parse()

	ctx = context.WithValue(ctx, "ServiceID", envString("APP_SERVICE_ID", defaultVersion))
	ctx = context.WithValue(ctx, "Version", envString("APP_VERSION", serviceID))
	ctx = context.WithValue(ctx, "ComponentId", envString("APP_COMPONENT_ID", componentID))
	ctx = context.WithValue(ctx, "ComponentType", envString("APP_COMPONENT_TYPE", componentType))
	ctx = context.WithValue(ctx, "startTime", startTime)



	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	fieldKeys := []string{"method"}

	var ix index.Service
	ix = index.NewService()
	ix = index.NewLoggingService(log.With(logger, "component", "index"), ix)
	ix = index.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "index_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "index_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		ix,
	)

	var hlth health.Service
	hlth = health.NewService(ctx)

	// Create http server
	srv := server.New(ix, hlth, log.With(logger, "component", "http"))

	errs := make(chan error, 2)
	go func() {
		_ = logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, srv)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	_ = logger.Log("terminated", <-errs)
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
