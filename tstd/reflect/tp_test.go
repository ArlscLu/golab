package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

var _ fmt.Stringer = reflect.TypeOf(10)

func TestBase(t *testing.T) {
	r := reflect.TypeOf(3)
	v := reflect.ValueOf(3)
	fmt.Printf("%+v\n", r)
	fmt.Printf("%+v\n", v)
}
