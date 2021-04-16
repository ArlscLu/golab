package df

import (
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestRune(t *testing.T) {
	rs := make([]rune, 10, 20)
	rs[0] = 'a' //normal
	rs[1] = '中' //中文字符
	rs[2] = '\377'
	rs[3] = '\x10'
	rs[4] = '\u9646'
	rs[5] = '\u4f1f'
	rs[6] = '\u6770'
	name := [3]rune{
		'\u9646', '\u4f1f', '\u6770',
	}
	t.Log(rs)
	for _, r := range rs {
		t.Logf("%d\n", utf8.RuneLen(r))
		if unicode.IsLetter(r) {
			t.Logf("%c是字母\n", r)
		}
	}
	t.Logf("%c", name)

}

func TestStringLiteral(t *testing.T) {
	var ss = []string{
		"a",
		"\u9646\u4f1f\u6770",
	}
	t.Log(ss)
	t.Logf("len: %d  caps: %d", len(ss), cap(ss))
}

const count = 10
