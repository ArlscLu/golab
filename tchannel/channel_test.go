package tchannel

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestMakeChannel(t *testing.T) {
	// var ch chan int
	var wg sync.WaitGroup
	wg.Add(3)
	go printOnce(&wg)
	wg.Wait()
	fmt.Print("ALL DONE")
}

func printOnce(wg *sync.WaitGroup) {
	for {
		wg.Done()
		fmt.Print(111, "\n")
	}
}

func TestWaitUntilNotify(t *testing.T) {
	done := make(chan struct{})
	go func() {
		//花费时间的io操作
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()
	fmt.Print("111\n")
	fmt.Print("111\n")
	<-done
}

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

func TestSort(t *testing.T) {
	ori := []int64{1, -1, 15, 23, 29, 1, 22, -3, -9}
	sort.Slice(ori, func(i int, j int) bool {
		return true
	})
	t.Log(ori)
}
