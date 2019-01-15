package notfound

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

func (s *instrumentingService) NotFound() (say []byte, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "NotFound").Add(1)
	}(time.Now())

	return s.next.NotFound()
}
