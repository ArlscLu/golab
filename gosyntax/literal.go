package gosyntax

/**
 literal：
 integer, floating-point, string,rune,imaginary,
 function ，composite
**/
import "fmt"

var ints = []int{
	0b11111111, //二进制255
	0o2_7,      //8进制23
	0x111,      // 16进制 1*16^2+1*16+1*1 = 273
	5,
	0xFACE,
}

var floats = []float32{
	0.5,
	1.111,
	1E10,
	1E-2,
	1.5E4,
}
var runes = []int{
	'\a',
	'a',
	'中',
	'\u0012',
	'\xff',
	'\377',
	'\a',
}

func Pliteral() {
	fmt.Println(ints)
	fmt.Println(floats)
	fmt.Println(runes)
	fmt.Printf("%T", '\a')

	fmt.Printf("%T\n", add)

	// "func(int, int) int"

	fmt.Printf("%T\n", sub)

	// "func(int, int) int"

	fmt.Printf("%T\n", first) // "func(int, int) int"

	fmt.Printf("%T\n", zero)
	if zero(1, 2) == zero1(3, 4) {
		fmt.Println("zero and zero1 are the same func")
	}
	// "func(int, int) int"
}
func add(x int, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	z = x - y
	return
}

func first(x int, _ int) int { return x }

func zero(int, int) uint8 {
	return 0
}
func zero1(int, int) uint8 {
	return 10
}
