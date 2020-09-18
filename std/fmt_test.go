package std_test

import (
	"fmt"
	"testing"
)

var bool1 bool = true

var mint1 int = 97

var mfloat1 float32 = 0.56

var mstr1 string = "ahello中"

var mslice1 []byte = []byte("abc")

var mpint1 *int = &mint1

var mstruct1 = struct {
	Name, Name2 string
	Age         int
}{
	Name: "Tom",
	Age:  18,
}

func TestFmtVerb(t *testing.T) {
	// base 2
	fmt.Printf("%b\n", mint1)
	// the character represented by the corresponding Unicode code point
	fmt.Printf("%c\n", mint1)
	// base 10
	fmt.Printf("%d\n", mint1)
	// base 8
	fmt.Printf("%o\n", mint1)
	// a single-quoted character literal safely escaped with Go syntax.
	fmt.Printf("%q\n", mint1)
	// base 16, with lower-case letters for a-f
	fmt.Printf("%x\n", mint1)
	// base 16, with lower-case letters for A-F
	fmt.Printf("%X\n", mint1)
	// Unicode format: U+1234; same as "U+%04X"
	fmt.Printf("%U\n", mint1)

	fmt.Printf("%x", mstr1)
	fmt.Printf("%d", 11)

	//small is beautiful
}
