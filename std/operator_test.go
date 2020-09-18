package std_test

import (
	"fmt"
	"testing"
)

var b1 = 0b_1001
var b2 = 0b_1111

func TestOperator(t *testing.T) {

	fmt.Printf("%b\n", b1|b2)
	fmt.Printf("%b\n", b1&b2)
	fmt.Printf("%b\n", b1^b2)
	fmt.Printf("%b\n", b1&^b2)
}
