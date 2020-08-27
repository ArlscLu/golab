package std

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("【UserCenter】")
}

type rst struct {
	num  int
	Time time.Duration
}
type srst []rst

func (s *srst) ExOutput() {
	//will change
	for k := range *s {
		(*s)[k].Exchange()
	}
	//won't change
	// for _, i := range *s {
	// 	i.Exchange()
	// }
}
func (s *rst) Exchange() {
	s.num += 0
}
func RunFile() {
	filename := "go.mod"
	s1 := make(srst, 3)
	//第一种方式 os.Open
	s1[0] = osopen(filename)
	//第二种方式 ioutil
	s1[1] = osioutil(filename)
	//第三种方式 bufio
	s1[2] = osbufio(filename)
	s1.ExOutput()
	fmt.Println(s1)
}

func osopen(filename string) rst {
	start := time.Now()
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf := make([]byte, 10)
	num := 0
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		num += n
	}
	end := time.Now()
	return rst{num, end.Sub(start)}
}

func osioutil(filename string) rst {
	start := time.Now()
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = ioutil.ReadAll(file)
	end := time.Now()
	return rst{0, end.Sub(start)}
}

func osbufio(filename string) rst {
	start := time.Now()
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bufReader := bufio.NewReader(file)
	buf := make([]byte, 1024)
	num := 0
	for {
		readNum, err := bufReader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == readNum {
			break
		}
		num += readNum
	}
	end := time.Now()
	return rst{num, end.Sub(start)}
}
