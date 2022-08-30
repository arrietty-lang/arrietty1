package parse

import (
	"github.com/x0y14/arrietty/tokenize"
	"log"
)

type Node struct {
	Kind NodeKind

	Init *Node // for
	Cond *Node // for / while
	Loop *Node // for

	NumFloat float64
	NumInt   int
	Str      string

	Lhs      *Node
	Rhs      *Node
	Children []*Node
}

func NewNode(kind NodeKind, lhs, rhs *Node) *Node {
	return &Node{
		Kind: kind,
		Lhs:  lhs,
		Rhs:  rhs,
	}
}

func NewNodeImmediate(tok *tokenize.Token) *Node {
	switch tok.Kind {
	case tokenize.Float:
		return &Node{
			Kind:     Float,
			NumFloat: tok.NumFloat,
		}
	case tokenize.Int:
		return &Node{
			Kind:   Int,
			NumInt: tok.NumInt,
		}
	case tokenize.String:
		return &Node{
			Kind: String,
			Str:  tok.Str,
		}
	case tokenize.Raw:
		return &Node{
			Kind: Raw,
			Str:  tok.Str,
		}
	case tokenize.True:
		return &Node{
			Kind: True,
		}
	case tokenize.False:
		return &Node{
			Kind: False,
		}
	case tokenize.Null:
		return &Node{
			Kind: Null,
		}
		//case tokenize.Comment:
		//	return &Node{
		//		Kind: Comment,
		//		Str:  tok.Str,
		//	}
	}
	log.Fatalf("unexpected token: %s", tok.Str)
	return nil
}

func NewNodeWithExpr(kind NodeKind, init, cond, loop *Node, body []*Node) *Node {
	return &Node{
		Kind:     kind,
		Init:     init,
		Cond:     cond,
		Loop:     loop,
		Children: body,
	}
}

func NewNodeWithChildren(kind NodeKind, children []*Node) *Node {
	return &Node{
		Kind:     kind,
		Children: children,
	}
}

func NewNodeFunction(ident *Node, params *Node, body *Node) *Node {
	return &Node{
		Kind:     Function,
		Children: []*Node{ident, params, body},
	}
}

func NewNodeCall(ident *Node, args *Node) *Node {
	return &Node{
		Kind:     Call,
		Children: []*Node{ident, args},
	}
}

func NewNodeKV(key, value *Node) *Node {
	return &Node{
		Kind: KV,
		Lhs:  key,
		Rhs:  value,
	}
}

func NewNodeIdent(ident string) *Node {
	return &Node{
		Kind: Ident,
		Str:  ident,
	}
}

func NewNodeReturn(value *Node) *Node {
	return &Node{
		Kind:     Return,
		Children: []*Node{value},
	}
}

func NewNodeAccess(ident *Node, dest *Node) *Node {
	return &Node{
		Kind:     Access,
		Children: []*Node{ident, dest},
	}
}
