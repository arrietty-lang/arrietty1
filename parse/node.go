package parse

import (
	"fmt"
	"github.com/x0y14/arrietty/tokenize"
	"log"
)

type Node struct {
	Kind NodeKind
	Pos  *tokenize.Position // 関数定義ノードなどではnilにしても良い気がするが、記号をノードして置き換える場合は位置情報が失われるので必須

	Init *Node // for
	Cond *Node // for / while
	Loop *Node // for

	// 即値s
	F float64
	I int
	S string

	Lhs      *Node
	Rhs      *Node
	Children []*Node
}

func (n *Node) String() string {
	v := ""
	switch n.Kind {
	case String:
		v = n.S
	case Int:
		v = fmt.Sprintf("%d", n.I)
	case Float:
		v = fmt.Sprintf("%f", n.F)
	}
	if v != "" {
		v = " : " + v
	}

	return fmt.Sprintf("[%d:%d(%d)] %s%s", n.Pos.LineNo, n.Pos.Lat, n.Pos.Wat, n.Kind.String(), v)
}

func NewNode(pos *tokenize.Position, kind NodeKind, lhs, rhs *Node) *Node {
	return &Node{
		Pos:  pos,
		Kind: kind,
		Lhs:  lhs,
		Rhs:  rhs,
	}
}

func NewNodeImmediate(pos *tokenize.Position, tok *tokenize.Token) *Node {
	switch tok.Kind {
	case tokenize.Float:
		return &Node{
			Pos:  pos,
			Kind: Float,
			F:    tok.F,
		}
	case tokenize.Int:
		return &Node{
			Pos:  pos,
			Kind: Int,
			I:    tok.I,
		}
	case tokenize.String:
		return &Node{
			Pos:  pos,
			Kind: String,
			S:    tok.S,
		}
	case tokenize.RawString:
		return &Node{
			Pos:  pos,
			Kind: RawString,
			S:    tok.S,
		}
	case tokenize.True:
		return &Node{
			Pos:  pos,
			Kind: True,
		}
	case tokenize.False:
		return &Node{
			Pos:  pos,
			Kind: False,
		}
	case tokenize.Null:
		return &Node{
			Pos:  pos,
			Kind: Null,
		}
	}
	log.Fatalf("unexpected token: %s", tok.S)
	return nil
}

func NewNodeWithExpr(pos *tokenize.Position, kind NodeKind, init, cond, loop *Node, body []*Node) *Node {
	return &Node{
		Pos:      pos,
		Kind:     kind,
		Init:     init,
		Cond:     cond,
		Loop:     loop,
		Children: body,
	}
}

func NewNodeWithChildren(pos *tokenize.Position, kind NodeKind, children []*Node) *Node {
	return &Node{
		Pos:      pos,
		Kind:     kind,
		Children: children,
	}
}

func NewNodeFunctionDefine(pos *tokenize.Position, retType *Node, ident *Node, params *Node, body *Node) *Node {
	return &Node{
		Pos:      pos,
		Kind:     FuncDef,
		Children: []*Node{retType, ident, params, body},
	}
}

func NewNodeCall(pos *tokenize.Position, ident *Node, args *Node) *Node {
	return &Node{
		Pos:      pos,
		Kind:     Call,
		Children: []*Node{ident, args},
	}
}

func NewNodeKV(pos *tokenize.Position, key, value *Node) *Node {
	return &Node{
		Pos:  pos,
		Kind: KV,
		Lhs:  key,
		Rhs:  value,
	}
}

func NewNodeIdent(pos *tokenize.Position, ident string) *Node {
	return &Node{
		Pos:  pos,
		Kind: Ident,
		S:    ident,
	}
}

func NewNodeReturn(pos *tokenize.Position, value *Node) *Node {
	return &Node{
		Pos:      pos,
		Kind:     Return,
		Children: []*Node{value},
	}
}

func NewNodeAccess(pos *tokenize.Position, ident *Node, dest *Node) *Node {
	return &Node{
		Pos:      pos,
		Kind:     Access,
		Children: []*Node{ident, dest},
	}
}

func NewNodeComment(pos *tokenize.Position, comment string) *Node {
	return &Node{
		Kind: Comment,
		Pos:  pos,
		S:    comment,
	}
}
