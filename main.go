package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/kscarlett/kmonkey/evaluator"
	"github.com/kscarlett/kmonkey/lexer"
	"github.com/kscarlett/kmonkey/object"
	"github.com/kscarlett/kmonkey/parser"
	"github.com/kscarlett/kmonkey/repl"

	log "github.com/sirupsen/logrus"
	"github.com/ttacon/chalk"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app     = kingpin.New("kmonkey", "The kmonkey interpreter")
	verbose = app.Flag("verbose", "Enable verbose logging.").Short('v').Bool()

	// TODO: run interactive mode if no subcommand was given
	interactive = app.Command("interactive", "Interactive REPL")

	run     = app.Command("run", "Run kmonkey code")
	runFile = run.Flag("file", "Code to run").Required().Short('f').ExistingFile()
)

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
}

func main() {
	app.Version("0.1.0")
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
	fmt.Printf("Starting kmonkey interactive interpreter v%s\n", "0.2.0") // Find a way to get this from app.version
	repl.Start(os.Stdin, os.Stdout)
}

func printParserErrors(out io.Writer, errors []string) {
	errColor := chalk.Red.NewStyle().WithTextStyle(chalk.Bold).Style

	io.WriteString(out, errColor("PARSER ERROR!\n"))
	for _, msg := range errors {
		io.WriteString(out, errColor("  [!] ")+msg+"\n")
	}
}
