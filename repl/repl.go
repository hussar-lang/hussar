package repl

import (
	"bufio"
	"io"

	"github.com/fatih/color"

	"github.com/kscarlett/kmonkey/lexer"
	"github.com/kscarlett/kmonkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	pColor := color.New(color.FgCyan, color.Bold)

	for {
		pColor.Printf(PROMPT)
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

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	err := color.New(color.FgRed, color.Bold)

	err.Println("ERROR!")
	for _, msg := range errors {
		err.Print("  [!] ")
		io.WriteString(out, msg+"\n")
	}
}
