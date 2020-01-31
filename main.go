package main

import (
	"github.com/hussar-lang/hussar/cmd"
)

var (
	version = "dev"
	commit  = "none"
)

func main() {
	cmd.Setup(version, commit)
	cmd.Execute()
}
