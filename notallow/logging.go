package notallow

import (
	"github.com/go-kit/kit/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	next   Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) MethodNotAllow() (res []byte, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "MethodNotAllow",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.next.MethodNotAllow()
}
