package labdb

import (
	"container/list"
	"testing"
)

func TestListBase(t *testing.T) {
	l := list.New()
	_ = l.PushBack(10)
	_ = l.PushBack(10)
	_ = l.PushBack(10)
	t.Log(l.Len())
	t.Logf("%#v,%+v,%v", l.Front(), l.Back(), l.Front())
}
