package gosyntax

import (
	"fmt"
)

func Ppanic() {
	go func() {
		fmt.Println(`i am from routine`)
	}()
	go func() {
		fmt.Println(`i am another`)
	}()
}
