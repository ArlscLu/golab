package std

import (
	"fmt"
	"time"
)

var ct = make(chan bool)

//10次
var addTime = make(chan int, 10)

func RunChan() {
	fmt.Println("wait  fro  5 seconds")
	<-time.After(time.Second * 5)
	fmt.Println("end of  5 seconds")
	go func() {
		time.Sleep(5 * time.Second)
		ct <- true
	}()
	for {
		select {
		case <-ct:
			fmt.Println("start")
		default:
			fmt.Println("default ")
		}
	}

}

func sendDone() {
	fmt.Println("wait for 10 seconds")
	time.Sleep(time.Second * 10)
	fmt.Println("waited for 10 seconds")
	ct <- true
}
func timeout() {
	for i := 0; i < 10; i++ {
		// time.Sleep(time.Second * 1)
		<-addTime
	}
}

func Main() {
	alpaTable := []string{"a", "b", "c", "d"}
	for k, v := range alpaTable {
		go func(k int, v string) {
			k++
			for i := 0; i < k; i++ {
				fmt.Println(v)
			}
		}(k, v)
	}
	for i := 0; i < 10; i++ {
		addTime <- i
	}
	go sendDone()
	go timeout()
	for {
		select {
		case <-ct:
			fmt.Println("receive")
		case addTime <- 1:
			fmt.Println("可以insert了")
		}
	}
}

func Ds() {
	var ch = make(chan int)
	go func() {
		// time.Sleep(time.Second * 3)
		ch <- 1
	}()
	<-ch
	fmt.Println("end ")
}
