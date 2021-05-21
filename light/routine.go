package main

import (
	"fmt"
	"time"
)

func gopher() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("%s\n", "hello")
	}

}

func mainer() {
	go gopher()
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("%s\n", "from main")
	}
}

func main() {
	mainer()
	var a []int = nil
	c := make(chan struct{})

	go func() {
		a = make([]int, 3)
		c <- struct{}{}
	}()

	<-c
	a[0], a[1], a[2] = 0, 1, 2 // 绝不会造成恐慌
}
