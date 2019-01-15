package index

import (
	"github.com/go-kit/kit/metrics"
	"time"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           Service
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		next:           s,
	}
}

func (s *instrumentingService) SayHello() (say []byte, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "say").Add(1)
		s.requestLatency.With("method", "say").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.SayHello()
}
