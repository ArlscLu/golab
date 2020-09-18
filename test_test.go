package main

import (
	"fmt"
	"testing"
)

var expect = []struct {
	in  int
	out bool
}{
	{1, false},
	{2, true},
	{3, false},
	{4, false},
	{5, true},
}

func TestA(t *testing.T) {
	fmt.Printf("%d\n", 123)
}

func IsEven(i int) (isEven bool) {
	isEven = false
	if i%2 == 0 {
		isEven = true
	}
	return
}
func TestCompare(t *testing.T) {
	for _, v := range expect {
		real := IsEven(v.in)
		if real != v.out {
			fmt.Println("结果不符合预期", v)
		}
	}

}

//benchmark testing
func BenchmarkTest(b *testing.B) {
	fmt.Println("ewew")
}

// func ExamplePrintln() {
// 	fmt.Println("The output of\nthis example.")
// 	// Output: The output of111
// 	// this example.
// }
func ExampleIsEven() {
	// fmt.Println(IsEven(10))
	//Output:true
	fmt.Println(IsEven(10), "\n", IsEven(11))
	//Unordered output:true
	//false

}
