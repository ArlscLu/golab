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
	flags := flag.Args()
	fmt.Println(flags)
}

func thisIsEven(i int) (isEven bool) {

	isEven = false
	if i%2 == 0 {
		isEven = true
	}
	return
}
