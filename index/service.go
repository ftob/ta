package index

//
type Service interface {
	SayHello() ([]byte, error)
}

//
type service struct{}

//
func NewService() Service {
	return &service{}
}

//
func (s *service) SayHello() ([]byte, error) {
	return []byte(`Hello world`), nil
}
