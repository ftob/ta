package notallow

import (
	"github.com/go-kit/kit/metrics"
	"time"
)

type instrumentingService struct {
	requestCount metrics.Counter
	next         Service
}

func NewInstrumentingService(counter metrics.Counter, s Service) Service {
	return &instrumentingService{
		requestCount: counter,
		next:         s,
	}
}

func (s *instrumentingService) MethodNotAllow() (say []byte, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "MethodNotAllow").Add(1)
	}(time.Now())

	return s.next.MethodNotAllow()
}
