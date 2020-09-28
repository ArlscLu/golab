package dp_test

import (
	"testing"

	"github.com/arlsclu7/golab/dp"
)

func TestSingle(t *testing.T) {
	t.Logf("%v", dp.Single)
}

func TestMessagePool(t *testing.T) {
	msg0 := msgpool.Instance().GetMsg()
	if msg0.Count != 0 {
		t.Errorf("expect msg count %d, but actual %d.", 0, msg0.Count)
	}
	msg0.Count = 1
	msgpool.Instance().AddMsg(msg0)
	msg1 := msgpool.Instance().GetMsg()
	if msg1.Count != 1 {
		t.Errorf("expect msg count %d, but actual %d.", 1, msg1.Count)
	}
}
