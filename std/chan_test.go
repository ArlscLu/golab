package std

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

var chanSum int

func TestChan1(t *testing.T) {
	c1 := make(chan bool)
	f, err := os.OpenFile("std.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	l := logrus.New()
	l.SetOutput(f)
	l.SetLevel(logrus.DebugLevel)
	if err != nil {
		t.Log(err)
	}
	go func() {
		for i := 1; i <= 1000; i++ {
			chanSum += i
		}
		c1 <- true
	}()
	<-c1
	fmt.Println("end of file")
}

func TestChansDanger(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestChansDangerWithMute(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// 协程安全Demo
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex    // 互斥锁
	var wg sync.WaitGroup // 等待队列
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1) // 加个任务
		go func() {
			defer func() {
				mut.Unlock() //函数调用完成后：解锁，保证协程安全
			}()
			mut.Lock() // 函数将要调用前：加锁，保证协程安全
			counter++
			wg.Done() // 做完任务
		}()
	}
	wg.Wait() //等待所有任务执行完毕
	t.Logf("counter = %d", counter)
}

func TestRoutineOrder(t *testing.T) {

}
