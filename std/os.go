package std

import "os"

func RunOs() {
	c := "helloworld"
	cb := []byte(c)
	os.Stdout.Write(cb)
}
