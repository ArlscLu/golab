/**
字面量
*/
package literal

import "fmt"

var A uint8 = 0b_101
var B uint8 = 0o_17
var C uint8 = 15
var D uint8 = 0x_1A

func Show() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
}
