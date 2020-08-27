package std

import (
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	errorF, err := os.OpenFile("a.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("%s", err)
	}
	Info = log.New(os.Stdout, "Info ", log.Ldate|log.Ltime)
	Warning = log.New(os.Stdout, "Warning ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errorF), "Error ", log.Ldate)
}

func RunLog() {
	Info.Printf("\n%b\n%b\n%b\n", os.O_CREATE, os.O_WRONLY, os.O_APPEND)
	Info.Println("飞雪无情的博客:", "http://www.flysnow.org")
	Warning.Printf("飞雪无情的微信公众号：%s\n", "flysnow_org")
	Error.Println("欢迎关注留言")
	// log.SetPrefix("enterprise_id:10000")
	// log.Printf("%s", "from log")

	// buf := &bytes.Buffer{}
	// logger := log.New(buf, "enterprise-id", log.Lshortfile|log.LstdFlags)
	// logger.Printf("%s%d", "asr", 12)
	// logrus.Printf("%d", 111)
	// logrus.Info(os.Args)
	// fmt.Print(buf.String())

}
