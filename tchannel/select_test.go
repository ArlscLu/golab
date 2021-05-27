package tchannel

import (
	"fmt"
	"testing"
	"time"
)

func TestBaseSelect(t *testing.T) {
	var ch1 = make(chan struct{}, 1)
	var ch2 = ch1
	var ch3 = ch1
	go func() {
		// time.Sleep(2 * time.Second)
		ch1 <- struct{}{}
	}()

	select {
	case v := <-ch1:
		fmt.Println(111, v)
	case v := <-ch2:
		fmt.Println(222, v)
	case v := <-ch3:
		fmt.Println(333, v)
		// default:
		// 	fmt.Print("NO new line")
	}
}

func TestSelectSend(t *testing.T) {
	var c = make(chan byte)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Printf("%d\n", <-c)
		}
	}()
	for {
		select {
		case c <- 0:
		case c <- 1:
		}

	}

}
