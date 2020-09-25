package std_test

import (
	"io"
	"os"
	"testing"
)

//读取源
var filename = "./osfile.log"

var F *os.File

//输出
var out1 = os.Stdout

var bottle = make([]byte, 100)

func TestOsFile(t *testing.T) {

	for {
		n, err := F.Read(bottle)
		if err != nil && err != io.EOF {
			t.Error(`读到bottle发生异常`, err)
			panic(err)
		}
		if n == 0 {
			break
		}
		// t.Logf("%s", bottle)
	}
	//重置file index
	F.Seek(0, 0)
	t.Log(countSomeAlpha())
	F.Seek(0, 0)
}

func countSomeAlpha() (int, error) {
	bottle := make([]byte, 100)
	var count int
	for {
		n, err := F.Read(bottle)
		if err != nil && err != io.EOF {
			return count, nil
		}
		if n == 0 {
			break
		}
		for _, v := range bottle {
			if lineHasSomeAlpha(v) {
				count++
			}
		}
	}
	return count, nil
}

func writeTo(b []byte) (int, error) {
	return 0, nil
}
func lineHasSomeAlpha(b byte) bool {
	if b == 'a' {
		return true
	}
	return false
}

func init() {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	//package file now has value
	F = f
}
