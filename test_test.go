package main

import (
	"fmt"
	"testing"
)

var expect = []struct {
	in  int
	out bool
}{
	{1, false},
	{2, true},
	{3, false},
	{4, false},
	{5, true},
}

func IsEven(i int) (isEven bool) {
	isEven = false
	if i%2 == 0 {
		isEven = true
	}
	return
}
func TestCompare(t *testing.T) {
	for _, v := range expect {
		real := IsEven(v.in)
		if real != v.out {
			fmt.Println("结果不符合预期", v)
		}
	}

}

//example形的testoutput,unordered
func ExampleIsEven() {
	// fmt.Println(IsEven(10))
	//Output:true
	fmt.Print(IsEven(10), "\n", IsEven(11))
	//Unordered output:true
	//false

}

func TestFail(t *testing.T) {
	t.Log("xxx no error")
	// t.Error("xxx error happens")
}

//判断是否为回文，正读与反读一致
func IsPalindrome(s string) bool {
	//未考虑unicode，标点符号
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
	//完善的
	// letters := []rune{}
	// for _, r := range s {
	// 	if unicode.IsLetter(r) {
	// 		letters = append(letters, unicode.ToLower(r))
	// 	}
	// }
	// for i := range letters {
	// 	if letters[i] != letters[len(letters)-1-i] {
	// 		return false
	// 	}
	// }
	// return true
}

func TestSimplePalindrome(t *testing.T) {
	if !IsPalindrome("ABCBA") {
		t.Error(`TestSimplePalindrome 暴露了bug`)
	}
}

//需要考虑unicode字符
func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`TestFrenchPalindrome 暴露了bug`)
	}
}

//需要考虑标点符号忽略
func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Error(`TestCanalPalindrome 暴露了bug`)
	}
}

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayaK", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},
		// semi-palindrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v want %v", test.input, got, test.want)
		}
	}
}
