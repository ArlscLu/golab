package std

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var url = "https://wangchujiang.com/linux-command/list.html#!kw=ps"

func TestHttp(t *testing.T) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", body)
}
