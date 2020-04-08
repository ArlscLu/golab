package main

import (
	"fmt"

	"github.com/arlsclu7/golab/types"
)

func main() {

	var a = 1
	a = a << 8

	fmt.Printf("%b", a)
	fmt.Printf("%d", a-1)

	fmt.Println()
	var b = 10
	var c = &b
	fmt.Println(c)

	fmt.Println(types.B)
	fmt.Println(types.Num)
	fmt.Printf("%d \n", types.Num)
	fmt.Printf("%b \n", types.Num)
	fmt.Printf("%o \n", types.Num)
	fmt.Printf("%x,%X \n", types.Num, types.Num)
	fmt.Printf("%x\n", types.X)

}
