package self

import (
	"fmt"
	"testing"
)

var ss = "hello永"

func TestUnicode(t *testing.T) {
	fmt.Printf("%s\n", ss)
	fmt.Printf("%c\n", []byte(ss))
	b := []byte(ss)
	for i := 0; i < len(b); i++ {
		fmt.Printf("%b\n", b[i])
	}
}
