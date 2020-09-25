package std_test

/*
	Find(All)?(String)?(Submatch)?(Index)?
*/
import (
	"fmt"
	"io"
	"os"
	"regexp"
	"testing"
)

var pattern = `local`

var filename1 = `./d.txt`

var bs []byte

var reg *regexp.Regexp

func TestRegexp(t *testing.T) {
	//查找index组成的[]int
	fmt.Println(findIndex(reg, bs), string(bs[76:81]))
}

func findIndex(r *regexp.Regexp, s []byte) []int {
	return r.FindIndex(s)
}

func init() {
	createReg()
	setFind()
}

//初始化reg
func createReg() {
	r, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	reg = r
}

//文件源
func setFind() {
	f, err := os.OpenFile(filename1, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	cache := make([]byte, 10)
	for {
		n, err := f.Read(cache)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		bs = append(bs, cache...)
	}
}
