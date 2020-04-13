package types

var Arr0 = [3]int{1, 2, 3}

var Slice0 = []int{4, 5, 7, 8, 9, 10}

var Slice1 = make([]int, 5, 10)

var Slice2 = []int{1, 2, 3}

type Employee struct {
	ID   int
	name string
}

var E1 Employee

type Complex struct {
	ID        int
	Source    []int
	IsMarried bool
	Remark    string
	Employee  Employee
	AgeExp    [5]float32
}
type Weekday int
