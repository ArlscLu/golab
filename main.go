package main

import (
	"os"

	"github.com/arlsclu7/golab/tools/g2s"
	"github.com/arlsclu7/golab/tools/j2m"
	"github.com/urfave/cli"
)

func main() {
	var AppVer string = "v1.0.0"
	app := cli.NewApp()
	app.Name = "tools for mqj"
	app.Usage = "小工具"
	app.Version = AppVer
	app.Author = "cjs"
	app.Email = "42282367@qq.com"
	app.Commands = []cli.Command{
		g2s.Tool,
		j2m.Tool,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
