package std_test

import (
	"flag"
	"fmt"
	"testing"
)

func TestFlag(t *testing.T) {
	age := flag.Int(`age`, 17, `print out your age`)
	flag.Parse()
	fmt.Print(*age)
	// fmt.Print(flag.CommandLine.Arg(0))
	// fmt.Print(flag.Arg(0))
}
