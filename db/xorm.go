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
	// err = engine.Sync2(new(User))
	s.Infof("%+v", engine.Logger().Level)
	if err != nil {
		s.Fatal(err)
	}
	user1 := &User{
		Name: "Tom",
		Age:  18,
	}
	affected, err := engine.Insert(user1)
	if err != nil {
		s.Panicln("发生了插入错误", err)
	}
	s.Printf("insert %d rows", affected)
	has, _ := engine.ID(5).Get(&User{})
	if has {
		fmt.Printf("user1: %v\n", user1)
	}
}
