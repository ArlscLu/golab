package gosyntax

import (
	"container/list"
	"fmt"
)

func ShowList() {
	f1 := func() {
		fmt.Println("f1")
	}
	f2 := func() {
		fmt.Println("f2")
	}
	f3 := func() {
		fmt.Println("f3")
	}
	l := list.New()
	// e4 := l.PushBack(4)
	// e1 := l.PushFront(1)
	// l.InsertBefore(3, e4)
	// l.InsertAfter(2, e1)
	l.PushBack(f1)
	l.PushBack(f2)
	l.PushBack(f3)
	for e := l.Front(); e != nil; e = e.Next() {
	}
}
