package std

import (
	"testing"
	"unicode"
)

var upkgAssert = []struct {
	input  rune
	expect bool
}{
	{1, true},
	{97, true},
	{20, true},
	{50, true},
	{1, true},
}

func TestUnicodePkg(t *testing.T) {
	for _, v := range upkgAssert {
		rst := isTrueOrNot(v.input)
		if rst != v.expect {
			t.Errorf("**error**isTrueOrNot(%d)= %t,want %t \n", v.input, rst, v.expect)
		} else {
			t.Logf("**success**isTrueOrNot(%d)= %t,want %t \n", v.input, rst, v.expect)
		}
	}
	for _, v := range upkgAssert {
		rst := unicode.IsDigit(v.input)
		if rst != v.expect {
			t.Errorf("**error**unicode.IsDigit(%d)= %t,want %t \n", v.input, rst, v.expect)
			continue
		}
		t.Logf("**success**unicode.IsDigit(%d)= %t,want %t \n", v.input, rst, v.expect)
	}

	for _, v := range upkgAssert {
		rst := unicode.IsNumber(v.input)
		if rst != v.expect {
			t.Errorf("**error**unicode.IsNumber(%d)= %t,want %t \n", v.input, rst, v.expect)
			continue
		}
		t.Logf("**success**unicode.IsNumber(%d)= %t,want %t \n", v.input, rst, v.expect)
	}
}

func isTrueOrNot(r rune) bool {
	if unicode.IsControl(r) {
		return true
	}
	return false
}
