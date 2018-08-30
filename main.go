package main

import (
	"github.com/hussar-lang/hussar/cmd"
)

var (
	GitCommit     string
	VersionString string
)

func main() {
	cmd.Setup(VersionString, GitCommit)
	cmd.Execute()
}
