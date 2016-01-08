package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kenchankunsan/genchankunsan/internal/kenchankunsan"
)

func main() {
	app := cli.NewApp()
	app.Name = "genchankunsan"
	app.Author = "hrysd"
	app.Usage = "Generate kenchankunsan"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "length, l",
			Value: 140,
			Usage: "Length for kenchankunsan",
		},
	}
	app.Action = func(c *cli.Context) {
		ken := kenchankunsan.Kenchankunsan(c.Int("length"))
		fmt.Println(ken)
	}
	app.Run(os.Args)
}
