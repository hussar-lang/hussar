package object

import "github.com/ttacon/chalk"

type Error struct {
	Severity string
	Message  string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string {
	// TODO: Formatting of text shouldn't be happening in any package other than console I/O
	errColor := chalk.Red.NewStyle().WithTextStyle(chalk.Bold).Style
	warnColor := chalk.Yellow.NewStyle().WithTextStyle(chalk.Bold).Style

	switch e.Severity {
	case "warn":
		return warnColor("[WARN]  ") + e.Message
	default:
		return errColor("[ERROR] ") + e.Message
	}
}
