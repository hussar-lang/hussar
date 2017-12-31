package repl

import (
	"bufio"
	"io"

	"github.com/ttacon/chalk"

	"github.com/kscarlett/kmonkey/evaluator"
	"github.com/kscarlett/kmonkey/lexer"
	"github.com/kscarlett/kmonkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	promptColor := chalk.Cyan.NewStyle().WithTextStyle(chalk.Bold).Style

	for {
		io.WriteString(out, promptColor(PROMPT))
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	errColor := chalk.Red.NewStyle().WithTextStyle(chalk.Bold).Style

	io.WriteString(out, errColor("PARSER ERROR!\n"))
	for _, msg := range errors {
		io.WriteString(out, errColor("  [!] ")+msg+"\n")
	}
}
