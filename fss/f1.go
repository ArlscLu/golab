package fss

import (
	"fmt"

	"github.com/arlsclu7/golab/types"
)

//F10 all
func F10() {
	fmt.Println("***来自fss.f1***")
	fmt.Println(
		F11(),
		"\n",
		F12(),
		"\n",
		F13(),
	)
}

//F11 xxxx
func F11() []types.M4 {
	a := make([]types.M4, 10, 100)
	i := 0
	for i < 10 {
		a[i] = types.M4{
			i + 4,
			i + 5,
			i + 6,
		}
		i++
	}
	return a
}

//F12 xxxx
func F12() []types.M5 {
	a := make([]types.M5, 10, 100)
	i := 0
	for i < 10 {
		a[i] = types.M5{
			X: float32(i) + 4.1,
			Y: float32(i) + 5.1,
		}
		i++
	}
	return a
}

//F13 xxx
func F13() []int {
	a := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	return a
}
