package applog

import (
	"log"
	"os"
)

var fname string = "/usr/local/var/log/golab/go.log"

func Applog() *log.Logger {
	f, err := os.OpenFile(fname, os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal("创建文件发生了异常")
	}
	logger := log.New(f, "[golog]...", log.Ldate|log.Ltime)
	return logger
}
