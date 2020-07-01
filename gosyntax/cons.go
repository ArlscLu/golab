package gosyntax

import "fmt"

const (
	c0, c00 float32 = iota + 0.1, iota
	c1, c11
	_, _
	c3, c33 //iota = 3
	c4, c44
)

const (
	d0 = 1 << (iota + 1)
	d1 = 1 << (iota + 1)
)

func Prin() {
	fmt.Println(`from pakc`)
	fmt.Println(c0, c00, c1, c11, c3, c33, c4, c44, d0, d1)
}
