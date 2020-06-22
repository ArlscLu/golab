package block

import "fmt"

var a int = 10 //file block a

func MyBlock() {
	x := 1000
	x = x << 2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", 1000)
	type t struct {
		name        string
		age         uint8
		chineseName string
		englishName []byte
	}
	t1 := t{
		name:        "tom",
		age:         18,
		chineseName: "中山新",
		englishName: []byte{14, 20},
	}
	us := []rune{'\u4e16', '\u0005'}
	ch := "hello中山新"
	for i, v := range ch {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", t1.chineseName)
	fa := func() int {
		return 10 * 100
	}
	fmt.Println(fa())
	/**
	[104][101][108][108][111]
	5[20013][][]    8[23665][][]    11[26032][][]

	**/
	fmt.Println(`new line`)
	fmt.Printf("%d", '\u4e16')
	fmt.Println(us)
	fmt.Printf("%c", us[1])
	fmt.Printf("%#v, %[1]T,%c", t1, t1.chineseName[0])
	fmt.Printf("%c", t1.chineseName[0])
	fmt.Println()
	return
	// if a := true; a {
	// 	fmt.Println(a)
	// }
	// for a := 1; a < 10; a++ {
	// 	fmt.Println(a)
	// }
	// if a == 10 {
	// 	fmt.Println("a from func block")
	// }
}
