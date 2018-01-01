package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/kscarlett/kmonkey/ast"
	"github.com/ttacon/chalk"
)

type ObjectType string

const (
	ERROR_OBJ        = "ERROR"
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	FUNCTION_OBJ     = "FUNCTION"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

// === Errors ===
type Error struct {
	Severity string
	Message  string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string {
	errColor := chalk.Red.NewStyle().WithTextStyle(chalk.Bold).Style
	warnColor := chalk.Yellow.NewStyle().WithTextStyle(chalk.Bold).WithTextStyle(chalk.Dim).Style

	switch e.Severity {
	case "warn":
		return warnColor("[WARN]  ") + e.Message
	default:
		return errColor("[ERROR] ") + e.Message
	}
}

// === Integer ===
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// === Boolean ===
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// === Null ===
type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

// === Return ===
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

// === Function ===
type Function struct {
	//Name string		TODO: to be used to not have to bind to a variable explicitly
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
