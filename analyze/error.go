package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type SemanticErrKind int

const (
	_ SemanticErrKind = iota
	UnexpectedNodeErr
	UnimplementedErr
	AlreadyDefinedErr
	UndefinedErr
	UnsupportedTypeErr
	TypeErr
)

type SemanticErr struct {
	Kind SemanticErrKind
	Node *parse.Node
	ToplevelKind
	Ident  string
	Ident2 string
}

func NewUnexpectNodeErr(node *parse.Node) error {
	return &SemanticErr{Kind: UnexpectedNodeErr, Node: node}
}

func NewUnimplementedErr(t ToplevelKind) error {
	return &SemanticErr{Kind: UnimplementedErr, ToplevelKind: t}
}

func NewAlreadyDefinedErr(ident1 string, ident2 string) error {
	return &SemanticErr{Kind: AlreadyDefinedErr, Ident: ident1, Ident2: ident2}
}

func NewUndefinedErr(ident string) error {
	return &SemanticErr{Kind: UndefinedErr, Ident: ident}
}

func NewUnsupportedTypeErr(n *parse.Node) error {
	return &SemanticErr{Kind: UnsupportedTypeErr, Node: n}
}

func (e *SemanticErr) Error() string {
	switch e.Kind {
	case UnexpectedNodeErr:
		return fmt.Sprintf("%s : unexpect node", e.Node.String())
	case UnimplementedErr:
		return fmt.Sprintf("unimplemented: %s", e.ToplevelKind.String())
	case AlreadyDefinedErr:
		return fmt.Sprintf("already defined in %s: %s", e.Ident, e.Ident2)
	case UndefinedErr:
		return fmt.Sprintf("undefined: %s", e.Ident)
	case UnsupportedTypeErr:
		return fmt.Sprintf("the valutype %s is unsupported", e.Node.String())
	case TypeErr:
		return fmt.Sprintf(e.Ident)
	}

	return ""
}
