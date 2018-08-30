package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var env = &cobra.Command{
	Use:   "env",
	Short: "Displays the current environment setup used by Hussar",
	Run: func(cmd *cobra.Command, args []string) {
		printEnvironment()
	},
}

func printEnvironment() {
	fmt.Printf("arch:  %s", runtime.GOARCH)
	fmt.Printf("os:    %s", runtime.GOOS)
	fmt.Printf("bin:   %s", os.Args[0])
	fmt.Printf("gc:    %s", runtime.Version())
	fmt.Printf("vers:  %s", version)
	fmt.Printf("build: %s", build)
}
