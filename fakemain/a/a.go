package main

import (
	"fmt"
	"time"
)

var i = 0

func main() {
	var ch = make(chan struct{})
	go func() {
		//只阻塞一次
		// <-ch
		for {
			<-ch
			fmt.Println("i is", i)
			time.Sleep(time.Second)
		}
	}()

	for {
		i += 1
		ch <- struct{}{}
		// fmt.Println("我爱你 ", i)
	}
}
