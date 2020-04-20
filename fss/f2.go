package fss

import (
	"fmt"
)

//F20 all
func F20() {
	// fmt.Println(f21()())
	// fmt.Println(f22(1, 2, 3))
	// fmt.Println(f23())
	// fmt.Println(f24())
	// fmt.Println(f25())
	// f26()
	// f27(y)
	// fmt.Println("函数外的x值", x)
	// f28()
	// f29(z)
	// f290(&t11)
	// f291(&t22)
	// f292()
	f293()
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

func f26() {
	a := 10 //int
	b := &a
	*b = 11
	fmt.Println(a, b, *b)
	fmt.Printf("%T,%T,%T\n", a, b, *b)
}

var x = 10
var y = &x

//f27 测试值（指针，值），引用  。引用只能用*xx的形式，而不能使用y= *xx形式
//系统使用make()返回的是  ，只能系统完成
func f27(xx *int) {
	// y := xx
	// *y++
	y := *xx
	// y++
	// y++
	*xx++
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	// helper.NewLine()
	fmt.Println(*xx)
	fmt.Printf("%T\n", *xx)
	// fmt.Println("value of y:", y)
}

func f28() {
	x := make([]int, 5)
	y := [...]int{4, 5, 6}
	z := make(map[int]string)
	zz := z //引用不能赋值
	zz = map[int]string{
		1: "hello",
		2: "world",
		3: "!",
	}
	zzz := 65
	fmt.Println(x, y, z, zz)
	fmt.Printf("%T,%T,%T,%T,%c \n", x, y, z, zz, zzz)
}

var z = make(map[int]string)

func f29(zz map[int]string) {
	zz = map[int]string{
		1: "hellos",
		2: "worlds",
	}
	ff := map[int]string{
		1: "hello",
		2: "world",
	}
	xx := 100
	fmt.Printf("xx类型%T\n", xx)
	fmt.Println("值为", xx)
	fmt.Printf("&xx类型%T\n", &xx)
	fmt.Println("值为", &xx)
	fmt.Printf("&zz类型%T\n", &zz)
	fmt.Println("值为", &zz)
	// if zz == nil {
	// 	fmt.Println("空指针map")
	// }
	fmt.Printf("zz类型%T\n", zz)
	fmt.Println("值为", zz)
	fmt.Printf("ff类型%T\n", ff)
	fmt.Println("值为", ff)
}

/**
00000|111111111

类型/值  |   类型/值
*int/0xcbfa0124   |  int/10
*map[int]string/&map[1:hellos 2 worlds]	 |   map[int]string/map[1:hellos 2 worlds]
*/

type T1 struct {
	Name string
	Age  uint8
}

var t11 = T1{
	Name: "Tom",
	Age:  18,
}
var t22 = 10

func f290(t *T1) {
	fmt.Printf("类型%T\n", t)
	t.Name = "Rose"
	fmt.Println("值", t)
	fmt.Printf("t11类型%T\n", t11)
	fmt.Println("t11值", t11)
}
func f291(x *int) {
	fmt.Printf("类型%T\n", x)
	*x = 11
	fmt.Println("值", *x)
}

func f292() {
	y := 10
	yy := &y
	fmt.Printf("类型%T\n", yy)
	fmt.Println("值", yy)
}

func f293() {
	x := 10
	y := T1{
		Name: "Jask",
	}
	z = make(map[int]string, 5)
	zz := map[int]string{
		1: "da",
		2: "ees",
	}
	a := new(int)
	xx := new(uint8)
	*xx = 120
	fmt.Printf(" x 类型%T\n", x)
	fmt.Println(x)
	fmt.Printf("&x 类型%T\n", &x)
	fmt.Println(&x)
	fmt.Printf("y 类型%T\n", y)
	fmt.Println(y)
	fmt.Printf("&y 类型%T\n", &y)
	fmt.Println(&y)
	fmt.Printf("z 类型%T\n", z)
	fmt.Println(z)
	fmt.Printf("&z 类型%T\n", &z)
	fmt.Println(&z)
	fmt.Printf("zz 类型%T\n", zz)
	fmt.Println(zz)
	fmt.Printf("&zz 类型%T\n", &zz)
	fmt.Println(&zz)
	fmt.Printf("a 类型%T\n", a)
	fmt.Println(a)
	fmt.Printf("*a 类型%T\n", *a)
	fmt.Println(*a)
	fmt.Printf("xx 类型%T\n", xx)
	fmt.Println(xx)
	fmt.Printf("*xx 类型%T\n", *xx)
	fmt.Println(*xx)
}
