package tthird

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	dir := filepath.Dir("logrus.go")
	f, err := os.OpenFile(
		dir+string(os.PathSeparator)+".."+string(os.PathSeparator)+"all.log",
		os.O_APPEND|os.O_CREATE|os.O_RDWR,
		0777,
	)
	if err != nil {
		panic(err)
	}
	Log.SetOutput(f)
	Log.SetLevel(logrus.DebugLevel)
}
