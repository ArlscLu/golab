package dp_test

import (
	"testing"

	"github.com/arlsclu7/golab/dp"
)

func TestSingle(t *testing.T) {
}

func TestMessagePool(t *testing.T) {
	msg0 := dp.Instance().GetMsg()
	if msg0.Count != 0 {
		t.Errorf("expect msg count %d, but actual %d.", 0, msg0.Count)
	}
	msg0.Count = 1
	dp.Instance().AddMsg(msg0)
	t.Errorf("%d", msg0.Count)
	msg1 := dp.Instance().GetMsg()
	if msg1.Count != 1 {
		t.Errorf("expect msg count %d, but actual %d.", 1, msg1.Count)
	}
}
