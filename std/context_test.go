package std

import (
	"context"
	"fmt"
	"testing"
)

func TestCtx(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
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
