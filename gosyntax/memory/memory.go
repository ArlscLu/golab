package memory

type bowl []interface{}
type person struct {
	name string
	age  uint8
}

func Debug() []interface{} {
	bowls := make(bowl, 10)
	var a [3]int = [3]int{1, 2, 3}  //[3]int
	var b []int = []int{10, 10, 10} //[]int
	var c = person{                 //struct
		name: "Tom",
		age:  18,
	}
	var d string = "aerewertt" //string
	var f func() int = func() int {
		return 10
	}
	bowls[0] = a
	bowls[1] = &a
	bowls[2] = b
	bowls[3] = &b
	bowls[4] = c
	bowls[5] = &c
	bowls[6] = d
	bowls[7] = &d
	bowls[8] = f
	bowls[9] = f()
	return bowls
	//引用类型 in   （slice,map ,channel)  强转
	//struct  隐式转换
}
