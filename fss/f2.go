package fss

import (
	"fmt"
)

//F20 all
func F20() {
	fmt.Println(f21()())
	fmt.Println(f22(1, 2, 3))
	fmt.Println(f23())
	fmt.Println(f24())
	fmt.Println(f25())
	fmt.Println(f26())
}

//F21 xxxclosure
func f21() func() int {
	x := 1
	return func() int {
		x = x << 1
		return x
	}
}

//F22 xxx multi input
func f22(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func f23() ([]string, [5]int) {
	x := make([]string, 5, 5)
	x = []string{"h", "e", "l", "l", "o"}
	y := [5]int{1, 2, 3, 4, 5}
	return x, y
}

func f24() *int {
	x := new(int)
	return x
}

func f25() []string {
	x := make([]string, 5, 5)
	return x
}

func f26() *int {
	a := 10
	b := &a
	return b
}
