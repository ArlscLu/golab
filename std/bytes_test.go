package std

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBytes(t *testing.T) {
	a := []byte("hello")
	b := []byte("world")
	fmt.Println(bytes.Equal(a, b))   // false
	fmt.Println(bytes.Compare(a, b)) // -1
	fmt.Println(bytes.Compare(b, a)) // 1
	a = []byte("hello ϕ")
	b = []byte("Hello Φ")
	fmt.Println(bytes.EqualFold(a, b)) // true
}
