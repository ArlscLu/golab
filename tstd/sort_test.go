package tstd

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"testing"
)

var d = []int{1, 23, -1, 32, -9, -100, 0, 3, 8, 15}

func TestSlice(t *testing.T) {
	var d = []int{1, 23, -1, 32, -9, -100, 0, 3, 8, 15}
	var b int = 10
	var c *int = &b
	_ = c
	sort.Slice(d, func(i, j int) bool {
		return d[i] < d[j]
	})
	Show()
}

func Show() {
	fmt.Print(d)
}

func TestMain(m *testing.M) {
	fmt.Print("XXXX\n", os.Args)
	flag.Parse()
	os.Exit(m.Run())
}

func TestF1(t *testing.T) {
	fmt.Print("F1\n")
}
func TestF2(t *testing.T) {
	fmt.Print("F2\n")
}
