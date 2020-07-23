package at

import (
	"bufio"
	"fmt"
	"os"
)

/*
	读写相关
*/
var (
	firstName, lastName string
	i                   int
	f                   float32
	input               = "56.12/5212/Go"
	format              = "%f/%d/%s"
)

func ShowRw() {
	fmt.Println("Please enter your full name: ")

	fmt.Scanln(&firstName, &lastName)
	// fmt.Println(firstName)
	// fmt.Println(lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)

	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels

	fmt.Sscanf(input, format, &f, &i, &s)

	fmt.Println("From the string we read: ", f, i, s)

}

func ShowRw123() {
	var (
		a float32
		b int
		d string
	)
	// a, d := "", 0
	// n, err := fmt.Sscanf("tomred/1/123", "%s/%b/%d", &a, &b, &d)
ex:
	n, err := fmt.Sscanf("56.12/5212/Go", "%f/%d/%s", &a, &b, &d)
	if err != nil {
		goto ex
		// panic("不合法的输入")
	}
	fmt.Println(n, a, b, d)
}

func ShowRw321() {
	var (
		a int
		b int
		c int
	)
	t_sum := 10
	fmt.Println("输入3个数 结果为", t_sum)
next:
	fmt.Scanln(&a, &b, &c)
	sum := a + b + c
	if sum != t_sum {
		goto next
	}
	fmt.Println("you success")
}

func Std() {
n:
	fmt.Printf("input your name \n")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	if input == "s\n" {
		fmt.Printf("hello %s", input)
		return
	}
	goto n

}
