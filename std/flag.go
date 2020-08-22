package std

import (
	"flag"
	"fmt"
	"log"
)

var (
	averageAge int = 18
)
var age int

type birthyear struct {
}

func (b birthyear) String() string {
	return "hello"
}

func (b birthyear) Set(s string) error {
	return nil
}
func Runflag() {
	var ip *int
	ip = flag.Int("ip", 127, "address of ip")

	flag.IntVar(&age, "age", 11, "help message for name")

	fmt.Println(*ip, age)
	var b birthyear
	flag.Var(&b, "birthyear", "help msg for birthyear")

	flag.Parse()
	if age > averageAge {
		log.Println("you are older than average")
	} else {
		log.Println("you are older than younger")
	}
}
