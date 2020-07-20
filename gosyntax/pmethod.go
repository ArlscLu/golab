package gosyntax

/*
	变量提升相关
	S包含T时,      s,*s 包含t    *s包含*t
	S包含*T时，    s,*s 包含天，*t的所有方法

	语法糖的存在   t.method  转换为&t.method()
*/

type base struct {
	pagesize uint8
}

func (b base) getValue() uint8 {
	return b.pagesize
}

func (b *base) changeValue() {
	b.pagesize = b.pagesize + 10
}

var bb1 = base{
	15,
}

type extend struct {
	*base
}

func Ppmethod() []interface{} {
	// var r []interface{}
	r := []interface{}{
		1,
		2,
		"asd",
		[3]int{1, 2, 3},
	}
	r = append(r, 10)
	r = append(r, map[string]int{
		"title": 19,
		"name":  11,
	})
	return r
	// fmt.Println(extend.getValue(extend{}))
	// ee1 := extend{}
	// extend.changeValue(ee1)
	// fmt.Println(ee1)
	// fmt.Println(ee1.getValue())
	// extend.getValue(extend{})
	// f := extend.getValue
	// ff := ee1.getValue
	// fmt.Println("来自提取的method", f(extend{}))
	// fmt.Println("来自instance提mm取的method", ff())
}
