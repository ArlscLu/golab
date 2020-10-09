package std

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type typeme struct {
	id int
}

func (t typeme) String() string {
	return `this is string `
}
func (t typeme) Set(string) error {
	return nil
}

func main() {
	//create default flagset
	var age2 string
	var s string
	var me1 typeme
	age := flag.Int(`age`, 17, `print out your age`)
	flag.StringVar(&age2, `age2`, `shiwu`, `print out your age`)
	flag.StringVar(&s, `name`, `tom`, `print out your name`)
	flag.Var(&me1, `isfemale`, `check is female`)

	ismale := flag.Bool(`ismale`, false, `check is male`)
	girlFriend := flag.String(`gf`, `rose`, `print your gf`)
	flag.Parse()
	if *age > 19 {
		fmt.Println(`you are older than 19`)
	}
	fmt.Println(age2)
	fmt.Println(me1)
	if *ismale {
		if strings.IndexRune(*girlFriend, 'n') == 6 {
			os.Stdout.WriteString(*girlFriend + "\n")
		}
		os.Stdout.WriteString(s)
	} else {
		fmt.Println(`not is male`)
	}
}
