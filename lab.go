package main

import (
	"fmt"

	"github.com/arlsclu7/golab/helper"
	"github.com/arlsclu7/golab/types"
)

func main() {
	helper.StartLine()
	d := []int{}
	fmt.Println(cap(d))
	for i := 0; i < 16; i++ {
		d = append(d, i)
		fmt.Print(cap(d), " ")
	}
	fmt.Println(types.E1)
	helper.DivideLine()
	a := make([]int, 5, 6)
	b := a[0:4]
	a = append(a, 1)
	a[1] = 5
	fmt.Println(b)
	// [0 0 0 0]
	fmt.Println(a)
	// [0 5 0 0 0 1]</pre>
	// var a = types.Float0
	// var b byte = '\v'
	var c rune = '中'
	var mystr1 string = "hello中国"
	var mystr2 [5]byte = [5]byte{104, 101, 65, 108, 111}
	fmt.Println(types.Arr0)
	fmt.Println(types.Slice0[1:5])
	fmt.Println(types.Slice1)
	fmt.Printf("值为%d,类型为%T\n", types.Slice0[0:2], types.Slice0[0:2])
	println(len(mystr1))
	println(len(mystr2))
	helper.DivideLine()
	var x []int = append(types.Slice2, []int{4, 5, 6}...)
	var y []int = append(append([]int{-1, -2, -3}, types.Slice2...), types.Slice2[:3]...)
	var z []int = types.Slice0[0:2:5]
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%s \n", mystr1)
	fmt.Printf("%s \n", mystr2)
	fmt.Printf("%.3f\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%U\n", c)
	fmt.Println(c)
	helper.NewLine()
}
