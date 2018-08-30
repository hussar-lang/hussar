package cmd

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"

	"github.com/hussar-lang/hussar/evaluator"
	"github.com/hussar-lang/hussar/lexer"
	"github.com/hussar-lang/hussar/object"
	"github.com/hussar-lang/hussar/parser"
)

var run = &cobra.Command{
	Use:   "run",
	Short: "Run the given script",
	Run: func(cmd *cobra.Command, args []string) {
		runFromFile()
	},
}

func init() {
	run.Flags().String("source.file", "", "the source file to run")
	viper.BindPFlags(run.Flags())
}

func runFromFile() {
	file, err := getSourceFile(viper.GetString("source.file"))
	if err != nil {
		log.Fatal(err)
	}

	l := lexer.New(string(file))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(os.Stderr, p.Errors())
		os.Exit(1)
	}

	env := object.NewEnvironment()
	eval := evaluator.Eval(program, env)
	if eval.Inspect() != "NULL" {
		fmt.Println(eval.Inspect())
	}
}

func getSourceFile(sourceFile string) ([]byte, error) {
	if sourceFile == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.WithError(err).Fatal()
		}

		var files []string
		filepath.Walk(cwd, func(path string, f os.FileInfo, _ error) error {
			if !f.IsDir() {
				if filepath.Ext(path) == "hss" {
					files = append(files, f.Name())
				}
			}
			return nil
		})

		if len(files) == 0 {
			log.WithError(errors.New("no source files found in current directory")).Fatal()
		} else if len(files) > 1 {
			for _, f := range files {
				if strings.ToLower(f) == "main.hss" {
					sourceFile = f
					break
				}
			}

			// Hack, I know
			if sourceFile == "" {
				log.WithError(errors.New("no main source file found in current directory - specify one with the source.file flag")).Fatal()
			}
		} else {
			sourceFile = files[0]
		}
	}

	source, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	return source, nil
}

func printParserErrors(out io.Writer, errors []string) {
	errColor := chalk.Red.NewStyle().WithTextStyle(chalk.Bold).Style

	io.WriteString(out, errColor("PARSER ERROR!\n"))
	for _, msg := range errors {
		io.WriteString(out, errColor("  [!] ")+msg+"\n")
	}
}
