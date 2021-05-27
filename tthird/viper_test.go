package tthird

import (
	"testing"

	"github.com/spf13/viper"
)

func TestX(t *testing.T) {
	viper.SetDefault("name", "tom")

	t.Log(viper.Get("name"))
	t.Parallel()
}
