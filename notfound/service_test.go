package notfound

import (
	"bytes"
	"testing"
)


func TestNotFound(t *testing.T)  {
	s := NewService()
	if r, err := s.NotFound(); !(bytes.Equal(r, []byte("Not found")) && err == ErrorNotFound) {
		t.Fail()
	}
}

