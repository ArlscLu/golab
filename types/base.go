package types

//B bbb
var B bool = true

//Num num
var Num int = 13

//X string
var X string = "U+000A"

//Float0 test
var Float0 float32 = 12.23222

const (
	C0 = 1<<iota - 1
	C1 = 1 << iota
	C2 = 1 << iota
	C4 = '中'
)
