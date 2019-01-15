package index

import (
	"bytes"
	"testing"
)

func TestSayHello(t *testing.T) {
	s := NewService()
	if say, err := s.SayHello(); !(bytes.Equal(say, []byte("Hello world")) || err == nil) {
		t.Fail()
	}
}
