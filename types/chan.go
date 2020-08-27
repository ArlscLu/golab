package types

import (
	"fmt"
	"time"
)

var stop = make(chan bool)

func RunChan() {

	go rrunChan()
	time.Sleep(10 * time.Second)
	fmt.Println("可以停止监控了")
	stop <- true
}

func rrunChan() {
	fmt.Println("开始r")
	for {
		select {
		case <-stop:
			fmt.Println("结束监控")
			return
		default:
			fmt.Println("继续监控")
			time.Sleep(2 * time.Second)
		}
	}
}
