package main

import (
	"flag"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func main() {

	flag.Parse()

	//If the number of arguments is zero, exit the main() function
	if flag.NArg() == 0 {
		println("give a search query, e.g. \"google hello world\" ")
		return
	}

	s := strings.Join(flag.Args(), "+") //Concatenates the args with '+'
	println("let me google", s)

	cmd := new(exec.Cmd) //Pointer to newly allocated exec.Cmd type

	switch runtime.GOOS {

	case "linux":
		cmd = exec.Command("xdg-open", "https://google.com/#q="+s)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "https://google.com/#q="+s)
	case "darwin":
		cmd = exec.Command("open", "https://google.com/#q="+s)
	default:
		println("I don't know how to google it, on this OS.")
		println("Open an issue at github.com/hariharan-uno/google")

	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

}
