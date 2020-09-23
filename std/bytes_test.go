package std

import (
	"bytes"
	"testing"
)

var a1 = []byte("abe")
var b1 = []byte("abc")
var c1 = []byte{'\377'}

func TestBytes(t *testing.T) {
	t.Log(bytes.Compare(a1, b1))
	if bytes.Compare(a1, b1) == 1 {
		t.Log("a is bigger or  b")
	} else {

	}
	t.Log(bytes.Compare([]byte{'\n'}, []byte{9, 1}))
	t.Log(bytes.Contains([]byte("你好"), []byte("好")))
	t.Log(bytes.ContainsAny([]byte("你好"), "好"))
	t.Log(bytes.ContainsRune([]byte("你好"), '你'))
	t.Log(bytes.Index([]byte("chicken"), []byte("ken")))
}
