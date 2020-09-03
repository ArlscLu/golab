package std

import (
	_ "encoding/json"
	"fmt"
	"time"
)

type student struct {
	Name  string
	Grade uint8
}

type classRoom []*student

func modify(c *classRoom) {

}

type School interface {
	Study()
	Play()
}

type Shangta struct {
	Address string
}

func (sc *Shangta) Study()  { fmt.Println("school is for study") }
func (sc *Shangta) Play()   { fmt.Println("school is for play") }
func (sc *Shangta) Travel() { fmt.Println("school is for travel") }

var _ School = (*Shangta)(nil)
var sc1 *Shangta = &Shangta{"shanghai shangta road 1032"}
var sc2 School

var Class1 = classRoom{
	&student{"tom", 3},
	&student{"jack", 4},
	&student{"rose", 5},
	&student{"andy", 6},
}

type route chan int

func ShowSlice(c classRoom) {
	r1 := make(route)
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Second * 1)
			r1 <- i % 2
		}
	}()
	for {
		if tmp := <-r1; tmp == 0 {
			fmt.Println(`过了1秒`)
		}
	}

	// 	sc1.Play()
	// 	sc1.Study()
	// 	sc1.Travel()
	// 	sc2 = &Shangta{"shanghai shangta road 1032"}
	// 	sc2.Play()
	// 	fmt.Printf("外层class的值  %p\n", Class1)
	// 	fmt.Printf("形参的值%p\n", c)
	// 	fmt.Printf("外层class的保存地址  %p\n", &Class1)
	// 	fmt.Printf("class的保存地址%p\n", &c)
	// }
}
