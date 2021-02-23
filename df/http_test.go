package df

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPost(t *testing.T) {
	url := "http://st1-api.mingqijia.com/v1_goods_search/index"
	fmt.Println("URL:>", url)
	//也可以使用marshall 一个struc map array ....
	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast.","site_id":1}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//自定义header
	//可以定义 cookie authorization
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ENTERPRISE-id", "10000")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
