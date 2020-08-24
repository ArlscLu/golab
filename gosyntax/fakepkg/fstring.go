package fakepkg

import (
	"errors"
	"fmt"
	"strings"
)

var (
	Filename string
)

func Enterfstring() {
	rst := make(map[string]interface{}, 10)
	rst["compare"] = strings.Compare("a", "b")
	rst["ContainsRune"] = strings.ContainsRune("helloworld电风扇", '电')
	rst["count"] = strings.Count("helloworld电风扇电电", "电")

	for k, v := range rst {
		fmt.Println(k, v)
	}
	err := errors.New("test error")
	if err != nil {
		panic("eror")
	}
	defer func() {
	}()
}

func comp(a, b string) int {
	return strings.Compare(a, b)
}
