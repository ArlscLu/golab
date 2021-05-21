package tchannel

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//nil,closed,no closed
func TestStatusChannel(t *testing.T) {
	var ch1 chan struct{}
	<-ch1
	var ch2 = make(chan struct{}, 1)
	close(ch2)
	var ch3 = make(chan struct{}, 1)

	/*
		close, send, receive
	*/
	//close
	// close(ch1) //关闭nil通道
	// close(ch2) //关闭已关闭的通道
	// close(ch3) //成功
	//send
	// ch1 <- struct{}{}
	// ch2 <- struct{}{}
	ch3 <- struct{}{}

	//receive
	// <-ch1
	<-ch2
	<-ch3

	for i := 0; i < 10; i++ {
		fmt.Print(<-ch2)
	}

}

//channel用作 future/promose
func TestPromose(t *testing.T) {
	done := makePromise()
	fmt.Printf("%T", done)
	fmt.Print(<-done)
}

//返回参数隐式转换为 <-chan
func makePromise() <-chan bool {
	r := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second) //massive job
		r <- true
	}()
	return r
}

//channel 用作最快通道
func TestFastestChannel(t *testing.T) {
	ch1 := make(chan int, 1)

	for i := 0; i < cap(ch1); i++ {
		go func(ii int) {
			time.Sleep(time.Duration(rand.Int31()))
			ch1 <- ii
		}(i)
	}
	t.Logf("fastest is %d", <-ch1)
}
