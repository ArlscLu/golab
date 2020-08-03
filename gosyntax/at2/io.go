package at2

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

var r io.Reader

type s1 []int

var somes = &ss{}

var gs = ss{}

func Showio() {

	fmt.Println(info(somes, 100))
	r = os.Stdin
	r = bufio.NewReader(r)
	r = new(bytes.Buffer)
	f, _ := os.Open("../at/rw.go")
	r = bufio.NewReader(f)

}

func Showinter() {
	var a interface{}
	a = s1{1, 2, 3}
	// if _, ok := a.(s1); ok {
	// 	fmt.Println("是s1类型")
	// }
	//特殊语法
	// switch a.(type) {
	switch a.(type) {
	case s1:
		fmt.Println("是s1类型")
	case int:
		fmt.Println("是int类型")
	}

}

type Simpler interface {
	Get() int
	Set(int)
	SimplerOther
}

type SimplerOther interface {
	Run(string) string
}

type ss struct {
	age int
}

func (s *ss) Get() int {
	return s.age
}

func (s *ss) Set(a int) {
	s.age = a
}

func (s *ss) Run(a string) string {
	return a
}

func info(sim Simpler, age int) int {
	sim.Set(age)
	return sim.Get()
}
