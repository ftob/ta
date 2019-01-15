package notfound

import "errors"

var ErrorNotFound = errors.New("404")

type Service interface {
	NotFound() ([]byte, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) NotFound() ([]byte, error) {
	return []byte(`Not found`), ErrorNotFound
}
