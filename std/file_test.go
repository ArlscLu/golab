package std

/*
	os.file 测试
	实现了 io.reader  io.writer
*/

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/arlsclu7/golab/third"
)

// var filename = "../resource/a.dat"
var filename = "../all.log"

func TestFile(t *testing.T) {
	f, err := open(filename)
	if err != nil {
		panic(err)
	}
	third.Log.Debug(f.Name())
	in := make([]byte, 200)
	var sp [][]byte
	// out := make([]byte, 5)
	for {
		_, err = f.Read(in)
		if err != nil && err != io.EOF {
			panic("发生了除eof的错误")
		}
		if err == io.EOF {
			break
		}
		bs := bytes.Split(in, []byte{'\n'})
		sp = append(sp, bs...)
	}
	for _, v := range sp {
		os.Stdout.WriteString(string(v))
	}

}

func open(fn string) (*os.File, error) {
	f, err := os.OpenFile(fn, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	return f, nil
}
