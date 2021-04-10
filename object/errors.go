package object

import "github.com/ttacon/chalk"

type Error struct {
	Severity string
	Message  string
}

// Type returns the type of object represented
func (e *Error) Type() ObjectType { return ERROR_OBJ }

// Inspect returns a string representation of the value
func (e *Error) Inspect() string {
	// TODO: Formatting of text shouldn't be happening in any package other than console I/O
	//  ideally, we'd be returning errors as a whole object with all the necessary metadata
	errColor := chalk.Red.NewStyle().WithTextStyle(chalk.Bold).Style
	warnColor := chalk.Yellow.NewStyle().WithTextStyle(chalk.Bold).Style

	switch e.Severity {
	case "warn":
		return warnColor("[WARN]  ") + e.Message
	default:
		return errColor("[ERROR] ") + e.Message
	}
}
