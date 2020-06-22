package block

import "fmt"

var A []int = []int{1, 2, 3, 4}

func Change() {
	A := []string{"a", "b", "c"}
	A = append(A, "z")
	fmt.Println(A)
}
