package retrieve_test

import (
	"fmt"
	"testing"
)

//TestMain 测试主函数
func TestMain(t *testing.T) {
	g := new(GiftCard)
	fmt.Println(g.Enable())
}

type Status struct {
	Name       string
	EnableAble string
}

func (s *Status) Enable() string {
	s.Name = s.Name + "111"
	return s.Name
}
func (s *Status) Disable() {

}
func (s *Status) UpList() {

}
func (s *Status) DownList() {

}

type GiftCard struct {
	Status
}

func gtest() {

}
