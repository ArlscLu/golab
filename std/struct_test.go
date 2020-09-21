package std

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

type Base1 struct {
	Name string
	Age  byte
	Sex  byte
}

type Base2 struct {
	OtherName string
}

type Student struct {
	Base1
	Base2
	p *Student
}

func (b *Base1) Change() {
	// tmp := []byte(b.Name)
	// for k := range tmp {
	// 	tmp[k] = 'x'
	// }
	// b.Name = string(tmp)
	// tmp := bytes.Buffer{}
	// tmp.WriteString("--tom")
	tmp := bytes.Buffer{}
	tmp.WriteString("hello")
	tmp.WriteString(", ")
	tmp.WriteString("world")
	// b.Name = tmp.String()
	b.Name = "jello"

}
func (b *Base1) Now() string {
	return b.Name
}

var s1 = &Student{
	Base1: Base1{
		Name: "Tom",
	},
	Base2: Base2{
		"Jackson",
	},
}

func TestA(t *testing.T) {
	a := "hello"
	b := []byte(a)
	b = append(b, 's')
	a = string(b)
	t.Logf("%s", a)
	log.Printf("%+v", s1)
	s1.Change()
	s1.Base1.Change()
	fmt.Println(s1.Now())
}
