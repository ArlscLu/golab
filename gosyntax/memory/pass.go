package memory

var (
	V0 int = 8
	V1     = []int{1, 2, 3}
)

//passed by value
func Fpass(a int) *int {
	return &a //different  hence: pass by value
}
func FpassRefe(a []int) *int {
	return &a[0]
}
