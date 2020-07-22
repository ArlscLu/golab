package at

//slide  struct 直接取地址对应的值  0x00对应nil

import (
	"fmt"
	"time"
)

type Mtest struct {
	age uint8
}

func (m Mtest) getAge() uint8 {
	return m.age
}

func (m *Mtest) getAge2() uint8 {
	/*
		if '*' {
			then m=*m
		}
	*/
	return m.age
}

var mm Mtest = Mtest{}

type mTime struct {
	time.Time
}
type ttime = time.Time

var somett = ttime{}

type xxx struct {
}

func (t mTime) first3Chars() string {
	return t.String()[0:3]
}

func ShowMtest() {
	s := Mtest{12}
	m := &Mtest{10}
	fmt.Println(s.getAge())
	fmt.Println(s.getAge2())
	fmt.Println(m.getAge())
	fmt.Println(m.getAge2())
	mt := mTime{
		// time.Now(),
	}
	fmt.Println("full time now", mt.String())
	fmt.Println("first 3 chars", mt.first3Chars())
	fmt.Println(time.Time.String(time.Now()))
	fmt.Println(somett.String())
	fmt.Println(111)
	// fmt.Println(somett)
	// fmt.Println(mm)
	var xxx2 *xxx
	fmt.Println(xxx2)
	fmt.Println(new(xxx))
}
func ShowMtest2() {
	s4 := new(sl1)
	s5 := new(ar)
	s6 := new(st2)

	s7 := &st2{"tom"}  //&{tom}
	s8 := &ar{2, 3, 4} //&[2 3 4]
	fmt.Println(s1, s2, s3, s4, s5, s6, s7, &s7, s8, &s8)
}

type sl1 []int
type ar [3]int
type st2 struct {
	name string
}

var s1 sl1
var s2 *ar
var s3 *st2

/*
	任意 *addr


	typeofaddr = *addr

	if(*addr == point) {
		if(*add)== nil {
			return nil
		} else {
			return *addr
		}
	} else {    //no nil
		return *addr
	}





*/
