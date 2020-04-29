package types

import (
	"fmt"

	"github.com/arlsclu7/golab/helper"
)

/** print
 */
type (
	//M1 uint8
	M1 uint8
	//MM1 alias for M1 type
	MM1 = M1
	//M2 rune
	M2 rune
	//M3 bool
	M3 bool
	//M4  [3]int
	M4 [3]int
	//M5 structtype
	M5 struct{ X, Y float32 }
	//M6 maptype
	M6 map[int]string
	//M7 complex
	M7 struct {
		x0, x1 M1
		x11    MM1
		x2     M2
		x3     M3
		x4     M4
		x44    []int
		x5     M5
		x6     M6
	}
	M8 *int
	M9 func() M7
)

//Pc 测试type类型创建及赋值
func Pc() {
	var y1 M1 = 255
	var y11 MM1 = 0
	var y2 = '中'
	var y3 = true
	var y4 M4 = [3]int{4, 5, 6}
	var y5 M5 = M5{X: 1.2, Y: 4}
	var y6 = make(map[int]string, 10)
	var y7 M7 = M7{
		x0:  10,
		x1:  11,
		x11: 244,
		x2:  '\n',
		x3:  false,
		x4:  [3]int{222, 333, 444},
		x44: []int{2, 6, 78, 111},
		x5:  M5{2.3, 3.4},
		x6:  make(map[int]string, 7),
	}
	var y8 M8 = new(int)
	var y9 M9 = func() M7 { return y7 }
	var y10 = [3]int{3, 4, 5}
	helper.StartLine()

	fmt.Printf("%T,%T ,%T ,%T,%T,%T,%T,%T,%T,%T,%T\n\n",
		y1, y11, y2, y3, y4, y5, y6, y7, y8, y9, y10)

	helper.StartLine()

	fmt.Println(y1, y11, y2, y3, y4, y5, y6, y7, y8, y9, y10)

	helper.StartLine()

	fmt.Println(y9())

}

//Pn  test literals of all types
func Pn() {
	// x1 := true
	// x2, x22, x222 := 10, -0b111, 0x1f
	x3 := "hello"
	// x4 := [3]string{"hello", "world", "!"}
	// x5 := []bool{true, false, true, false}
	type t6 struct {
		Name string
		Age  uint8
	}
	// x6 := t6{
	// 	Name: "Tom",
	// 	Age:  18,
	// }
	// x7 := new(int)
	// x8 := func() int {
	// 	return 10
	// }
	x10 := map[string]int{
		"j": 12,
		"k": 32,
	}
	X11 := func() string {
		return x3
	}
	x12 := t6{
		Name: "jack",
	}
	x13 := struct {
		Name string
	}{
		Name: "rose",
	}

	fmt.Println(X11(), x10, x12, x13)

}

var s11 string = "helloworld!"

//X11  测试closure  闭包函数，可以使用当前scope的变量
var X11 = func() string {
	return s11
}

func Parray() {
	a := [5]int{1, 2, 3}
	a[4] = len(a)
	b := [...]int{15: 16}
	for i, v := range a {
		fmt.Println("index:", i, "value:", v)
	}
	for i, v := range b {
		fmt.Println("index:", i, "value:", v)
	}
}

func Pslice() {

	a := []string{"h", "e", "l", "l", "o"}
	// a = append(a, "!")
	fmt.Println(&a)
	aa := a[0:1]
	fmt.Println(&aa)
	if &a == &aa {
		println("地址相同")
	}
	for i, v := range a {
		fmt.Println("index:", i, "value:", v)
	}
	b := a[0:3]
	b[1] = "a"
	c := make([]int, 5)
	d := c[1:4]
	d[0] = 2
	fmt.Println(a, b, c, d)
}

func Pmap() {
	a := make(map[int]string)
	// b := map[int]string{}
	var c map[int]string
	var b *int
	a = map[int]string{
		1: "bob",
		2: "mary",
	}
	if a != nil {
		println("&a is not nil")
		println(&a)
		println(a)
		if len(a) == 0 {
			println("长度为0")
		}
	}
	if b == nil {
		println("b is nil")
	}
	if c == nil {
		println("c is nil")
	}
	key := 111
	name, ok := a[key]
	fmt.Println(name, ok)
	if !ok {
		println("不存在", key)
	}

}

func Ptype() {
	type person struct {
		Id       uint8
		Name     string
		Isfemale bool
		Expmonth []int
		Next     *person
		NextId   *uint8
		NextName *string
	}
	marry := person{
		Id:       2,
		Name:     "marry",
		Isfemale: true,
		Expmonth: []int{4, 5, 6},
		Next:     nil,
	}
	jack := person{
		Id:       1,
		Name:     "jack",
		Isfemale: false,
		Expmonth: []int{1, 2, 3},
		Next:     &marry,
		NextId:   &marry.Id,
		NextName: &(marry.Name),
	}
	// fmt.Println(marry)
	// fmt.Println(&marry)
	// fmt.Println(*&marry)
	fmt.Println(jack)
	// fmt.Println(&jack)

}

//Pnil  nil类型tes
func Pnil() {
	var a *int
	var b []int
	c := make([]int, 5)
	fmt.Println(a)
	if b == nil {
		fmt.Println(b)
		fmt.Println("b is nil ")
	}

	d := c
	d[1] = 10
	d = []int{1, 2, 4, 4, 4}
	fmt.Println(c, len(c), cap(c))
	fmt.Println(d, len(d), cap(d))
}

//Pslice2 slice的作用
func Pslice2() {
	a := "helloworld"
	b := "哈罗沃德xxx"
	// c := '哈罗沃德'
	aa := []int{}
	bb := []rune{}
	aaa := []int{}
	bbb := []rune{}
	ccc := []rune("哈罗沃德xxx")
	ddd := []byte(a)
	for k, v := range a {
		aa = append(aa, k)
		bb = append(bb, v)
	}
	for k, v := range b {
		aaa = append(aaa, k)
		bbb = append(bbb, v)
	}
	z := "z"
	fmt.Println(a, b, aa, z, ccc, ddd)
}

//Pall 测试数据类型使用
/**
可为nil  说明存的值为地址
包括了
point,map,slice,channel,func,interface (only)

address,value  （存储）
type，key （系统记录） 决定了怎么使用
address  关联前两者
*/
func Pall() {
	var a *int
	if a == nil {
		println("a is nil")
	}
	var b []int
	if b == nil {
		println("b is nil")
	}
	var c chan int
	if c == nil {
		println("c is nil")
	}
	var d func()
	if d == nil {
		println("d is nil")
	}
	var e error
	if e == nil {
		println("e is nil")
	}
	var f map[int]string
	if f == nil {
		println("f is nil")
	}
	var g int
	if g == 0 {
		println("g is 0")
	}
	type x struct {
		a int
	}
	h := x{a: 10}
	fmt.Println(h)
	b = []int{1, 2, 3, 4, 5}
	bb := b[2:3]
	bb[0] = 111
	z := "z"
	fmt.Println(bb, z)
}
