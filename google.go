package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "google"
	app.Usage = "Quick Search on google"
	app.Action = func(c *cli.Context) {
		if len(c.Args()) == 0 {
			println("give a search query")
		} else {
			println("let me google", c.Args()[0])
		}
	}
	app.Run(os.Args)
}
