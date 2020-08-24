package std

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

type interval []float64

func (i *interval) String() string {
	r := ""
	for _, v := range *i {
		r = fmt.Sprintf("%s", "使用了") + r + fmt.Sprintf("%f秒\n", v)
	}
	return r
}
func (i *interval) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("interval value alerady set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration.Seconds())
	}
	return nil
}

var intervalFlag interval

var period time.Duration

var gopherType string

const (
	defaultGopher = "pocket"
	usage         = "the variety of gopher"
)

func Runflag() {
	//e1
	var species = flag.String("species", "gopher", "the species we are studying")

	//e2

	// flag.Parse()
	if flag.Parsed() {
		fmt.Println(`111`)
	}
	fmt.Println("total args", flag.NArg()-1)
	fmt.Println(species, flag.Args())

	//e3
}

func init() {

	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")

	flag.Var(&intervalFlag, "delta", "comma-separated list of intervals to use between events")
}
