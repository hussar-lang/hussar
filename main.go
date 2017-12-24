package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/kscarlett/kmonkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the kmonkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type commands!\n")
	repl.Start(os.Stdin, os.Stdout)
}
