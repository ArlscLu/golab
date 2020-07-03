package gosyntax

import "fmt"

func Si() {
	const (
		a1 uint8 = 10
	)
	type stype struct {
		name  []byte
		sex   byte
		child *stype
	}
	s2 := stype{
		[]byte{1, 2, 3},
		1,
		nil,
	}
	s1 := stype{
		[]byte{1, 2, 3},
		1,
		&s2,
	}
	fmt.Println(s1)

}
