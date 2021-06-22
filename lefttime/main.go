package main

import (
	"fmt"
	"time"
)

func main() {
	var oneSec = make(chan time.Time)
	go func() {
		for {
			time.Sleep(5 * time.Second)
			oneSec <- time.Now()
		}

	}()
	for {
		v := <-oneSec
		n := v.Format(time.Stamp)
		fmt.Println(n)
	}
}
