package self

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	hour := flag.Int(`hour`, 18, `print the day target hour 24h`)
	second := flag.Int(`second`, 0, `print the day target hour 24h`)

	flag.Parse()
	location, err := time.LoadLocation(`Local`)
	if err != nil {
		panic(err)
	}
	now := time.Now()
	target := time.Date(now.Year(), now.Month(), now.Day(), *hour, *second, 0, 0, location)

	d := target.Sub(now)
	fmt.Printf("还剩余%.0f\n", d.Minutes())
	fmt.Printf("%s\n", d)
	fmt.Println(location.String())
}
