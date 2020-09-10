package std

import "testing"

type Yer interface {
	Mm()
}
type yy struct {
	Name string
	Age  int
	yn   *yy
}
type someint int

func (y yy) Mm() {}
func TestAssert(t *testing.T) {
	//x.(T)
	//T is not interface

	//T is interface
	const X = iota << 2
	var a interface{} = 7
	if v, ok := a.(someint); ok {
		t.Log(v)
	}
	y1 := *new(yy)
	testAss1(y1, t)
}

func testAss1(y1 Yer, t *testing.T) {
	if v, ok := y1.(yy); ok {
		t.Log(v)
	}
}
