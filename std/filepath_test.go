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

var log = logrus.New()

func TestFilePath(t *testing.T) {
	abs, err := filepath.Abs(fp)
	check(err)
	log.Infoln(abs)
	log.Infoln(filepath.Base(fp))
	log.Infoln(filepath.Ext(fp))
	log.Infoln(filepath.FromSlash(fp1))
	if matched, err := filepath.Match("json", fp); matched {
		fmt.Println("matched!")
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
	log.SetOutput(i)
	fmt.Println("默认的输出level", log.Level)
}
