package gosyntax

import (
	"fmt"
	"io"
	"strings"
)

type alphaReader struct {
	reader io.Reader
}

func Pio() {
	reader := strings.NewReader("Go语言中文网")
	p := make([]byte, 6)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
	fmt.Println("afafeww")
	fmt.Printf("%s, %d\n", p, n)
}

/*
新建实例
*/
func NewAlphaReader(reader io.Reader) *alphaReader {
	return &alphaReader{reader: reader}
}

/*
只选取字母
*/
func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}
	copy(p, buf)
	return n, nil
}
func Pior() {

}
