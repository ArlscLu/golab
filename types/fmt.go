package types

import "fmt"

type s1 struct {
	name string
	age  int
}

var ch chan int

func Run() {
	bowl := make([]interface{}, 10, 10)
	bowl[0] = 10
	bowl[1] = s1{"tom", 18}
	bowl[2] = []int{1, 2, 3, 4}
	bowl[3] = [3]int{100, 101, 102}
	bowl[4] = true
	bowl[5] = "helloworld"
	bowl[6] = 6.7777
	bowl[7] = &bowl[6]
	bowl[8] = 2 << 2
	var shift float32
	shift = 1 << 2
	fmt.Print(bowl)
	fmt.Printf("")
	fmt.Printf("\n%x\n", &bowl[6])
	fmt.Printf("%f", shift)
}
