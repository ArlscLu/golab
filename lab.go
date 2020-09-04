package main

import (
	"os"

	"github.com/arlsclu7/golab/db"
	"github.com/arlsclu7/golab/std"
	"github.com/arlsclu7/golab/types"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

const (
	dsn = "root:123@(127.0.0.1)/test?charset=utf8&parseTime=True&loc=Local"
)

func main() {
	os.Exit(3)
	std.Main()
	os.Exit(2)
	std.ShowSlice(std.Class1)
	os.Exit(2)
	std.RunChan()
	os.Exit(1)
	db.RunXorm()
	os.Exit(1)
	std.Log()
	os.Exit(1)
	std.ExecFilePkg()
	os.Exit(7)
	std.RunFile()
	os.Exit(1)
	std.RunOs()
	os.Exit(1)
	std.RunLog()
	os.Exit(3)
	std.RunWg()
	os.Exit(2)
	types.RunChan()
	os.Exit(2)
	std.RunLog()
	os.Exit(1)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/post", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"code":    1,
			"data":    "",
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080

	// start := time.Now()
	// a := 10
	// if a > 5 {
	// 	fmt.Println("bigger than 5")
	// } else if a > 7 {
	// 	fmt.Println("bigger than 7")
	// } else {
	// 	fmt.Println("other")
	// }
	// year, month, day := start.Date()
	// l := start.Location().String()
	// fmt.Println(year, month, day, l)
	// for i := 0; i < 100; {
	// 	fmt.Print(".")
	// 	i++
	// }
	// t := time.Now()
	// used := t.Sub(start)
	// fmt.Println(`使用了`, "\n", used)
	// fmt.Printf("%T\n", used)
	// s := "helloworld为什么gans"
	// fmt.Println(strings.Index(s, "什"), s[0])
	// fmt.Printf("%c%c%c%c\n", strings.IndexRune(s, '什'), s[13], s[1], s[0])
	// // fmt.Println(rand.New("1235677899").Intn(100))
	// os.Exit(1)
	// db, err := gorm.Open("mysql", dsn)
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()

	// // Migrate the schema
	// db.AutoMigrate(&Product{})

	// // Create
	// db.Create(&Product{Code: "L1212", Price: 1000})

	// // Read
	// var product Product
	// db.First(&product, 1)                   // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212

	// // Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)

	// // Delete - delete product
	// db.Delete(&product)
}
