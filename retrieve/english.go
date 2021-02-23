package retrieve

import (
	"fmt"
)

type alpha struct {
}

var TableName string = `alpha`

//List show all
func List() {
	fmt.Println(TableName)
}
