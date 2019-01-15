package notallow

import (
	"bytes"
	"testing"
)

func TestMethodNotAllow(t *testing.T)  {
	s := NewService()
	if r, err := s.MethodNotAllow(); !(bytes.Equal(r, []byte("Method not allow")) && err == ErrorMethodNotAllow) {
		t.Fail()
	}
}
