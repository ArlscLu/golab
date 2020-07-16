package gosyntax

import "fmt"

type weekday uint8

const (
	Monday weekday = iota<<2 + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func Pself() {
	fmt.Println(Monday)
}
