package std

import (
	"fmt"
	"time"
)

var ct chan bool

func RunChan() {

	go func() {
		time.Sleep(5 * time.Second)
		ct <- true
	}()
	for {
		select {
		case <-ct:
			fmt.Println("start")
		}
	}

}
