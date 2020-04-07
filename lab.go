package main

import (
	"fmt"

	"github.com/arlsclu7/golab/types"
)

func main() {
	fmt.Println(types.B)
	fmt.Println(types.Num)
	fmt.Printf("%d \n", types.Num)
	fmt.Printf("%b \n", types.Num)
	fmt.Printf("%o \n", types.Num)
	fmt.Printf("%x,%X \n", types.Num, types.Num)
	fmt.Printf("%x\n", types.X)
	fmt.Printf("%U", "上海")
	fmt.Print("111111\n")
	a := "中国"
	fmt.Println(len(a))
	for i := 0; i < len(a); i++ {
		fmt.Printf("%x ", a[i])
	}
	fmt.Println()
	for i := 0; i < len(a); i++ {
		fmt.Printf("%c ", a[i])
	}
	s := "中国"
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%x", r[i])
	}
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i])
	}
}
