package std

import (
	"fmt"
	"net/url"

	"github.com/sirupsen/logrus"
)

type st uint8

func (*st) String() string {
	return "stname"
}
func (*st) Error() string {
	return "i am error"
}

var s1 *st

func Runst() {
	fmt.Printf("%s\n", s1)
	fmt.Println(s1)
	fmt.Println(s1.String())

	var e error = s1
	func(ee error) {
		fmt.Println(ee)
	}(e)
	u1 := "www.baidu.com?user=tom&age=18"
	uu1, err := url.Parse(u1)
	if err != nil {
		logrus.Info(err)
	}
	fmt.Printf("%+v", uu1)
	fmt.Println("user", uu1.User)
	fmt.Println("scheme", uu1.Scheme)
	fmt.Println("host", uu1.Host)
}
