package notfound

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

func (s *loggingService) NotFound() (res []byte, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "NotFound",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.next.NotFound()
}
