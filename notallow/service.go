package notallow

import "errors"

var ErrorMethodNotAllow = errors.New("405")

type Service interface {
	MethodNotAllow() ([]byte, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) MethodNotAllow() ([]byte, error) {
	return []byte(`Method not allow`), ErrorMethodNotAllow
}
