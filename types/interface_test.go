package types

import (
	"fmt"
	"testing"
)

type Reader interface {
	Read()
}
type Repeatable interface {
	Repeat()
	Init()
}

//iskip接口
type Iskip interface {
	Skip()
	Skip1()
}

type Boy struct {
	Name string
}

func (b Boy) Skip() {
	fmt.Println("i implement skip")
}
func (b Boy) Skip1() {
	fmt.Println("i implement skip1")
}

var tom = &Boy{
	Name: "Tom",
}

func TestA(t *testing.T) {
	needISkip(tom)
}

func needISkip(sk Iskip) {
	sk.Skip()
	sk.Skip1()
}
