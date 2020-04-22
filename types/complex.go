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
	a = append(a, "!")
	for i, v := range a {
		fmt.Println("index:", i, "value:", v)
	}
	b := a[0:3]
	b[1] = "a"
	fmt.Println(a, b)
}

func Pmap() {
	a := make(map[int]string)
	b := map[int]string{}
	var c map[int]string
	fmt.Println(a, b, c)
}
