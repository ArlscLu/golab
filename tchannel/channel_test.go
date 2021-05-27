package tchannel

import (
	"fmt"
	"math/rand"
	"sync"
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
	<-done
	t.Log("all done")
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
func TestAfterTime(t *testing.T) {
	t.Log("yy")
	<-time.After(2 * time.Second)
	t.Log("xx")
}

func TestTickTime(t *testing.T) {
	c := time.Tick(2 * time.Second)
	for next := range c {
		fmt.Printf("%v %s\n", next, "exist once")
	}
}
func TestShared(t *testing.T) {
	var sh int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {

		for i := 0; i < 10000; i++ {
			sh++
		}
		wg.Done()

	}()
	wg.Wait()
	for i := 0; i < 10000; i++ {
		sh++
	}

	time.Sleep(1 * time.Second)
	t.Log(sh)
}

func TestTranslate(t *testing.T) {

	var sh1 int = 3000
	var sh2 int = 7000
	var mu sync.Mutex
	//trans
	trans := func() {
		mu.Lock()
		sh1 -= 2000
		sh2 += 2000
		mu.Unlock()
	}
	for i := 0; i < 10000; i++ {
		go trans()
	}
	t.Log(sh1 + sh2)
	t.Log(sh1 + sh2)
}
