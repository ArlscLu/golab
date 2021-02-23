package df

import (
	"testing"
	"unicode"
)

func TestUnicode(t *testing.T) {
	r := 'a'
	t.Logf("%c", unicode.ToUpper(r))
	t.Log(unicode.ToUpper(r))
}
