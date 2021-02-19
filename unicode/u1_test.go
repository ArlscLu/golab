package unicode_test

import (
	"fmt"
	"testing"
)

var testdata = []struct {
	input rune
	mark  string
}{
	{input: 1},
	{input: '\377', mark: "八进制"},
}

func TestUnicode(t *testing.T) {
	t.Cleanup(func() {
		t.Error("make failure ,all cleanup")
	})
	for k, v := range testdata {
		t.Log(k, v)
	}
}
func TestAssert(t *testing.T) {
	var x interface{}
	x = 10
	if xint, ok := x.(int); ok {
		t.Logf("%dxxxxxx", xint)
	} else if ybool, ok := x.(bool); ok {
		t.Logf("%T", ybool)
	}

}

func TestAssert2(t *testing.T) {
	var x interface{}
	x = new(implString)
	switch next := x.(type) {
	case string:
		t.Log("存的是string")
	case fmt.Stringer:
		t.Log("实现了fmt.Stringer接口")
		t.Logf("\n%s\n", next.String())
	default:
		t.Log("empty")
	}
}

type implString int

func (i *implString) String() string {
	return "helloworld"
}
