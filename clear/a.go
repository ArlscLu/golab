package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type f struct {
	N  int
	TT time.Time
}

func main() {

	t := f{
		5,
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
	}

	m, _ := json.Marshal(t)
	fmt.Printf("%s \n", m)

	fn := ff{10, InsideStruct{"world"}}
	xx, _ := json.Marshal(fn)
	fmt.Printf("%s \n", xx)
}

type ff struct {
	NN int
	InsideStruct
}
type InsideStruct struct {
	Name string `json:"name"`
}

func (is *InsideStruct) MarshalJSON() ([]byte, error) {
	return []byte("yyyy"), nil
}

var _ json.Marshaler = &InsideStruct{}
var _ json.Marshaler = time.Time{}

// var _ json.Marshaler = f{}
var _ json.Marshaler = &ff{}
