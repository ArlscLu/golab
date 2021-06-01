package tinterface

import (
	"fmt"
	"sort"
	"testing"
)

//writing generic algorithm （泛型编程）

type person struct {
	Age int
}
type byAge []person

func (ba byAge) Len() int           { return len(ba) }
func (ba byAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }
func (ba byAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age }

func TestGeneric(t *testing.T) {
	persons := []person{
		{18}, {19}, {-1}, {3},
	}
	fmt.Printf("原始信息%v\n", persons)
	sort.Sort(byAge(persons))
	fmt.Printf("处理后信息%v\n", persons)
}
