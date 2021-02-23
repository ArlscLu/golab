package df

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateEmptyFile(t *testing.T) {
	file := "new/1.md"
	emptyFile, err := os.Create(file)
	if err != nil {
		t.Log(err)
	}
	t.Log(emptyFile.Name())
	emptyFile.Close()
}

func TestCreateEmptyFolder(t *testing.T) {
	dir := "new"
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dir, 0755)
		if errDir != nil {
			t.Fatal(err)
		}
	}
}

func TestStat(t *testing.T) {
	find := "3.md"
	s, err := os.Stat(find)
	if os.IsNotExist(err) {
		t.Log("不存在", find)
		new, err := os.Create(find)
		if err != nil {
			t.Fatal("创建 ", find, `失败`)
		}
		s, _ = new.Stat()
	}
	t.Log(s.Name())
}

func TestAssert(t *testing.T) {
	var x I = ison{Name: "tom"}
	if is, ok := x.(ison); ok {
		t.Log(is.Name)
	}
	if is, ok := x.(II); ok {
		is.m()
	}
	switch v := x.(type) {
	case ison:
		t.Log(v.Name + "elllo")
	default:
		t.Fatal("fatelxxxx")
	}
}

type I interface {
	m()
	mm()
}
type II interface{ m() }
type ison struct {
	Name string
}

func (i ison) m() {
	fmt.Print("\nxxxxyyy\n")
}
func (i ison) mm() {}

// func f(y I) {
// 	_ := y.(ison)      // illegal: string does not implement I (missing method m)
// 	_ := y.(io.Reader) // r has type io.Reader and the dynamic type of y must implement both I and io.Reader
// }
