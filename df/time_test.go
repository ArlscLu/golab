package df

import (
	"testing"
	"time"
)

func TestTimeBase(t *testing.T) {
	var l *time.Location
	l = time.Local
	l = time.UTC
	_ = l
	n := time.Now()
	t.Logf("%d", n.Month())
	t.Log(n.Unix())
}
