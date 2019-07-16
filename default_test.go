package goempty

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {

	if !assert.Equal(t, "default", Default(MyName{})) {
		return
	}

	if assert.Equal(t, "default", Default(MyName{"tom"})) {
		return
	}
}
