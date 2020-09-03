package db

import (
	"fmt"
	"time"

	"github.com/arlsclu7/golab/std"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
type Police struct {
	Id   int64
	Name string
	Age  int
	Uid  int
}

var l *logrus.Logger

var s = logrus.New()

// var ss = &logrus.Logger{
// 	Out:       std.TodayFile(),
// 	Formatter: new(logrus.TextFormatter),
// 	Hooks:     make(logrus.LevelHooks),
// 	Level:     logrus.DebugLevel,
// }

func RunXorm() {
	s.SetOutput(std.TodayFile())
	s.SetLevel(logrus.DebugLevel)
	engine, err := xorm.NewEngine("mysql", "root:123@(127.0.0.1)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		s.Info(11)
	}
	defer engine.Close()
	user1 := &User{
		Name: "Gandof",
	}
	has, _ := engine.Get(user1)
	if has {
		fmt.Printf("user4: %v\n", user1)
	}
	users := make([]User, 0)
	engine.Where("age > ? and age <=?", 6, 16).Find(&users)
	for _, v := range users {
		fmt.Printf("%+v\n", v)
	}
}

func Search(id int) *User {
	engine, err := xorm.NewEngine("mysql", "root:123@(127.0.0.1)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		s.Info(11)
	}
	defer engine.Close()
	find := &User{}
	has, _ := engine.ID(id).Get(find)
	if has {
		return find
	}
	return nil
}
