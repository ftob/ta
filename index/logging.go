package index

import (
	"github.com/go-kit/kit/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	next Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) SayHello() (say []byte, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "SayHello",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.next.SayHello()
}


