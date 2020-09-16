package self

/*
bytes.Reader  strings.Reader  implements  io.Reader



*/
import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var pathNotExist error = errors.New("不存在的文件")
var pathIsDir error = errors.New("属于目录文件")
var bs = [][]byte{
	[]byte("hello world"),
	[]byte("chinese永char"),
	[]byte(""),
	[]byte("こんにちは!"),
}

//读取文件
func TestRename(t *testing.T) {
	d, err := read("../resource/a.dat")
	if err != nil {
		panic(err)
	}
	finfo, err := d.Stat()
	if err != nil {
		panic(err)
	}
	if finfo.IsDir() {
		panic(pathIsDir)
	}

	// bytes.Reader
	// A Reader implements the io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner interfaces by reading from a byte slice.
	// Unlike a Buffer, a Reader is read-only and supports seeking.
	// The zero value for Reader operates like a Reader of an empty slice.
	bread := bytes.NewReader(bs[3])
	fmt.Println(bread.Len(), bread.Size())

	bbuffer := bytes.NewBuffer(bs[0])
	fmt.Println(bbuffer.Len(), bbuffer.String())
	rename()
}

//打开文件
func read(d string) (*os.File, error) {
	f, err := os.Open(d)
	if err != nil {
		return nil, pathNotExist
	}
	return f, nil
}
func scan(f *os.File) string {
	return ""
}

func rename() {
	fss, err := ioutil.ReadDir("../resource")
	if err != nil {
		panic(err)
	}
	for _, v := range fss {
		fmt.Printf(
			"%+v\n%+v\n%+v\n%+v\n%+v\n%+v\n\n\n",
			v.Name(),
			v.Size(),
			v.Mode(),
			v.ModTime(),
			v.IsDir(),
			v.Sys(),
		)
		if v.Name() == "b.json" {
			content, err := ioutil.ReadFile("../resource/" + v.Name())
			if err != nil {
				panic("读取文件错误")
			}
			fmt.Println(string(content))
		}
	}
}
