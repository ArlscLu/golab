/**
Each type T has an underlying type:
If T is one of the predeclared boolean, numeric, or string types, or a type literal, the corresponding underlying type is T itself.
Otherwise, T's underlying type is the underlying type of the type to which T refers in its type declaration.
*/
package types

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Ta struct {
	name string
}

type a struct {
	name string
}
type addr struct {
	province uint8
	city     uint8
	district uint8
	home     string
}
type addrList []addr

var addr1 addr = addr{}
var addr2 addr = addr{}

var al1 addrList = addrList{addr1, addr2}

type aa a

var a1 = a{
	name: "jack",
}
var aa1 = aa{
	name: "mary",
}

func mint() {
	a := 10
	b := 0b10101
	c := '中'
	d := 0xdeadbeef
	e := []int{1, 2, 3}
	fmt.Println(a, b)
	fmt.Printf("%c,%[1]o \n", c)
	fmt.Printf("%c,%#[1]x  \n", d)
	fmt.Printf("%p \n", e)
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color"`
	Actors []string
}

var movies = []Movie{
	{Title: "Catereing", Year: 1942, Color: true, Actors: []string{"Tom", "rOSE"}},
	{Title: "Catereing", Year: 1942, Color: false, Actors: []string{"中国移动", "rOSE"}},
}

func Show() {
	data, err := json.MarshalIndent(movies, "", "   ")
	if err != nil {
		log.Fatalf("Json  Marshaling failder: $s", err)
	}
	fmt.Printf("%s \n ", data)
	mint()
	fmt.Println(a1, aa1)
	fmt.Println(addr1, addr2, al1)
	fmt.Println(os.Getenv("ZSH"))
}
