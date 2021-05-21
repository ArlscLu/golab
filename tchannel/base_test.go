package tchannel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func TestSyncWait(t *testing.T) {

	//同步的
	wg.Add(1)
	go printOne()
	wg.Wait()
	printMain()

}

func printOne() {
	for i := 0; i < 3; i++ {
		fmt.Print("z")
	}
	wg.Done()

}
func printMain() {
	fmt.Print("main")
}

func TestSyncChannel(t *testing.T) {
	var notic = make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		firstCalled()
		notic <- true
	}()
	fmt.Print(<-notic)
}

func firstCalled() {
	fmt.Print("right! i am first called!  Wheter how slow i am \n")
}

var shared int

var mu sync.Mutex

func TestAddWithoutSync(t *testing.T) {
	wg.Add(1)
	go func() {
		add10kTimes()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		add20kTimes()
		wg.Done()
	}()
	wg.Wait()
	fmt.Print(shared)
}

func TestAddWithSync(t *testing.T) {
	wg.Add(1)
	go func() {
		add10kTimes()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		add20kTimes()
		wg.Done()
	}()
	wg.Wait()
	fmt.Print(shared)
}

func add10kTimes() {
	mu.Lock()
	defer mu.Unlock()
	for i := 0; i < 2000; i++ {
		shared += 1
	}
}

func add20kTimes() {
	mu.Lock()
	defer mu.Unlock()
	for i := 0; i < 2000; i++ {
		shared += 1
	}
}
