package types

import (
	"fmt"
)

//B bbb
var B bool = true

//Num num
var Num int = 13

//X string
var X string = "U+000A"

//Float0 test
var Float0 float32 = 12.23222

const (
	C0 = 1<<iota - 1
	C1 = 1 << iota
	C2 = 1 << iota
	C4 = '中'
)

//Rst rst
type Rst struct {
	R1 int
	R2,
	R22 float32
	R3 []int
	R4 []string
	R5 bool
	R6 [5]string
	R7 rune
}

//Base printrst
func Base() Rst {
	var x = Rst{
		R1:  1,
		R2:  3.3,
		R22: 4.4,
		R3:  []int{1, 2, 3},
		R7:  '中'}
	return x
}

func Pbase() {
	type t1 struct {
		name string
		age  int
	}
	a := new(int)
	fmt.Println("a", a)
	fmt.Println("&a", &a)
	fmt.Println("*a", *a)
	aa := make(map[int]string)
	fmt.Println("aa", aa)
	fmt.Println("&aa", aa)
	bb := 10
	fmt.Println("bb", bb)
	fmt.Println("&bb", &bb)
	cc := make([]int, 5)
	fmt.Println("cc", cc)
	fmt.Println("&cc", &cc)
	dd := t1{
		name: "Tom",
		age:  18,
	}
	fmt.Println("dd", dd)
	fmt.Println("&dd", &dd)
	b, c, d := a, a, a
	*a = 10
	fmt.Println(a, b, c, d, *a, *b, *c, *d, &b)
}
