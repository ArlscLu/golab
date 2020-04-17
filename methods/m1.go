package methods

import (
	"fmt"
)

func M1() {
	fmt.Println("***来自methods.M1()")
	fmt.Println(p1)
	fmt.Println(f2(p2))
	fmt.Println(p1.Distance(16))
}

type Point struct {
	x, y float32
}

var p1 Point = Point{
	3.2,
	4.2,
}
var p2 *Point = &p1

func f2(p *Point) Point {
	*p = Point{4.4, 4.4}
	return *p
}

func (p Point) Distance(n uint8) float32 {
	return float32(n) * 1.6
}
