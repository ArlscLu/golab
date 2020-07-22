package at

import (
	"fmt"
	"time"
)

/*

The language predeclares certain type names.
Others are introduced with type declarations.
Composite types—
**array, struct, pointer,
	function, interface, slice,
	map, and channel types—
may be constructed using type literals.


Type      = TypeName | TypeLit | "(" Type ")" .
TypeName  = identifier | QualifiedIdent .
TypeLit   = ArrayType | StructType | PointerType | FunctionType | InterfaceType |
	    SliceType | MapType | ChannelType .
*/

/*
if !Typelit {
	then (predeclared)TypeName in (
		bool byte complex64 complex128 error float32 float64
		int int8 int16 int32 int64 rune string
		uint uint8 uint16 uint32 uint64 uintptr
	)
	eg.   type uint8  uint8
		  type rune int32
} else {
	is TypeName | TypeLit
}

*/
type ebool bool

//ArrayType   "type"+ identifier + typelit
type at [3]int

//StructType
type st struct {
	name string
}

//PointType
type pt *int

//FunctionType
type ft func(a []byte) int

//InterfaceType
type it interface {
	run()
}

//SliceType
type st1 []rune

//MapType
type mt map[uint8]string

//ChannelType
type ct <-chan int
type ct2 chan<- int
type ct3 chan int

var a [3]int

var b []string

var c it

var d = mt{}

var e = make(ct, 3)
var e2 = make(ct2, 3)
var e3 = make(ct3, 3)

var ee = make([]int, 4, 6)

var ex = []int{4, 5, 6}

var eee = make(map[int]string, 10)

// func changeE(c ct) {
// 	2 <- c
// }

func changeEE(s []int) {
	// s = append(s, []int{1, 2, 3}...)
	s[0] += 10
	s[1] += 10
	s[2] += 10
	//此时指向新array地址
	// s = append(s, []int{1, 2, 3, 4}...)
	//此时未指向新array地址,因为caps足够
	s = append(s, []int{1, 2}...)
	//差异点
	s[3] += 11
}

// type SomeTest struct {
// 	Name string
// 	Age  byte
// }

func Show() {
	e2 <- 1
	e2 <- 2
	e2 <- 3
	e3 <- 1
	e3 <- 2
	e3 <- 3

	close(e2)
	_ = <-e3
	fmt.Println(len(e3))

	_ = <-e3
	fmt.Println(len(e3))
	go func() {
		if len(e3) == 1 {
			fmt.Println("it is my turn")
		}
	}()
	// _ = <-e3

	fmt.Println(len(e3))
	fmt.Println("************")
	fmt.Println(ee)
	// fmt.Println(changeEE(ex))
	changeEE(ee)
	fmt.Println(1)
	fmt.Println(c, d, e, ee, eee, ee)
	time.Sleep(10)

}
