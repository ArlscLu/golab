package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/arlsclu7/golab/gosyntax"
	"github.com/arlsclu7/golab/gosyntax/at"
	"github.com/arlsclu7/golab/gosyntax/at2"
	"github.com/arlsclu7/golab/gosyntax/memory"
)

func main() {
	// gosyntax.ShowList()
	// os.Exit(2)
	// at.ShowRw()
	// at.ShowRw123()
	// at.Showm2()
	// os.Exit(6)

	at.NotChangeV(at.OriginSlice)
	at.ChangeV(&at.OriginSlice)
	os.Exit(9)
	at.ShowTypes()
	os.Exit(8)
	at2.Showio()
	at2.Showinter()
	os.Exit(7)
	for i, v := range memory.Debug() {
		if v == nil {
			continue
		}
		fmt.Println(i, v)
		// if i%2 != 0 {
		// 	fmt.Println("       ")
		// }
	}
	os.Exit(6)
	at.ShowMemory()
	os.Exit(5)
	at.ShowMtest2()
	os.Exit(4)
	at.Std()
	os.Exit(3)
	str := `
	begin
	*************
				**********

	**********

	*********************  **********
				*******
	
	`
	io.WriteString(os.Stdout, str)
	at.ShowRw321()
	os.Exit(2)
	at.Show()
	os.Exit(1)
	r := strings.NewReader("hello world大中国")
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
	fmt.Printf("%s", b)
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
