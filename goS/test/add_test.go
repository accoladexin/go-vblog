package add_test

import (
	"fmt"
	"testing"
)

func add(a, b int) int {
	fmt.Println("add_test.go")
	return a + b
}

func TestAdd(T *testing.T) {
	add(2, 3)
}

func TestAdd2(T *testing.T) {
	add(2, 3)
}
