package main

import (
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "hello",
		Usage: "hello world example",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Usage:   "preferred language",
				Value:   "english",
				Aliases: []string{"languange", "l"},
			},
			&cli.BoolFlag{
				Name:  "show",
				Usage: "wheter show",
			},
		},
		Action: func(c *cli.Context) error {
			name := "somebody"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if c.Bool("show") {
				fmt.Println("hello world", name)
			}
			fmt.Println("hello world", name)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
