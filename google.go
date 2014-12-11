// Copyright 2014 Hari haran. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

func main() {

	flag.Parse()

	// If the number of arguments is zero, exit the main() function.
	if flag.NArg() == 0 {
		fmt.Println(`give a search query, e.g. "google hello world" `)
		return
	}

	s := strings.Join(flag.Args(), "+") // concatenate the args with '+'
	fmt.Println("let me google", s)

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
