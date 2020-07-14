package gosyntax

import "fmt"

var a [6]byte = [6]byte{'a', 'b', 'c', 'd', 'e', 'f'}

func Pslide() {
	chi := "中国china"
	fmt.Println("range处理string")
	fmt.Printf("数据类型%T\n", chi)
	for k, v := range chi {
		fmt.Println(k, v)
	}

	chislide := []byte(chi) //转换为[]byte
	fmt.Println("range处理[]byte")
	fmt.Printf("数据类型%T\n", chislide)
	for k, v := range chislide {
		fmt.Println(k, v)
	}

	chisliderune := []rune(chi) //转换为[]rune
	fmt.Println("range处理[]rune")
	fmt.Printf("数据类型%T\n", chisliderune)
	for k, v := range chisliderune {
		fmt.Println(k, v)
	}
	fmt.Println(chi, chislide, chisliderune)
	pkg := []byte{'h', 'a'}
	fmt.Println(pkg)
	a := []string{
		"日本語",                                  // UTF-8 input text
		`日本語`,                                  // UTF-8 input text as a raw literal
		"\u65e5\u672c\u8a9e",                   // the explicit Unicode code points
		"\U000065e5\U0000672c\U00008a9e",       // the explicit Unicode code points
		"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e", // the explicit UTF-8 bytes
	}

	/**
	 ""解析  Interpreted




	**/

	fmt.Println(a)
	fmt.Printf("%T", a[2])
	for k, v := range a[2] {
		fmt.Println(k, v)
	}
	fmt.Println("\nddd")
}
