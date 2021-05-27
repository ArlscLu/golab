package tthird

import (
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestPing(t *testing.T) {
	rOption := redis.DialDatabase(1)
	c, err := redis.Dial("tcp", ":6379", rOption)
	if err != nil {
		t.Log(err)
	}
	defer c.Close()
	_, err = c.Do("SET", "mk01", "----chinese---")
	n, err := c.Do("GET", "mkey")

	t.Log(n.([]byte))
}
