package types

import (
	"errors"
	"os"
	"testing"
)

func TestNewError(t *testing.T) {
	_, err := square(-1)
	// if err != nil {
	// 	t.Log(err)
	// }
	// t.Log(rst)
	if e, ok := err.(*os.PathError); ok {
		t.Logf("%s", e)
	}
}

func square(f float64) (float64, error) {
	if f < 0 {
		// return 0, errors.New("not smaller than 0")
		return 0, &os.PathError{
			Op:   "square",
			Path: "/types/error_test.go",
			Err:  errors.New("not smaller than 0"),
		}
	}
	return f / 2, nil
}
