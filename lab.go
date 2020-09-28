package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// e := gin.Default()
	// // 添加 Get 请求路由
	// e.GET("/", func(context *gin.Context) {
	// 	context.String(http.StatusOK, "hello gin")
	// })
	// e.GET("/u", func(context *gin.Context) {
	// 	context.String(http.StatusOK, func() string {
	// 		return "ganbin"
	// 	}())
	// })

	// e.Run(":8088")
	fmt.Println(thisIsEven(10))
	env := os.Getenv("ENV3")
	envconvert, _ := strconv.ParseInt(env, 10, 0)
	envconvert2 := int(envconvert)
	if thisIsEven(envconvert2) {
		fmt.Printf("%T %s\n", envconvert2, `是even`)
	}
	fmt.Println(`from main`)
	// var delta *string
	var delta string
	flag.StringVar(&delta, `delta`, ``, `测试使用的delta`)
	var d int
	flag.IntVar(&d, `d`, 10, `useage of d`)
	name := flag.String(`name`, `tom`, `name of user`)
	flag.Parse()
	flags := flag.Args()
	fmt.Println(`non-flag command-line arguments. `, flags)
	fmt.Println(delta, d, *name)
}

func thisIsEven(i int) (isEven bool) {

	isEven = false
	if i%2 == 0 {
		isEven = true
	}
	return
}
