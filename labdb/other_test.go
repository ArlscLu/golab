package labdb

import (
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

//"github.com/gomodule/redigo/redis"
func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}

func TestPackageName(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
	}{
		{
			desc:  "",
			input: "github.com/gomodule/redigo/redis",
		},
		{
			desc:  "",
			input: "github.com/gomodule/redigo/redis.v1",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := getPackageName(tC.input)
			logrus.Print(result)
		})
	}
}
