package df

import "testing"

//testrune 测试rune字面量 的使用
func TestRuneLiteral(t *testing.T) {
	//rune == int32
	list := table{
		"\\n": []rune{
			'\n',
			'\u000A',
		},
	}
	t.Log(list)
}

type table map[string][]rune
