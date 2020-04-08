package helper

import (
	"fmt"
)

//NewLine 输出空行便于使用
func NewLine() {
	fmt.Println("\n\n******end*****")
}

//StartLine 输出开始行
func StartLine() {
	fmt.Println("*****start******")
}

//DivideLine 分割行
func DivideLine() {
	fmt.Println("********divide**********")
}
