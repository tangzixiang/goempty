package goempty

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

const zero = int64(0)

var value = rand.Int63()

func nothing(interface{}) {}

func BenchmarkEmptyValue_01(b *testing.B) {

	for i := 0; i < b.N; i++ {
		nothing(value == zero)
	}
}

func BenchmarkEmptyValue_02(b *testing.B) {

	for i := 0; i < b.N; i++ {
		nothing(value == int64(0))
	}
}

func BenchmarkEmptyValue_03(b *testing.B) {

	for i := 0; i < b.N; i++ {
		nothing(value&zero == zero)
	}
}

func BenchmarkEmptyValue_04(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IsZero(value)
	}

}

type MyName struct{ name string }

func (m MyName) Empty() bool {
	return m.name == ""
}

func (m MyName) Default() interface{} {
	return "default"
}

func TestIsZero(t *testing.T) {
	if !assert.False(t, IsZero(MyName{"tom"})) {
		return
	}

	if !assert.True(t, IsZero(MyName{""})) {
		return
	}

	if !assert.True(t, IsZero((*uint64)(nil))) {
		return
	}
}
