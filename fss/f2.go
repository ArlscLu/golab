package fss

import (
	"fmt"
)

//F21 xxxclosure
func F21() func() int {
	x := 1
	return func() int {
		x = x << 1
		return x
	}
}

//F22 xxx multi input
func F22(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

//F20 all
func F20() {
	fmt.Println("***来自fss.f2***")
	fmt.Println(
		F21()(),
		F22(1, 2, 3),
	)
}
