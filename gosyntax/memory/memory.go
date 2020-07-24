package memory

type bowl []interface{}
type person struct {
	name string
	age  uint8
}

type reference struct {
	label string `标签      a`
	rtype string `类型 int,[3]int,[]int...`
	addr  int32  `16进制表示的地址  0xc0000c0040`
}

type point struct {
	label string `标签      a`
	rtype string `类型 int,[3]int,[]int...`
	addr  int32  `16进制表示的地址  0xc0000c0040`
}

/*reference 是指向固定的 pointer
//变 ｜ 不变
//上层的rtype决定下层取值的方式
	a [3]int  0x1
	b []int len3 cap3   0x2
	c person{name,age}    0x3
	d *int   0x4
	if  type==[3]int  {
		value = 0x1 + [3]int
	}
	if type==[]int len3 cap3 {
		value = ox2 + []int len3 cap3
	}
	if type==*int {
		value = 0x3 + *int
	}
	根据父type 来决定是不是解引用  dereference
	初始    type+addr   =>*addr
	指针变量  *int  +addr   值为 *addr+int        =>*addr
	引用变量   []int len3 cap3 +addr   值为  *addr int len3 cap3    => (*addr int len3 cap3)()


	key  |  type  |  ptype  |addr
			*int       /      0x2    yes
			int        *int     0x3    no
			[]intl3c3   /       0x4    yes
			intl3c3     []intl3c3   0x5   yes
   type + addr
   普通类型   int    0x1
   指针类型   *int   0x2
   引用类型     []int  0x3

	存储：    addr+type
	指针 包含   addr+type
	三种形式：
	value                        []
	addr+type    				&[]
	addr+type()					*&[]

	根据type决定表达式返回什么

	anything  addr+type
	switch($type) {
		case *int
		then addr    //pointer类型
		case [3]int,int
		then *addr(value)
		case []int
		then *addr(reference)
	}

*/

var eparent = 8

func Debug() []interface{} {
	bowls := make(bowl, 15)
	var a [3]int = [3]int{1, 2, 3}  //[3]int
	var b []int = []int{10, 10, 10} //[]int
	var c = person{                 //struct
		name: "Tom",
		age:  18,
	}
	var d string = "aerewertt" //string
	var e *int = &eparent
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
	bowls[7] = &d //内存地址
	bowls[8] = f  //内存地址
	bowls[9] = f()
	bowls[10] = e //内存地址
	bowls[11] = eparent
	return bowls
	//引用类型 in   （slice,map ,channel)  强转
	//struct  隐式转换
}
