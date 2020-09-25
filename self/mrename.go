package self

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/sirupsen/logrus"
)

type Rename struct {
	sourceF        *os.File //source的file
	searchPattern  string   //查找的pattern
	replacePattern string   //替换为pattern
	dir            string   //某个文件下的文件
	pair           []pair   //struct组成的对应关系
}

type pair struct {
	From, To string
}

var oSource = "source.txt"
var oDst = "mid.txt"

var searchPattern = `\d{5}`
var replacePattern = `\d{7}`

//入口
func Run() {
	r := New()
	r.setPatternPair()
}

func New() *Rename {
	re := new(Rename)
	dir, err := os.Getwd()
	if err != nil {
		logrus.Errorf("%s", `获取当前路径发生异常`)
		panic(err)
	}
	re.SetDir(dir)
	re.SetSourceFile(oSource)
	re.SetPattern(searchPattern, replacePattern)
	return re
}

//把对应关系写到pattern里面
func (re *Rename) setPatternPair() {
	reg, err := regexp.Compile(re.searchPattern)
	regrep, err := regexp.Compile(re.replacePattern)
	if err != nil {
		logrus.Errorf("%s", "创建reg出错")
		panic(err)
	}
	lines := re.readFile()
	for _, line := range lines {
		from := string(reg.Find(line))
		to := string(regrep.Find(line))
		re.pair = append(re.pair, pair{
			From: from,
			To:   to,
		})
	}
	for _, v := range re.pair {
		fmt.Printf("%#v\n", v)
	}

}

//读对应关系
func (re *Rename) readFile() [][]byte {
	var lines []byte
	bottle := make([]byte, 10)
	for {
		n, err := re.sourceF.Read(bottle)
		if err != nil && err != io.EOF {
			logrus.Errorf("%s", `读取源文件发生异常`)
			panic(err)
		}
		if n == 0 {
			break
		}
		lines = append(lines, bottle...)
	}
	return bytes.Split(lines, []byte{'\n'})
}

func (re *Rename) rename() {

}

func (re *Rename) SetDir(d string) {
	re.dir = d
}

//设置替换的模式
func (re *Rename) SetPattern(s, r string) {
	re.searchPattern = s
	re.replacePattern = r
}

//设置源文件
func (re *Rename) SetSourceFile(source string) {
	s, err := os.OpenFile(source, os.O_RDONLY, 0777)
	if err != nil {
		logrus.Errorf("%s", "打开源文件发生错误")
		panic(err)
	}
	// defer s.Close()
	re.sourceF = s
}
