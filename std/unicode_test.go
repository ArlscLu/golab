package std

import (
	"fmt"
	"testing"
)

var s10 = "ahello永"
var s11 = []byte(s10)
var s12 = []rune(s10)

var raw = `a商\n\r
\t`

// var s13 = '\ud800'

func TestUnicode(t *testing.T) {
	fmt.Println(`string`)
	for i := 0; i < len(s10); i++ {
		fmt.Printf("%d,%[1]U,%[1]c,\n", s10[i])
	}

	fmt.Println(`[]byte`)
	for i := 0; i < len(s11); i++ {
		fmt.Printf("%x,%[1]U,%[1]c,\n", s11[i])
	}

	fmt.Println(`[]rune unicode points`)
	for i := 0; i < len(s12); i++ {
		fmt.Printf("%x,%[1]U,%[1]c,\n", s12[i])
	}

	fmt.Println(`raw`)
	for i := 0; i < len(raw); i++ {
		fmt.Printf("%b,%[1]x,%[1]U,%[1]c,\n", raw[i])
	}

}
