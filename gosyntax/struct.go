package gosyntax

import "fmt"

type emptyStruct struct {
}

type Nspace struct {
	MiddleName [3]byte
}

type Xspace struct {
	Nspace
	Sex  byte
	Name [2]byte
}

var xs = Xspace{}

func (n *Nspace) Xprint(a []rune) {
	fmt.Println(n.MiddleName, a)
}

func Pstruct() {
	// r := "我是风温柔委婉"
	xs.Xprint([]rune{'我', '是'})
}
