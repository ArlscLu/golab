package main

import (
	"fmt"
	"math"
)

func swap(x, y string) (string, string) {
	return y, x
}
func myprint(x string) {
	fmt.Println(x)
}
func my2print(x float64) {
	fmt.Println(x)
}

func main() {
	if v := math.Pow(2, 10); v > 99 {
		my2print(v)
	}
	var x int = 10
	if x > 4 {
		myprint("大于4")
	} else {
		myprint("小于4")
	}
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	sum2 := 1
	for sum2 < 10 {
		sum2 += sum2
	}
	sum3 := 15
	for sum3 > 10 {
		sum3--
		fmt.Println(sum3)
	}

	var (
		a uint = 1<<64 - 1 //把1移动到第i位 （i初始位置为0）
		b bool = true
	)
	var (
		aa int
		bb float32
		cc bool
		dd string
	)
	var i, j int = 0, 255
	fmt.Printf("type is %T value is %v\n", i, i)
	fmt.Printf("type is %T value is %v\n", j, j)
	fmt.Printf("type is %T value is %v\n", a, a)
	fmt.Printf("type is %T value is %v\n", b, b)
	fmt.Printf("%v,%v,%v,%q", aa, bb, cc, dd)

}
