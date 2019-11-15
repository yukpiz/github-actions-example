package example

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func Test_Hello1(t *testing.T) {
	assert.Equal(t, Hello1(), "hello1")
}

func Test_Hello2(t *testing.T) {
	assert.Equal(t, Hello2(), "hello2")
}
