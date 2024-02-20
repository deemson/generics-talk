package dao

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func UnmarshalJSON[T any](data []byte) (T, error) {
	var value T
	err := json.Unmarshal(data, &value)
	return value, err
}

type Struct struct {
	Field string
}

func TestStructUnmarshal(t *testing.T) {
	s, err := UnmarshalJSON[Stuct]([]byte(`{"Field": "hello world"}`))
	assert.NoError(t, err)
	assert.Equal(t, s.Field, "hello world")
}

func Sort[T any](slice []T, less func(t1 T, t2 T) bool) {
	sort.Slice(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
}

func TestSort(t *testing.T) {
	slice := []Struct{{"c"}, {Field: "a"}, {"b"}}
	Sort(slice, func(t1 Struct, t2 Struct) bool {
		return t1.Field < t2.Field
	})
	assert.Equal(t, []Struct{{"a"}, {Field: "b"}, {"c"}}, slice)
}
