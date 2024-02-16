package examples

import (
	"fmt"
	"testing"
)

func TestGoHasGenerics(t *testing.T) {
	stringSlice := []string{"one", "two", "three"}
	integerSlice := []int{1, 2, 3}
	//stringSlice = append(stringSlice, 4)
	fmt.Println(stringSlice)
	fmt.Println(integerSlice)
}
