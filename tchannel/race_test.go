package tchannel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var share int
var wg1 sync.WaitGroup

// var mu1 sync.Mutex

func TestRace(t *testing.T) {

	wg1.Add(3)
	go add()
	go add()
	go add()
	wg1.Wait()
	t.Logf("累计后得到值%d", share)
}

func add() {
	// mu1.Lock()
	// defer mu1.Unlock()
	for i := 0; i < 1000; i++ {
		share = share + 1
	}
	wg1.Done()
}

func longTime2() {
	time.Sleep(2 * time.Second)
	fmt.Print("I lost 2 seconds")
	wg1.Done()
}
func longTime1() {
	time.Sleep(1 * time.Second)
	fmt.Print("I lost 1 seconds")
	wg1.Done()
}

func shortTime() {
	// time.Sleep(1 * time.Second)
	fmt.Print("fast\n")
}

func TestNotWait(t *testing.T) {
	wg1.Add(3)
	shortTime()
	shortTime()
	shortTime()
	go longTime1()
	longTime2()

	go longTime2()
	shortTime()
	shortTime()
	shortTime()
	wg1.Wait()

}

func TestSingleInst(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go printErrorIfZero()
	}
}

var inst int

func singleInst() int {

	if inst == 0 {
		inst = 1
	}
	return inst
}
func printErrorIfZero() {
	if singleInst() == 0 {
		panic("xxxxx")
	}
}

func TestChannelRace(t *testing.T) {
	c := make(chan int) // 一个非缓冲通道
	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
		// <-ch        // 此操作编译不通过
		ch <- x * x // 阻塞在此，直到发送的值被接收
	}(c, 3)
	done := make(chan struct{})
	go func(ch <-chan int) {
		n := <-ch      // 阻塞在此，直到有值发送到c
		fmt.Println(n) // 9
		// ch <- 123   // 此操作编译不通过
		time.Sleep(time.Second)
		done <- struct{}{}
	}(c)
	<-done // 阻塞在此，直到有值发送到done
	fmt.Println("bye")
}

func TestBall(t *testing.T) {
	var ball = make(chan string)
	kickBall := func(playerName string) {
		for {
			fmt.Print(<-ball, "传球", "\n")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}
	go kickBall("张三")
	go kickBall("李四")
	go kickBall("王二麻子")
	go kickBall("刘大")
	ball <- "裁判"    // 开球
	var c chan bool // 一个零值nil通道
	<-c             // 永久阻塞在此
}

func TestGlobalVar(t *testing.T) {

	var i = 0
	go func() {
		for {
			fmt.Println("i is", i)
			time.Sleep(time.Second)
		}
	}()

	for {
		i += 1
		//fmt.Println("我爱你 ", i)
	}
}
