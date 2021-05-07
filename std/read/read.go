package read

import (
	"io"
	"os"
)

var _ readJsonFunc = (readBy1)(nil)
var _ readJsonFunc = (readBy2)(nil)
var _ readJsonFunc = (readBy3)(nil)
var _ readJsonFunc = (readBy4)(nil)

type readJsonFunc interface {
	readJson(path string) ([]byte, error)
}

//利用ioutil.ReadFile直接从文件读取到[]byte中
type readBy1 func(path string) ([]byte, error)

func (r readBy1) readJson(path string) ([]byte, error) {
	return r(path)
}

//先从文件读取到file中，在从file读取到buf, buf在追加到最终的[]byte
type readBy2 func(path string) ([]byte, error)

func (r readBy2) readJson(path string) ([]byte, error) {
	return r(path)
}

//先从文件读取到file, 在从file读取到Reader中，从Reader读取到buf, buf最终追加到[]byte
type readBy3 func(path string) ([]byte, error)

func (r readBy3) readJson(path string) ([]byte, error) {
	return r(path)
}

//读取到file中，再利用ioutil将file直接读取到[]byte中
type readBy4 func(path string) ([]byte, error)

func (r readBy4) readJson(path string) ([]byte, error) {
	return r(path)
}

var rb1 readBy1 = func(path string) ([]byte, error) {
	return nil, nil
}

var rb2 readBy2 = func(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var result []byte
	var buff = make([]byte, 1024)
	for {
		n, err := f.Read(buff)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if err == io.EOF {
			break
		}
		result = append(result, buff[:n]...)
	}
	return result, nil
}
