package tstd

import (
	"fmt"
	"sort"
	"testing"
)

var d = []int{1, 23, -1, 32, -9, -100, 0, 3, 8, 15}

func TestSlice(t *testing.T) {
	sort.Slice(d, func(i, j int) bool {
		return d[i] < d[j]
	})
	Show()
}

func Show() {
	fmt.Print(d)
}
