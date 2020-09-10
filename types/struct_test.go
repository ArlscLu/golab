package types

import (
	"fmt"
	"testing"
)

//组合
type boy struct {
}

func (b *boy) BasketBall() {
	fmt.Println("like play BasketBall")
}
func (b boy) BasketBallCommon() {
	fmt.Println("like play BasketBallCommon")
}

type girl struct {
}

func (g *girl) Dance() {
	fmt.Println("like Dance")
}
func (g girl) DanceCommon() {
	fmt.Println("like  DanceCommon")
}

type pair struct {
	pairName string
	boy      // embedded field T
	*girl    // embedded field *T
}

//S
var s1 = pair{}

//*S
var s2 = new(pair)

func TestStruct1(t *testing.T) {
	fmt.Printf("普通的%+v   \n", s1)
	fmt.Printf("指针的%+v \n", s2)
	// s1.BasketBall()
	// s1.BasketBallCommon()
	// s1.Dance()
	// s1.DanceCommon()
	s2.BasketBall()
	s2.BasketBallCommon()
	s2.Dance()
	s2.DanceCommon()
}

func TestStruct2(t *testing.T) {

}
func TestStruct3(t *testing.T) {

}
func TestStruct4(t *testing.T) {

}
