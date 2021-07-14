package main

import "fmt"

func main() {

	var ch1Done = make(chan struct{})
	var ch2Done = make(chan struct{})
	var ch3Done = make(chan struct{})
	var ch = make(chan struct{})

	go func() {
		for i := 0; i < 100; i++ {
			<-ch3Done
			fmt.Println("wolf")
			ch1Done <- struct{}{}
		}

	}()

	go func() {
		for i := 0; i < 3; i++ {
			<-ch1Done
			fmt.Println("sheep")
			ch2Done <- struct{}{}
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			<-ch2Done
			fmt.Println("grass")
			if i == 2 {
				ch <- struct{}{}
				ch <- struct{}{}
			}
			ch3Done <- struct{}{}

		}
	}()
	ch3Done <- struct{}{}
	<-ch
}
