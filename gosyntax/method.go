package gosyntax

import "fmt"

type Point struct {
	x, y int
}

type Path []Point

func (p Point) Distance() int {
	return p.x - p.y
}

func (pa Path) Distance() int {
	return len(pa)
}

var pa1 = Path{
	{10, 1},
	{9, 2},
	{8, 3},
	{7, 4},
	{6, 5},
	{5, 4},
}

func Pmethod() {
	fmt.Println(pa1.Distance())
}
