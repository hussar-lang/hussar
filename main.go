package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/hussar-lang/hussar/evaluator"
	"github.com/hussar-lang/hussar/lexer"
	"github.com/hussar-lang/hussar/object"
	"github.com/hussar-lang/hussar/parser"
	"github.com/hussar-lang/hussar/repl"

	log "github.com/sirupsen/logrus"
	"github.com/ttacon/chalk"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	GitCommit     string
	VersionString string

	app     = kingpin.New("hussar", "The Hussar interpreter")
	verbose = app.Flag("verbose", "Enable verbose logging.").Short('v').Bool()

	// TODO: run interactive mode if no subcommand was given (see #1)
	interactive = app.Command("interactive", "Interactive REPL")

	run     = app.Command("run", "Run Hussar code")
	runFile = run.Flag("file", "Code to run").Required().Short('f').ExistingFile()
)

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
}

func main() {
	app.Version(fmt.Sprintf("%s (%s)", VersionString, GitCommit))
	args, err := app.Parse(os.Args[1:])

	if *verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	switch kingpin.MustParse(args, err) {
	case run.FullCommand():
		log.WithFields(log.Fields{
			"File":    *runFile,
			"Verbose": *verbose,
		}).Debug("Received run command")

		runFromFile()
	case interactive.FullCommand():
		startRepl()
	}
}

func runFromFile() {
	file, err := ioutil.ReadFile(*runFile)
	if err != nil {
		log.Fatal(err)
	}

	l := lexer.New(string(file))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(os.Stdout, p.Errors())
		os.Exit(21)
	}

	env := object.NewEnvironment()
	eval := evaluator.Eval(program, env)
	if eval.Inspect() != "NULL" {
		fmt.Println(eval.Inspect())
	}
}

func startRepl() {
	fmt.Printf("Starting Hussar interactive interpreter v%s\n", VersionString)
	repl.Start(os.Stdin, os.Stdout)
}

func printParserErrors(out io.Writer, errors []string) {
	errColor := chalk.Red.NewStyle().WithTextStyle(chalk.Bold).Style

	io.WriteString(out, errColor("PARSER ERROR!\n"))
	for _, msg := range errors {
		io.WriteString(out, errColor("  [!] ")+msg+"\n")
	}
}
