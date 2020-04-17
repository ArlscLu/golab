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
	var y6 M6 = make(map[int]string, 10)
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

	helper.StartLine()

	fmt.Printf("%T,%T ,%T ,%T,%T,%T,%T,%T,%T,%T\n\n", y1, y11, y2, y3, y4, y5, y6, y7, y8, y9)

	helper.StartLine()

	fmt.Println(y1, y11, y2, y3, y4, y5, y6, y7, y8, y9)

	helper.StartLine()

	fmt.Println(y9())
}
