package self

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestRead(t *testing.T) {
	f, err := os.OpenFile("bak.txt", os.O_RDONLY, 0666)
	if err != nil {
		logrus.Debugln(err)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		logrus.Debugln(err)
	}
	t.Logf("\n%c\n", data)
	t.Log(f.Name())
}

func TestReadFile(t *testing.T) {
	f, err := ioutil.ReadFile("bak.txt")
	if err != nil {
		logrus.Debugln(err)
	}
	for i, v := range f {
		if i%2 == 0 {
			t.Log(string(v))
		}
	}
}
