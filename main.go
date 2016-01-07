package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"unicode/utf8"
    "bytes"

	"github.com/codegangsta/cli"
)

const KEN = "けん"

var suffixes = [3]string{
	"ちゃん",
	"くん",
	"さん",
}

func main() {
	app := cli.NewApp()
	app.Name = "genchankunsan"
	app.Author = "hrysd"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "length, l",
			Value: 140,
			Usage: "Length for kenchankunsan",
		},
	}
	app.Action = func(c *cli.Context) {
		var ken bytes.Buffer 
		ken.WriteString(KEN)

		for utf8.RuneCount(ken.Bytes()) < c.Int("length") {
			rand.Seed(time.Now().UnixNano())
			i := rand.Intn(len(suffixes))
			ken.WriteString(suffixes[i])
		}
		ken.WriteString("…")

		fmt.Println(ken.String())
	}
	app.Run(os.Args)
}
