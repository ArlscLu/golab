package std

import (
	"fmt"
	"testing"
)

var s1 = "hello永"
var s2 = []rune(s1)

func TestUnicode(t *testing.T) {
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%d,%b\n", i, s1[i])
	}
	for i, v := range s1 {
		fmt.Printf("%d,%U,%c,\n", i, v, v)
	}
}
