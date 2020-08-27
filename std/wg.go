package std

import (
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func RunWg() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1号完成")
		wg.Done()
	}()
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("2号完成")
		wg.Done()
	}()
	wg.Wait()
	logrus.Debug("1")
	fmt.Println("all done")
}
