package std

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/sirupsen/logrus"
)

var fp = "../resource/b.json"
var fp1 = "../ /resource//../../filepath"

var vlog = logrus.New()

func TestFilePath(t *testing.T) {
	abs, err := filepath.Abs(fp)
	check(err)
	vlog.Infoln(abs)
	vlog.Infoln(filepath.Base(fp))
	vlog.Infoln(filepath.Ext(fp))
	vlog.Infoln(filepath.FromSlash(fp1))
	matched, err := filepath.Match("json", fp)
	if matched {
		fmt.Println("matched!")
		check(err)
	} else {
		fmt.Println(" not matched!")
	}
	check(err)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func init() {
	i, err := os.OpenFile("std.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	check(err)
	vlog.SetOutput(i)
	fmt.Println("默认的输出level")
}
