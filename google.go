// Copyright 2014 Hari haran. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [query]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "example: %s who is yoda?\n\n", os.Args[0])
		os.Exit(2)
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
	}

	fmt.Printf("Googling %q\n", strings.Join(flag.Args(), " "))
	fmt.Printf("A browser window should open. If a window is already open, please visit it!\n\n")

	// concatenate the args with '+' to append them to URL
	s := strings.Join(flag.Args(), "+")

	link, err := url.Parse("https://google.com/#q=" + s)
	if err != nil {
		fmt.Printf("url parsing error: %s\n", err)
		return
	}

	cmd := new(exec.Cmd)

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", link.String())
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", link.String())
	case "darwin":
		cmd = exec.Command("open", link.String())
	default:
		fmt.Println("I don't know how to open a browser on this OS.")
		fmt.Println("Open an issue at https://github.com/hariharan-uno/google/issues")
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
