package cmd

import (
	"fmt"
	"os"

	"github.com/hussar-lang/hussar/repl"

	"github.com/spf13/cobra"
)

var interactive = &cobra.Command{
	Use:   "interactive",
	Short: "Start the interactive shell",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Hussar interactive interpreter.")
		repl.Start(os.Stdin, os.Stdout)
	},
}
