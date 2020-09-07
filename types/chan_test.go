package types

import (
	"testing"
	"time"
)

type tunnel chan int

func TestBase(t *testing.T) {
	ch1 := make(tunnel)
	go func() {
		// time.Sleep(time.Second * 3)
		ch1 <- 1
	}()
	time.Sleep(time.Second * 3)
	<-ch1
	t.Logf("%s", "helloworld")
}
