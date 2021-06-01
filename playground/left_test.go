package playground

import (
	"fmt"
	"testing"
	"time"
)

func TestLeftTime(t *testing.T) {
	var done = make(chan struct{})
	everyPeriod := func() {
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}
	target := time.Date(2021, 5, 28, 18, 30, 0, 0, time.Local)
	for {
		go everyPeriod()
		fmt.Printf("%v过了5秒钟,距离还有%v\n", <-done, time.Until(target).Minutes())
	}
}

func TestNum(t *testing.T) {
}
