package std

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCtx(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				fmt.Println("do once +++")
				// time.Sleep(1e9) //solid when not set
				select {
				case <-ctx.Done():
					fmt.Println("ctx done")
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		//之前执行了  do once *2
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	// for x := range newrange(2) {
	// 	fmt.Println(x)
	// 	if x == 2 {
	// 		break
	// 	}
	// }
}

// func newrange(i int) <-chan int {
// 	ch := make(chan int)
// 	go func() {
// 		for {
// 			ch <- i
// 			i++
// 		}
// 	}()
// 	return ch
// }

func TestRange(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case ch <- n:
					n++
				case <-ctx.Done():
					t.Logf("%+v", <-ctx.Done())
					fmt.Println("ctxx done")
					return
				}
			}
		}()
		return ch
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

}

func TestZs(t *testing.T) {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go func() {
		for i := 1; i <= 300; i++ {
			ch1 <- i
		}
	}()
	go func() {
		for j := 1; j <= 400; j++ {
			ch2 <- j
		}
	}()
	go func() {
		for j := 1; j <= 400; j++ {
			ch2 <- j
		}
	}()
	time.Sleep(2 * 1e9)
	for x := 0; x <= 10; x++ {
		//select 如果没值则阻塞
		select {
		case <-make(chan int, 1):
			fmt.Println("self chan int")
		case i := <-ch1:
			fmt.Println("ch1输出", i)
		case y := <-func() chan int {
			fmt.Println("xxxxxxxx")
			ch := make(chan int)
			go func() {
				ch <- 1
			}()
			return ch
		}():
			fmt.Println("func 输出", y)
		case <-ch2:
			fmt.Println("ch2输出")
		case <-make(chan []int, 1):
			fmt.Println("self chan []int")

			// default:
			// 	fmt.Println("default 输出")
		}
	}

}
