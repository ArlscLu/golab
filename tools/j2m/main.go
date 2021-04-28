package j2m

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/urfave/cli"
	"github.com/yosuke-furukawa/json5/encoding/json5"
)

func stringFlag(name, value, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

// json转换为md文档
var Tool = cli.Command{
	Name:   "j2m",
	Usage:  "json转换为md文档",
	Action: run,
	Flags: []cli.Flag{
		stringFlag("json,j", "", "json文件"),
		stringFlag("pattern,p", "", "模版文件"),
	},
}

func run(ctx *cli.Context) {
	jsonFile := ctx.String("json")
	j, err := rb2.readJson(jsonFile)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	if len(j) == 0 {
		fmt.Printf("未读取任何数据")
		os.Exit(1)
	}
	patternFile := ctx.String("pattern")
	pattern, err := rb2.readJson(patternFile)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	if len(pattern) == 0 {
		fmt.Printf("未读取pattern任何数据")
		os.Exit(1)
	}
	var s handleStruct
	err = json5.Unmarshal(j, &s)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	var in []string
	err = handle(&in, "", s.In, 0)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	var out []string
	err = handle(&out, "", s.Out, 0)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	patternStr := string(pattern)
	inOrl, _ := json5.MarshalIndent(s.In, "", "	")
	outOrl, _ := json5.MarshalIndent(s.Out, "", "	")
	fmt.Printf(patternStr, s.Topic, s.Url, strings.Join(in, "\n"), strings.Join(out, "\n"), s.Url, string(inOrl), string(outOrl))
}

//依赖类型
//需要key值产生一条md记录
// 根据type对value值做进一步处理
func handle(result *[]string, key string, data interface{}, depth int) error {
	if data == nil {
		return nil
	}
	switch value := data.(type) {
	case map[string]interface{}:
		makeOneLine(result, key, "object", depth)
		if len(key) != 0 {
			depth++
		}
		for k, v := range value {
			handle(result, k, v, depth)
		}
	case []interface{}:
		makeOneLine(result, key, "array", depth)
		if len(key) != 0 {
			depth++
		}
		handle(result, "", value[0], depth)
	case bool:
		makeOneLine(result, key, "boolean", depth)
	case float64:
		makeOneLine(result, key, "integer", depth)
	case string:
		makeOneLine(result, key, "string", depth)
	}
	return nil
}

//产生一条md记录
var format = `| %s%s | %s | Yes | %s |`

func makeOneLine(result *[]string, key, keyType string, depth int) {
	if len(key) == 0 {
		return
	}
	remark := key + "的注释"
	var prefix []string
	for i := 0; i < depth; i++ {
		prefix = append(prefix, "↳")
	}
	add := fmt.Sprintf(format, strings.Join(prefix, ""), key, keyType, remark)
	*result = append(*result, add)
}

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

type handleStruct struct {
	Topic string      `json:"topic,omitempty"`
	Url   string      `json:"url,omitempty"`
	In    interface{} `json:"in,omitempty"`
	Out   interface{} `json:"out,omitempty"`
}
