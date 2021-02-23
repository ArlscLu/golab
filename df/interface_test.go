package df

import (
	"testing"
)

type Summary interface {
	showFrom() string
}

type Toy struct {
}

func (t Toy) showFrom() string {
	return "china"
}

//TestImplInterface 测试是否实现接口
func TestImplInterface(t *testing.T) {
	t.Log(1)
	var _ Summary = Toy{}
	var _ Summary = (*Toy)(nil)
}
