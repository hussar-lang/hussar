package ast

type Node interface {
	TokenLiteral() string // Only needed for testing and debugging
	String() string       // Used for comparison and printing in debugging
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}
