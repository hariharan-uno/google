package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "google"
	app.Usage = "Quick Search from terminal on google"
	app.Action = func(c *cli.Context) {
		if len(c.Args()) == 0 {
			println("give a search query, e.g. \"google hello world\" ")
		} else {
			s := strings.Join(c.Args(), "+")
			println("let me google", s)

			cmd := new(exec.Cmd)

			switch runtime.GOOS {

			case "linux":
				cmd = exec.Command("xdg-open", "https://google.com/#q="+s)
			case "windows":
				cmd = exec.Command("cmd", "/c", "start", "https://google.com/#q="+s)
			case "darwin":
				cmd = exec.Command("open", "https://google.com/#q="+s)
			default:
				cmd = exec.Command("xdg-open", "https://google.com/#q="+s)
				//TODO: Add support for more Operating Systems
			}

			err := cmd.Start()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	app.Run(os.Args)
}
