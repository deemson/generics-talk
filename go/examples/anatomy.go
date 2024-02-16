package examples

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type GenericStruct[T any] struct {
	Value T
}

func GenericFunction[T any](value T) T {
	return value
}

func TestShow(t *testing.T) {
	s := GenericStruct[int]{Value: 42}
	fmt.Println(s)
	assert.Equal(t, "hello", GenericFunction[string]("hello"))
	assert.Equal(t, "hello", GenericFunction("hello"))
}
