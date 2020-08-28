package std

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	logredis "github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
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

var (
	f   *os.File
	err error
)

func ExecFilePkg() {
	f, err = os.Create("a1.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println("name is ", f.Name())
	err := f.Chmod(0777)
	if err != nil {
		panic("权限配置error")
	}
	fi, err := f.Stat()
	if err != nil {
		panic("fileinfo获取失败")
	}
	/*
		type FileInfo interface {
			Name() string       // base name of the file
			Size() int64        // length in bytes for regular files; system-dependent for others
			Mode() FileMode     // file mode bits
			ModTime() time.Time // modification time
			IsDir() bool        // abbreviation for Mode().IsDir()
			Sys() interface{}   // underlying data source (can return nil)
		}
	*/
	fmt.Println(
		"base name of the file\n", fi.Name(),
		"length in bytes for regular files; system-dependent for others\n", fi.Size(),
		"file mode bits\n", fi.Mode(),
	)
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		fmt.Println("regular file")
	case mode.IsDir():
		fmt.Println("directory")
	case mode.Perm().IsDir():

	}
}

var now time.Time = time.Now()

//one day one logfile
func todayFile() *os.File {
	filename := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
	filename += ".log"
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal("创建日志文件发生错误", err)
	}
	return f
}

func Log() {
	f := todayFile()
	// l := log.New(f, "[user-service] ", log.LstdFlags)
	// lrus := logrus.New()
	// lrus.SetOutput(f)
	// lrus.Println(111)
	// logrus.Println(222)
	// l.Printf("%s", "helloworld")
	lrus1 := &logrus.Logger{
		Out:       f,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
	hookConfig := logredis.HookConfig{
		Host:     "localhost",
		Key:      "mykey",
		Format:   "v0",
		App:      "aweosome",
		Hostname: "localhost",
		TTL:      3600,
		Port:     6379,
	}
	hook, err := logredis.NewHook(hookConfig)
	if err != nil {
		// lrus1.Errorf("%s", err)
		panic(err)
	} else {
		lrus1.AddHook(hook)
	}
	// lrus1.SetFormatter(&logrus.JSONFormatter{})
	le := lrus1.WithFields(logrus.Fields{
		"name": "dj",
		"age":  18,
	})
	le.Debug("from debugger debug level")
	le.Info("from debugger info level")
}
