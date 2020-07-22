package at

import "fmt"

var qat at
var qat_ *at

func ShowMemory() {
	qat2_ := &at{1, 2, 3}
	sl0 := 8
	sl1 := []int{4, 5, 6}
	sl2 := &[]int{7, 8, 9}
	sl3 := [3]int{10, 11, 12}
	sl4 := &[3]int{13, 14, 15}

	/*
		qat2_    *at    &[1 2 3]
	*/
	fmt.Println(qat, &qat) //空struct
	newline()
	fmt.Println(qat_, &qat_) //空struct 指针
	newline()
	fmt.Println(qat2_, &qat2_, *qat2_) //struct 指针
	newline()
	fmt.Println(`sl0@int       `, sl0, &sl0)
	newline()
	fmt.Println(`sl1@slice的情况     `, sl1, &sl1)
	newline()
	fmt.Println(`sl2@slice的情况#指针#    `, sl2, &sl2, *sl2)
	newline()
	fmt.Println(`sl3@array的情况    `, sl3, &sl3)
	newline()
	fmt.Println(`sl4@array的情况#指针#     `, sl4, &sl4, *sl4)
}
func newline() {
	fmt.Println()
}
