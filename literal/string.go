package literal

import "fmt"

func ShowString() {
	a := `hello golab  外国人`
	for k, v := range a {
		fmt.Print(k, v, " ")
		fmt.Printf("%U %x\n", v, v)
	}
}
