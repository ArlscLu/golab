package at

import "fmt"

var qat at
var qat_ *at
var sl7 *[3]int

type S struct {
	a int
	b float32
}
type tslice []int

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // has type Vertex
	q = &Vertex{1, 2} // has type *Vertex
	r = Vertex{X: 1}  // Y:0 is implicit
	s = Vertex{}      // X:0 and Y:0
)

func Showm2() {
	/*

		t := *q makes a copy of the struct pointed to by q

	*/
	// t := *q
	// q.X = 4
	// u := *q
	// q.X = 100
	// uu := *q
	// fmt.Println(p, q, r, s, t, u, t == u, uu)
	t := q
	q.X = 4
	u := *q
	fmt.Println(p, q, r, s, t, u, *t == u)
}

func ShowMemory() {
	//done
	ss := new(S)

	gs := new(tslice)

	qat2_ := &at{1, 2, 3}
	sl0 := 8
	sl1 := []int{4, 5, 6}
	sl2 := &[]int{7, 8, 9}
	sl3 := [3]int{10, 11, 12}
	sl4 := &[3]int{13, 14, 15}
	sl5 := st{"tom"}
	var sl6 *[3]int
	// var slx *int
	// *slx = 10
	// *sl6 = [3]int{7, 9, 11}
	sl7 = &[3]int{1, 2, 3}
	// sl6 := &[3]int{4, 5, 6}

	/*
		qat2_    *at    &[1 2 3]
	*/
	newline()
	fmt.Println(ss.a) //struct解引用了  ss.a=>(*ss).a 非强制性 原来的还在，只在需要的时候转换
	fmt.Println((*ss).a)
	newline()
	fmt.Println(gs)
	fmt.Println(*gs) //会输出[], 实际为nil
	fmt.Println(*gs == nil)
	newline()
	fmt.Println(qat, &qat) //空array
	newline()
	fmt.Println(qat_, &qat_) //空array 指针    强制获得地址后直接取对应值（附注：addr+remark 这样copy的也算地址）
	newline()
	fmt.Println(qat2_, &qat2_, *qat2_) //array 指针   存的只有addr  和上面做比较
	newline()
	fmt.Println(`sl0@int       `, sl0, &sl0)
	newline()
	fmt.Println(`sl1@slice的情况     `, sl1, &sl1) //强制获得地址后直接取对应值
	newline()
	fmt.Println(`sl2@slice的情况#指针#    `, sl2, &sl2, *sl2)
	newline()
	fmt.Println(`sl3@array的情况    `, sl3, &sl3, *&sl3)
	newline()
	fmt.Println(`sl4@array的情况#指针#     `, sl4, &sl4, *sl4)
	newline()
	fmt.Println(`sl5@struct的情况     `, sl5, &sl5, *&sl5)
	newline()
	fmt.Println(`sl6@array的情况#指针#     `, sl6, &sl6, *&sl6) // 强制获得地址后直接取对应值
	newline()
	fmt.Println(`sl7@*int的情况#指针#    `, sl7, &sl7, *sl7, *&sl7) //struct解引用了
}
func newline() {
	fmt.Println()
}

/*
	指针的时候  存的地址为 &[1,2,3]
	申明的也会分配一个内存地址，但其中的值为nil
	当为引用类型如(slide,map,interface 时) 存储的是地址

	[key|type|addr]
	if(type in (pointers, functions, interfaces, slices, channels, and maps)) {
		addr = 0x0;
		value = nil
	} elseif(type = point) {
		addr != 0x0
		value = &[1,2,3]
	} else {
		addr != 0x0
		value = {tom}
	}


	[key | type  | addr]

	a   *point 0x0
	b   []int  0x0
	c  int 	   0x1
	d   *int   0x2
*/
