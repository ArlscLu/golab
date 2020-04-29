package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func mains() {

	url := "http://st1-api.handeson.com/v1_goods_detail/index"
	method := "POST"

	payload := strings.NewReader("{\n \"spu_id\": 14296,\n \"activity_type\":\"vip\",\n \"token\":\"ce5444b0-32d1-11ea-8f84-a37164b398e7\"\n}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Type", "text/plain")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
