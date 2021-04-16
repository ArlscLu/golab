package ts

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"testing"
)

func TestJsonUnmarshal(t *testing.T) {
	testCases := []struct {
		desc string
		j    string
	}{
		{
			desc: "",
			j: `
			{
				"name" : "testJSON",
				"servers" : [
					{
						"serverName" : "Shanghai_VPN",
						"serverIP" : "127.0.0.1"
					},
					{
						"serverName" : "Beijing_VPN",
						"serverIP" : "127.0.0.2"        
					}
				],
				"status" : 0
			}
			`,
		},
		{
			desc: "空json null",
			j:    `null`,
		},
		{
			desc: "数组",
			j:    `[1,2,3,4,5]`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var result interface{}
			_ = json.Unmarshal([]byte(tC.j), result)
			t.Logf("%v", result)
		})
	}
}

//json.decode{}
func TestJsonDecoder(t *testing.T) {
	const jsonStream = `
	{"nn": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name string `json:"nn"`
		Text string `json:"text,omitempty"`
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
