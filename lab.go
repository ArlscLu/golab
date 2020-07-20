package main

import (
	"fmt"
	"os"

	"github.com/arlsclu7/golab/gosyntax"
)

func main() {
	fmt.Println(gosyntax.Ppmethod())
	os.Exit(11)
	gosyntax.Pmethod()
	os.Exit(22)
	st := gosyntax.SomeTest{
		Name: "Tom",
		Age:  11,
	}
	gosyntax.NeedTester(st)
	os.Exit(12)

	// gosyntax.Pstruct()
	// fmt.Println(gschild.Ta)
	// one := make(chan int)
	// two := make(chan int)

	// gosyntax.Ppanic()
	// one := make(chan int)
	// two := make(chan int)

	// go func() {
	// 	one <- 100
	// }()

	// go func() {
	// 	v := <-one
	// 	two <- v
	// }()

	// fmt.Println(<-two)
	// os.Exit(1)

	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"Blog":   "www.flysnow.org",
	// 		"wechat": "flysnow_org",
	// 	})
	// })
	// r.Run(":8080")
	// fmt.Println(3 << 5)
	// gosyntax.Pstruct()
	// os.Exit(6)
	// reader := gosyntax.NewAlphaReader(strings.NewReader("Hello! It's 9am, where is the sun?"))
	// p := make([]byte, 4)
	// for {
	// 	n, err := reader.Read(p)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Print(string(p[:n]))
	// }
	// os.Exit(5)
	// fmt.Println()
	// gosyntax.Pio()
	// os.Exit(4)
	// gosyntax.HtmlParse()
	// os.Exit(3)
	// gosyntax.Pliteral()
	// os.Exit(2)
	// gosyntax.Pslide()
	// os.Exit(2)
	// gosyntax.R()
	// rst := gosyntax.Callback(5, func(a int, b int) int {
	// 	return a * b
	// })
	// fmt.Println(rst)
	// f()
	// cfibonacci(10)
	// os.Exit(1)
	// gosyntax.Pliteral()
	// gosyntax.Prin()
	// gosyntax.Si()
	// fmt.Println(show())
	// os.Exit(1)
}

func show() uint8 {
	defer func() {
		fmt.Println(111)
	}()
	fmt.Println(122222)
	return 10
}

/**
defer fmt.Printf("%d ", 0)
defer fmt.Printf("%d ", 1)
defer fmt.Printf("%d ", 2)
defer fmt.Printf("%d ", 3)
defer fmt.Printf("%d ", 4)
key  type *address
*/
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func fibonacci(n int) (rst int) {
	if n <= 1 {
		rst = 1
	} else {
		rst = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func cfibonacci(n int) int {
	rst := 0
	for i := 0; i < n; i++ {
		rst += fibonacci(i)
		fmt.Println(i, n, rst)
	}
	return rst
}
