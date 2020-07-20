package gosyntax

import "fmt"

type Point struct {
	x, y int
}

type Path []Point

func (p Point) Distance() int {
	return p.x - p.y
}

func (pa Path) Distance() []int {
	s := []int{}
	for i := range pa {
		s = append(s, pa[i].Distance())
	}
	return s
}

var pa1 = Path{
	{10, 1},
	{9, 2},
	{8, 3},
	{7, 4},
	{6, 5},
	{5, 4},
}

type Book struct {
	title  []rune
	price  float32
	author string
	bdis   Discount
}
type Discount = float32

func (b *Book) ForRead() {
	fmt.Println(b.title)
}
func (b Book) discount(dis Discount) float32 {
	return b.price * dis
}
func (b *Book) CheapPrice() float32 {
	return b.price - b.discount(0.3)
}
func (b *Book) UpdatePrice() {
	if b.bdis == 0.1 {
		b.price -= b.discount(0.1)
	}
	b.price = b.CheapPrice()
}

func (b *Book) detail() ([]rune, float32, string) {
	return b.title, b.price, b.author
}

var b1 = Book{
	title:  []rune{'麦', '赛', 'Q', 'a'},
	price:  75.8,
	author: "mm.风味.yy",
	bdis:   0.1,
}

func Pmethod() {
	fmt.Println(pa1.Distance())
	fmt.Println(b1.CheapPrice())
	b1.UpdatePrice()
	fmt.Println(b1.detail())
	fmt.Println(b1)
	fmt.Println(b1.discount(0.11))
}
