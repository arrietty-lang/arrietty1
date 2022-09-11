package parse

import (
	"fmt"
	"github.com/x0y14/arrietty/tokenize"
)

var token *tokenize.Token

func isEof() bool {
	return token.Kind == tokenize.Eof
}

func consume(kind tokenize.TokenKind) *tokenize.Token {
	if token.Kind == kind {
		tok := token
		token = token.Next
		return tok
	}
	return nil
}

func consumeIdent(id string) *tokenize.Token {
	if token.Kind == tokenize.Ident && id == token.S {
		tok := token
		token = token.Next
		return tok
	}
	return nil
}

func expect(kind tokenize.TokenKind) (*tokenize.Token, error) {
	if token.Kind == kind {
		tok := token
		token = token.Next
		return tok, nil
	}
	return nil, NewUnexpectedTokenErr(kind.String(), token)
}

func program() ([]*Node, error) {
	var nodes []*Node
	for !isEof() {
		n, err := toplevel()
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, n)
	}
	return nodes, nil
}

func toplevel() (*Node, error) {
	if c_ := consume(tokenize.Comment); c_ != nil {
		return NewNodeComment(c_.Pos, c_.S), nil
	}

	retType, err := types()
	if err != nil {
		return nil, err
	}

	idPos := token.Pos
	id, err := expect(tokenize.Ident)
	if err != nil {
		return nil, err
	}
	ident := NewNodeIdent(idPos, id.S)

	// (
	_, err = expect(tokenize.Lrb)
	if err != nil {
		return nil, err
	}

	var params *Node = nil

	// )
	if consume(tokenize.Rrb) == nil {
		// params
		params, err = funcParams()
		// )
		_, err = expect(tokenize.Rrb)
		if err != nil {
			return nil, err
		}
	}

	// { stmt* }
	codeBlock, err := block()
	if err != nil {
		return nil, err
	}

	return NewNodeFunctionDefine(retType.Pos, retType, ident, params, codeBlock), nil
}

func block() (*Node, error) {
	var nodes []*Node
	lcb, err := expect(tokenize.Lcb)
	if err != nil {
		return nil, err
	}

	for consume(tokenize.Rcb) == nil {
		n, err := stmt()
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, n)
	}

	return NewNodeWithChildren(lcb.Pos, Block, nodes), nil
}

func stmt() (*Node, error) {
	var node *Node

	if c_ := consume(tokenize.Comment); c_ != nil {
		return NewNodeComment(c_.Pos, c_.S), nil
	}

	if return_ := consumeIdent("return"); return_ != nil {
		if consume(tokenize.Semi) != nil {
			node = NewNodeReturn(return_.Pos, nil)
			return node, nil
		}

		n, err := expr()
		if err != nil {
			return nil, err
		}
		node = NewNodeReturn(return_.Pos, n)

		_, err = expect(tokenize.Semi)
		if err != nil {
			return nil, err
		}
	} else if if_ := consumeIdent("if"); if_ != nil {
		_, err := expect(tokenize.Lrb)
		if err != nil {
			return nil, err
		}

		cond, err := expr()
		if err != nil {
			return nil, err
		}

		_, err = expect(tokenize.Rrb)
		if err != nil {
			return nil, err
		}

		ifBlock, err := block()
		if err != nil {
			return nil, err
		}

		if else_ := consumeIdent("else"); else_ != nil {
			elseBlock, err := block()
			if err != nil {
				return nil, err
			}
			node = NewNodeWithExpr(if_.Pos, IfElse, nil, cond, nil, []*Node{ifBlock, elseBlock})
		} else {
			node = NewNodeWithExpr(if_.Pos, If, nil, cond, nil, []*Node{ifBlock})
		}
	} else if while_ := consumeIdent("while"); while_ != nil {
		_, err := expect(tokenize.Lrb)
		if err != nil {
			return nil, err
		}

		cond, err := expr()
		if err != nil {
			return nil, err
		}

		_, err = expect(tokenize.Rrb)
		if err != nil {
			return nil, err
		}

		whileBlock, err := block()
		if err != nil {
			return nil, err
		}

		node = NewNodeWithExpr(while_.Pos, While, nil, cond, nil, []*Node{whileBlock})
	} else if for_ := consumeIdent("for"); for_ != nil {
		_, err := expect(tokenize.Lrb)
		if err != nil {
			return nil, err
		}
		var init *Node = nil
		var cond *Node = nil
		var loop *Node = nil

		// init
		if consume(tokenize.Semi) == nil {
			_init, err := expr()
			if err != nil {
				return nil, err
			}
			init = _init
			_, err = expect(tokenize.Semi)
			if err != nil {
				return nil, err
			}
		}
		// cond
		if consume(tokenize.Semi) == nil {
			_cond, err := expr()
			if err != nil {
				return nil, err
			}
			cond = _cond
			_, err = expect(tokenize.Semi)
			if err != nil {
				return nil, err
			}
		}
		// loop
		if consume(tokenize.Rrb) == nil {
			_loop, err := expr()
			if err != nil {
				return nil, err
			}
			loop = _loop
			_, err = expect(tokenize.Rrb)
			if err != nil {
				return nil, err
			}
		}
		forBlock, err := block()
		if err != nil {
			return nil, err
		}
		node = NewNodeWithExpr(for_.Pos, For, init, cond, loop, []*Node{forBlock})
	} else {
		n, err := expr()
		if err != nil {
			return nil, err
		}
		node = n
		_, err = expect(tokenize.Semi)
		if err != nil {
			return nil, err
		}
	}

	return node, nil
}

func expr() (*Node, error) {
	return assign()
}

func assign() (*Node, error) {
	var node *Node

	if var_ := consumeIdent("var"); var_ != nil {
		id, err := expect(tokenize.Ident)
		if err != nil {
			return nil, err
		}

		typ, err := types()
		if err != nil {
			return nil, err
		}

		node = NewNode(var_.Pos, VarDecl, NewNodeIdent(id.Pos, id.S), typ)

		if assign_ := consume(tokenize.Assign); assign_ != nil {
			// decl and assign
			n, err := andor()
			if err != nil {
				return nil, err
			}
			node = NewNode(assign_.Pos, Assign, node, n)
		}
		// only decl
		return node, nil
	}

	node, err := andor()
	if err != nil {
		return nil, err
	}

	if assign_ := consume(tokenize.Assign); assign_ != nil {
		// assign
		rhs, err := andor()
		if err != nil {
			return nil, err
		}
		node = NewNode(assign_.Pos, Assign, node, rhs)
	} else if colonAssign_ := consume(tokenize.ColonAssign); colonAssign_ != nil {
		// :=
		rhs, err := andor()
		if err != nil {
			return nil, err
		}
		node = NewNode(colonAssign_.Pos, ShortVarDecl, node, rhs)
	}
	return node, nil
}

func andor() (*Node, error) {
	node, err := equality()
	if err != nil {
		return nil, err
	}

	for {
		if and_ := consume(tokenize.And); and_ != nil {
			n, err := equality()
			if err != nil {
				return nil, err
			}
			node = NewNode(and_.Pos, And, node, n)
		} else if or_ := consume(tokenize.Or); or_ != nil {
			n, err := equality()
			if err != nil {
				return nil, err
			}
			node = NewNode(or_.Pos, Or, node, n)
		} else {
			return node, nil
		}
	}
}

func equality() (*Node, error) {
	node, err := relational()
	if err != nil {
		return nil, err
	}

	for {
		if eq_ := consume(tokenize.Eq); eq_ != nil {
			n, err := relational()
			if err != nil {
				return nil, err
			}
			node = NewNode(eq_.Pos, Eq, node, n)
		} else if ne_ := consume(tokenize.Ne); ne_ != nil {
			n, err := relational()
			if err != nil {
				return nil, err
			}
			node = NewNode(ne_.Pos, Ne, node, n)
		} else {
			return node, nil
		}
	}
}

func relational() (*Node, error) {
	node, err := add()
	if err != nil {
		return nil, err
	}

	for {
		if lt_ := consume(tokenize.Lt); lt_ != nil {
			n, err := add()
			if err != nil {
				return nil, err
			}
			node = NewNode(lt_.Pos, Lt, node, n)
		} else if le_ := consume(tokenize.Le); le_ != nil {
			n, err := add()
			if err != nil {
				return nil, err
			}
			node = NewNode(le_.Pos, Le, node, n)
		} else if gt_ := consume(tokenize.Gt); gt_ != nil {
			n, err := add()
			if err != nil {
				return nil, err
			}
			node = NewNode(gt_.Pos, Gt, node, n)
		} else if le_ := consume(tokenize.Ge); le_ != nil {
			n, err := add()
			if err != nil {
				return nil, err
			}
			node = NewNode(le_.Pos, Ge, node, n)
		} else {
			return node, nil
		}
	}
}

func add() (*Node, error) {
	node, err := mul()
	if err != nil {
		return nil, err
	}

	for {
		if add_ := consume(tokenize.Add); add_ != nil {
			n, err := mul()
			if err != nil {
				return nil, err
			}
			node = NewNode(add_.Pos, Add, node, n)
		} else if sub_ := consume(tokenize.Sub); sub_ != nil {
			n, err := mul()
			if err != nil {
				return nil, err
			}
			node = NewNode(sub_.Pos, Sub, node, n)
		} else {
			return node, nil
		}
	}
}

func mul() (*Node, error) {
	node, err := unary()
	if err != nil {
		return nil, err
	}

	for {
		if mul_ := consume(tokenize.Mul); mul_ != nil {
			n, err := unary()
			if err != nil {
				return nil, err
			}
			node = NewNode(mul_.Pos, Mul, node, n)
		} else if div_ := consume(tokenize.Div); div_ != nil {
			n, err := unary()
			if err != nil {
				return nil, err
			}
			node = NewNode(div_.Pos, Div, node, n)
		} else if mod_ := consume(tokenize.Mod); mod_ != nil {
			n, err := unary()
			if err != nil {
				return nil, err
			}
			node = NewNode(mod_.Pos, Mod, node, n)
		} else {
			return node, nil
		}
	}
}

func unary() (*Node, error) {
	var node *Node

	if add_ := consume(tokenize.Add); add_ != nil {
		n, err := primary()
		if err != nil {
			return nil, err
		}
		node = NewNode(add_.Pos, Plus, n, nil)
	} else if sub_ := consume(tokenize.Sub); sub_ != nil {
		n, err := primary()
		if err != nil {
			return nil, err
		}
		node = NewNode(sub_.Pos, Minus, n, nil)
	} else if not_ := consume(tokenize.Not); not_ != nil {
		n, err := primary()
		if err != nil {
			return nil, err
		}
		node = NewNode(not_.Pos, Not, n, nil)
	} else {
		n, err := primary()
		if err != nil {
			return nil, err
		}
		node = n
	}

	return node, nil
}

func primary() (*Node, error) {
	return access()
}

func access() (*Node, error) {
	node, err := literal()
	if err != nil {
		return nil, err
	}

	for {
		if lsb_ := consume(tokenize.Lsb); lsb_ != nil {
			index, err := expr()
			if err != nil {
				return nil, err
			}
			node = NewNodeAccess(lsb_.Pos, node, index)
			_, err = expect(tokenize.Rsb)
			if err != nil {
				return nil, err
			}
		} else {
			return node, nil
		}
	}
}

func literal() (*Node, error) {
	var node *Node

	if lrb_ := consume(tokenize.Lrb); lrb_ != nil {
		n, err := expr()
		if err != nil {
			return nil, err
		}
		node = NewNode(lrb_.Pos, Parenthesis, n, nil)
		_, err = expect(tokenize.Rrb)
		if err != nil {
			return nil, err
		}
		return node, nil
	}

	if id_ := consume(tokenize.Ident); id_ != nil {
		if consume(tokenize.Lrb) != nil {
			if rrb_ := consume(tokenize.Rrb); rrb_ != nil {
				// no call args
				node = NewNodeCall(id_.Pos, NewNodeIdent(id_.Pos, id_.S), nil)
				return node, nil
			} else {
				// with call args
				n, err := callArgs()
				if err != nil {
					return nil, err
				}
				node = NewNodeCall(id_.Pos, NewNodeIdent(id_.Pos, id_.S), n)
				_, err = expect(tokenize.Rrb)
				if err != nil {
					return nil, err
				}
				return node, nil
			}
		} else {
			node = NewNodeIdent(id_.Pos, id_.S)
			return node, nil
		}
	}

	if lsb_ := consume(tokenize.Lsb); lsb_ != nil {
		// 開始地点の情報はこちら側が所持しているので、付け替えてあげる
		l, err := list()
		if err != nil {
			return nil, err
		}
		l.Pos = lsb_.Pos
		return l, nil
	}
	if lcb_ := consume(tokenize.Lcb); lcb_ != nil {
		// 開始地点の情報はこちら側が所持しているので、付け替えてあげる
		d, err := dict()
		if err != nil {
			return nil, err
		}
		d.Pos = lcb_.Pos
		return d, nil
	}
	return immediate()
}

func immediate() (*Node, error) {
	pos := token.Pos
	if v := consume(tokenize.Float); v != nil {
		return NewNodeImmediate(pos, v), nil
	} else if v := consume(tokenize.Int); v != nil {
		return NewNodeImmediate(pos, v), nil
	} else if v := consume(tokenize.String); v != nil {
		return NewNodeImmediate(pos, v), nil
	} else if v := consume(tokenize.RawString); v != nil {
		return NewNodeImmediate(pos, v), nil
	} else if v := consume(tokenize.True); v != nil {
		return NewNodeImmediate(pos, v), nil
	} else if v := consume(tokenize.False); v != nil {
		return NewNodeImmediate(pos, v), nil
	} else if v := consume(tokenize.Null); v != nil {
		return NewNodeImmediate(pos, v), nil
	}
	return nil, NewUnexpectedTokenErr("immediate", token)
}

func types() (*Node, error) {
	if f_ := consumeIdent("float"); f_ != nil {
		return NewNodeWithChildren(f_.Pos, Float, nil), nil
	}
	if i_ := consumeIdent("int"); i_ != nil {
		return NewNodeWithChildren(i_.Pos, Int, nil), nil
	}
	if s_ := consumeIdent("string"); s_ != nil {
		return NewNodeWithChildren(s_.Pos, String, nil), nil
	}
	if b_ := consumeIdent("bool"); b_ != nil {
		return NewNodeWithChildren(b_.Pos, Bool, nil), nil
	}
	if v_ := consumeIdent("void"); v_ != nil {
		return NewNodeWithChildren(v_.Pos, Void, nil), nil
	}
	if a_ := consumeIdent("any"); a_ != nil {
		return NewNodeWithChildren(a_.Pos, Any, nil), nil
	}

	// list
	if lsb_ := consume(tokenize.Lsb); lsb_ != nil {
		var length *Node = nil
		if consume(tokenize.Rsb) == nil {
			l_, err := expr()
			if err != nil {
				return nil, err
			}
			if l_.Kind != Int {
				return nil, fmt.Errorf("list size must be positive int")
			}
			length = l_
			_, err = expect(tokenize.Rsb)
			if err != nil {
				return nil, err
			}
		}
		itemType, err := types()
		if err != nil {
			return nil, err
		}
		return NewNodeWithChildren(lsb_.Pos, List, []*Node{length, itemType}), nil
	}

	// dict
	if d_ := consumeIdent("dict"); d_ != nil {
		// dict[k]v
		//     ^
		_, err := expect(tokenize.Lsb)
		if err != nil {
			return nil, err
		}
		// dict[k]v
		//      ^
		keyType, err := types()
		if err != nil {
			return nil, err
		}
		if keyType.Kind != String {
			return nil, fmt.Errorf("dict key is not support other than string")
		}
		// dict[k]v
		//       ^
		_, err = expect(tokenize.Rsb)
		if err != nil {
			return nil, err
		}
		// dict[k]v
		//        ^
		valueType, err := types()
		if err != nil {
			return nil, err
		}
		return NewNodeWithChildren(d_.Pos, Dict, []*Node{keyType, valueType}), nil
	}

	// named type
	id_, err := expect(tokenize.Ident)
	if err != nil {
		return nil, err
	}
	return NewNodeIdent(id_.Pos, id_.S), nil
}

func callArgs() (*Node, error) {

	firstId, err := expr()
	if err != nil {
		return nil, err
	}
	nodes := []*Node{firstId}

	for consume(tokenize.Comma) != nil {
		n, err := expr()
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, n)
	}

	return NewNodeWithChildren(firstId.Pos, Args, nodes), nil
}

func funcParams() (*Node, error) {
	firstId := consume(tokenize.Ident)
	firstType, err := types()
	if err != nil {
		return nil, err
	}
	nodes := []*Node{NewNode(firstId.Pos, Param, NewNodeIdent(firstId.Pos, firstId.S), firstType)}

	for consume(tokenize.Comma) != nil {
		id_ := consume(tokenize.Ident)
		type_, err := types()
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, NewNode(id_.Pos, Param, NewNodeIdent(id_.Pos, id_.S), type_))
	}
	return NewNodeWithChildren(firstId.Pos, Params, nodes), nil
}

func list() (*Node, error) {
	var nodes []*Node

	for consume(tokenize.Rsb) == nil {
		n, err := unary()
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, n)

		if consume(tokenize.Comma) == nil {
			_, err = expect(tokenize.Rsb)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	return NewNodeWithChildren(nil, List, nodes), nil
}

func dict() (*Node, error) {
	var nodes []*Node

	for consume(tokenize.Rcb) == nil {
		n, err := kv()
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, n)

		if consume(tokenize.Comma) == nil {
			_, err = expect(tokenize.Rcb)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	return NewNodeWithChildren(nil, Dict, nodes), nil
}

func kv() (*Node, error) {
	s := consume(tokenize.String)
	key := NewNodeImmediate(s.Pos, s)
	_, err := expect(tokenize.Colon)
	if err != nil {
		return nil, err
	}
	value, err := unary()
	if err != nil {
		return nil, err
	}
	return NewNodeKV(s.Pos, key, value), nil
}

func Parse(tok *tokenize.Token) ([]*Node, error) {
	token = tok
	return program()
}
