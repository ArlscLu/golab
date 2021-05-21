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
	wg.Done()
}
func longTime1() {
	time.Sleep(1 * time.Second)
	fmt.Print("I lost 1 seconds")
	wg.Done()
}

func shortTime() {
	// time.Sleep(1 * time.Second)
	fmt.Print("fast\n")
}

func TestNotWait(t *testing.T) {
	wg.Add(3)
	shortTime()
	shortTime()
	shortTime()
	go longTime1()
	longTime2()

	go longTime2()
	shortTime()
	shortTime()
	shortTime()
	wg.Wait()

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
