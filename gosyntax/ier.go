package gosyntax

import "fmt"

func NeedTester(t Tester) {
	fmt.Println("mircle happens")
	fmt.Println(t.outName())
	fmt.Println(t.outAge())
}

type Tester interface {
	outName() string
	outAge() byte
	Gogo()
}

type SomeTest struct {
	Name string
	Age  byte
}

func (st SomeTest) outName() string {
	return st.Name
}
func (st SomeTest) outAge() byte {
	return st.Age
}
func (SomeTest) Gogo() {
}
