package df

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

func TestSnippet(t *testing.T) {
	testCases := []struct {
		desc    string
		example []string
	}{
		{
			desc:    "测试基本",
			example: []string{"h", "e", "l", "l", "o"},
		},
		{
			desc:    "下一个",
			example: []string{"get", "away"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			for i, v := range tC.example {
				t.Logf("第%d %s\n", i, v)
			}
		})
	}
}

type (
	A struct{}
	B interface{}
	C struct {
		name string
		age  int
	}
)

var a = C{}
var bb B

//TestBaseSlice slice的存储
func TestBaseSlice(t *testing.T) {
	testCases := []struct {
		desc string
		hold [11]int
	}{
		{
			desc: "slice与array的比较",
			hold: [...]int{7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Logf("原来 len %d,caps%d", len(tC.hold), cap(tC.hold))
			part := tC.hold[0:3]
			t.Logf("part len %d,caps%d", len(part), cap(part))
			//如果能使用底层数组则使用,不够时就新建底层数据
			extend := tC.hold[4:7]
			extend = append(extend, []int{100, 101, 102}...)
			t.Log(extend)
		})
	}
}

// The underlying types of the following ones are both int.
type (
	MyInt int
	Age   MyInt
)

func Test(t *testing.T) {
	base := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := base[3:5:7]
	s2 := base[:]
	s1 = append(s1, 10)
	s2 = append(s2, 20)
	t.Log(s1, s2)
}

func TestMap(t *testing.T) {
	var mContainer = make(map[string]int, 2)
	mContainer["Tom"] = 10
	mContainer["Jack"] = 12
	mContainer["Rose"] = 22
	t.Log(len(mContainer))
}
