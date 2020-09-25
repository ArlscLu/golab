package std

import (
	"fmt"
	"strings"
	"testing"
)

var s10 = "又ahello永"
var s11 = []byte(s10)
var s12 = []rune(s10)

var raw = `an\r
\t`

// var s13 = '\ud800'

var r1 rune = '"'

func TestUnicode(t *testing.T) {
	fmt.Println(`string`)
	for i := 0; i < len(s10); i++ {
		fmt.Printf("%d,%[1]U,%[1]c,\n", s10[i])
	}

	fmt.Println(`[]byte s11`)
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
	fmt.Println(`rune`)
	fmt.Printf("%+v,%[1]b,%[1]T,%[1]q\n", r1)
	fmt.Print(strings.IndexAny(s10, "永"))
}
