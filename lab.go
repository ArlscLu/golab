package main

import (
	"fmt"

	"github.com/arlsclu7/golab/types"
)

func newInt() *int {
	return new(int)
}
func main() {

	fmt.Println(types.P1())

	for i := 1; i < 10; i++ {
		fmt.Printf("%x \n", i)
	}

	ptr := new(float32)
	a := "mystring"
	b := &a
	fmt.Println(b)
	fmt.Println("ptr address", b)
	fmt.Println("ptr value", *b)
	fmt.Println("b adress", ptr)
	fmt.Println("b value", *ptr)
}
