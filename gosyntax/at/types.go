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

type typesl []int

var a [3]int

var slice1 = []int{1, 2, 3}

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

type stcombine struct {
	name string
	info []int
}

var stc1 = stcombine{
	name: "tom",
	info: []int{2, 3},
}

//pointer  =addr + type
func ShowTypes() {
	// return fmt.Sprintln(111)
	i := 5
	var pi *int
	pi = &i
	fmt.Println(pi, &i)
	if pi == &i {
		fmt.Println("是相等的")
	}
	pi2 := pi
	*pi = 100
	fmt.Println(*pi, *&i, *pi2)
	fmt.Println(stc1)
	changestru(stc1)
	fmt.Println(stc1)
}

func changestru(s stcombine) {
	s.name = "jack"
	// s.info[1] = 333
	s.info = append(s.info, []int{7, 8, 9}...)
	// s.info = []int{}
	// change := make([]int, 2, 2)
	// s.info = change
	// fmt.Printf("s.info类型%T    \n change类型%T \n", s.info, change)
	// // s.info = []int{11, 22}

}

//数据类型信息   用于  编译过程中
//类型决定了怎么使用地址
//变量的类型   决定了怎么使用存储的信息
//数据类型 -> 决定了怎么使用内存地址
/*
	编译阶段 变量key   addr_to_get(unknown now)   type
	编译器知道  根据type决定怎么使用addr
*/

var OriginSlice = []int{1, 2, 3}

func NotChangeV(x []int) {
	// var a = [4]int{1, 2, 3, 4}
	// var pi *int
	// someint := 10
	// pi = &someint

	// x = []int{4, 5, 6} //赋值
	x[1] = x[1] * 10
	b := x
	// b = []int{4, 5, 6, 7, 9}
	b = append(b, 5)
	// xx = append(xx, 5)
	// b = make([]int, 3, 3)
	// b[2] = 100
	// *pi = 20
	fmt.Println(x, b, &b, OriginSlice)
	// fmt.Println(pi, someint, *pi)
}

func ChangeV(x *[]int) {
	*x = append(*x, 300)
	fmt.Println(x, *x, OriginSlice)
}
