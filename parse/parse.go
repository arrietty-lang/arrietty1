package parse

import (
	"log"

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
	if token.Kind == tokenize.Ident && id == token.Str {
		tok := token
		token = token.Next
		return tok
	}
	return nil
}

func expect(kind tokenize.TokenKind) *tokenize.Token {
	if token.Kind == kind {
		tok := token
		token = token.Next
		return tok
	}
	log.Fatalf("unexpected token: %v", token)
	return nil
}

func program() []*Node {
	var nodes []*Node
	for !isEof() {
		nodes = append(nodes, toplevel())
	}
	return nodes
}

func toplevel() *Node {
	// define function
	retType := types()

	id := NewNodeIdent(expect(tokenize.Ident).Str)
	expect(tokenize.Lrb)

	var params *Node
	if consume(tokenize.Rrb) == nil {
		// has parameters
		params = funcParams()
		expect(tokenize.Rrb)
	} else {
		// no parameters
		params = NewNodeWithChildren(Params, nil)
	}

	codeBlock := block()

	return NewNode(FuncDef, retType, NewNodeFunction(id, params, codeBlock))
}

func block() *Node {
	var nodes []*Node
	expect(tokenize.Lcb)
	for consume(tokenize.Rcb) == nil {
		nodes = append(nodes, stmt())
	}
	return NewNodeWithChildren(Block, nodes)
}

func stmt() *Node {
	var node *Node
	if consume(tokenize.KWReturn) != nil {
		node = NewNodeReturn(expr())
		expect(tokenize.Semi)
	} else if consume(tokenize.KWIf) != nil {
		expect(tokenize.Lrb)
		cond := expr()
		expect(tokenize.Rrb)
		ifBlock := block()
		if consume(tokenize.KWElse) != nil {
			elseBlock := block()
			node = NewNodeWithExpr(IfElse, nil, cond, nil, []*Node{ifBlock, elseBlock})
		} else {
			node = NewNodeWithExpr(If, nil, cond, nil, []*Node{ifBlock})
		}
	} else if consume(tokenize.KWWhile) != nil {
		expect(tokenize.Lrb)
		cond := expr()
		expect(tokenize.Rrb)
		whileBlock := block()
		node = NewNodeWithExpr(While, nil, cond, nil, []*Node{whileBlock})
	} else if consume(tokenize.KWFor) != nil {
		expect(tokenize.Lrb)
		var initExpr *Node
		var condExpr *Node
		var loopExpr *Node
		if consume(tokenize.Semi) == nil {
			initExpr = expr()
			expect(tokenize.Semi)
		}
		if consume(tokenize.Semi) == nil {
			condExpr = expr()
			expect(tokenize.Semi)
		}
		if consume(tokenize.Rrb) == nil {
			loopExpr = expr()
			expect(tokenize.Rrb)
		}
		forBlock := block()
		node = NewNodeWithExpr(For, initExpr, condExpr, loopExpr, []*Node{forBlock})
	} else {
		node = expr()
		expect(tokenize.Semi)
	}

	return node
}

func expr() *Node {
	return assign()
}

func assign() *Node {
	var node *Node
	if consumeIdent("var") != nil {
		// varDecl
		id := expect(tokenize.Ident)
		node = NewNode(VarDecl, NewNodeIdent(id.Str), types())

		if consume(tokenize.Assign) != nil {
			// and assign
			node = NewNode(Assign, node, andor())
		}
		// only decl
		return node
	}

	// shortVarDecl or assign or andor
	node = andor()

	if consume(tokenize.Assign) != nil {
		node = NewNode(Assign, node, andor())
	} else if consume(tokenize.ColonAssign) != nil {
		node = NewNode(ShortVarDecl, node, andor())
	}
	return node
}

func andor() *Node {
	node := equality()
	for {
		if consume(tokenize.And) != nil {
			node = NewNode(And, node, equality())
		} else if consume(tokenize.Or) != nil {
			node = NewNode(Or, node, equality())
		} else {
			return node
		}
	}
}

func equality() *Node {
	node := relational()
	for {
		if consume(tokenize.Eq) != nil {
			node = NewNode(Eq, node, relational())
		} else if consume(tokenize.Ne) != nil {
			node = NewNode(Ne, node, relational())
		} else {
			return node
		}
	}
}

func relational() *Node {
	node := add()
	for {
		if consume(tokenize.Lt) != nil {
			node = NewNode(Lt, node, add())
		} else if consume(tokenize.Le) != nil {
			node = NewNode(Le, node, add())
		} else if consume(tokenize.Gt) != nil {
			node = NewNode(Gt, node, add())
		} else if consume(tokenize.Ge) != nil {
			node = NewNode(Ge, node, add())
		} else {
			return node
		}
	}
}

func add() *Node {
	node := mul()
	for {
		if consume(tokenize.Add) != nil {
			node = NewNode(Add, node, mul())
		} else if consume(tokenize.Sub) != nil {
			node = NewNode(Sub, node, mul())
		} else {
			return node
		}
	}

}

func mul() *Node {
	node := unary()
	for {
		if consume(tokenize.Mul) != nil {
			node = NewNode(Mul, node, mul())
		} else if consume(tokenize.Div) != nil {
			node = NewNode(Div, node, mul())
		} else if consume(tokenize.Mod) != nil {
			node = NewNode(Mod, node, mul())
		} else {
			return node
		}
	}
}

func unary() *Node {
	var node *Node
	if consume(tokenize.Add) != nil {
		node = primary()
	} else if consume(tokenize.Sub) != nil {
		node = NewNode(Sub, NewNodeImmediate(tokenize.NewInt(nil, nil, 0)), primary())
	} else if consume(tokenize.Not) != nil {
		node = NewNode(Not, primary(), nil)
	} else {
		node = primary()
	}
	return node
}

func primary() *Node {
	return access()
}

func access() *Node {
	node := literal()
	for {
		if consume(tokenize.Lsb) != nil {
			node = NewNodeAccess(node, expr())
			expect(tokenize.Rsb)
		} else {
			return node
		}
	}
}

func literal() *Node {
	var node *Node
	if consume(tokenize.Lrb) != nil {
		node = expr()
		expect(tokenize.Rrb)
		return node
	}

	if n := consume(tokenize.Ident); n != nil {
		if consume(tokenize.Lrb) != nil {
			if consume(tokenize.Rrb) != nil {
				node = NewNodeCall(NewNodeIdent(n.Str), NewNodeWithChildren(Args, nil))
				return node
			} else {
				node = NewNodeCall(NewNodeIdent(n.Str), callArgs())
				expect(tokenize.Rrb)
				return node
			}
		} else {
			node = NewNodeIdent(n.Str)
			return node
		}
	}

	if consume(tokenize.Lsb) != nil {
		node = array()
		//expect(tokenize.Rsb)
		return node
	} else if consume(tokenize.Lcb) != nil {
		node = dict()
		//expect(tokenize.Rcb)
		return node
	}

	return immediate()
}

func immediate() *Node {
	if v := consume(tokenize.Float); v != nil {
		return NewNodeImmediate(v)
	} else if v := consume(tokenize.Int); v != nil {
		return NewNodeImmediate(v)
	} else if v := consume(tokenize.String); v != nil {
		return NewNodeImmediate(v)
	} else if v := consume(tokenize.Raw); v != nil {
		return NewNodeImmediate(v)
	} else if v := consume(tokenize.True); v != nil {
		return NewNodeImmediate(v)
	} else if v := consume(tokenize.False); v != nil {
		return NewNodeImmediate(v)
	} else if v := consume(tokenize.Null); v != nil {
		return NewNodeImmediate(v)
	}
	log.Fatalf("unexpected token: %v", token)
	return nil
}

func types() *Node {
	if consumeIdent("float") != nil {
		return NewNodeWithChildren(Float, nil)
	}
	if consumeIdent("int") != nil {
		return NewNodeWithChildren(Int, nil)
	}
	if consumeIdent("string") != nil {
		return NewNodeWithChildren(String, nil)
	}
	if consumeIdent("bool") != nil {
		return NewNodeWithChildren(Bool, nil)
	}
	if consumeIdent("void") != nil {
		return NewNodeWithChildren(Void, nil)
	}
	if consume(tokenize.Lsb) != nil {
		// list
		// lengthはevalしてintが出ることを期待している
		// [?]type
		// ^^^
		var length *Node
		if consume(tokenize.Rsb) == nil {
			// length
			length = expr()
			expect(tokenize.Rsb)
		} else {
			// dynamic
			length = &Node{Kind: Int, NumInt: -1}
		}
		// [?]type
		//    ^^^^
		itemType := types()
		return NewNodeWithChildren(List, []*Node{length, itemType})
	}
	if consumeIdent("dict") != nil {
		// dict
		// dict[k-type]v-type
		//     ^^^^^^^
		expect(tokenize.Lsb)
		keyType := types()
		expect(tokenize.Rsb)
		valueType := types()
		return NewNodeWithChildren(Dict, []*Node{keyType, valueType})
	}
	// named ident
	return NewNodeIdent(expect(tokenize.Ident).Str)
}

func callArgs() *Node {
	nodes := []*Node{expr()}
	for consume(tokenize.Comma) != nil {
		nodes = append(nodes, expr())
	}
	return NewNodeWithChildren(Args, nodes)
}

func funcParams() *Node {
	nodes := []*Node{NewNodeIdent(consume(tokenize.Ident).Str)}
	for consume(tokenize.Comma) != nil {
		nodes = append(nodes, NewNodeIdent(consume(tokenize.Ident).Str))
	}
	return NewNodeWithChildren(Params, nodes)
}

func array() *Node {
	var nodes []*Node
	for consume(tokenize.Rsb) == nil {
		nodes = append(nodes, primary())
		if consume(tokenize.Comma) == nil {
			expect(tokenize.Rsb)
			break
		}
	}
	return NewNodeWithChildren(List, nodes)
}

func dict() *Node {
	var nodes []*Node
	for consume(tokenize.Rcb) == nil {
		nodes = append(nodes, kv())
		if consume(tokenize.Comma) == nil {
			expect(tokenize.Rcb)
			break
		}
	}
	return NewNodeWithChildren(Dict, nodes)
}

func kv() *Node {
	key := NewNodeImmediate(consume(tokenize.String))
	expect(tokenize.Colon)
	value := primary()
	return NewNodeKV(key, value)
}

func Parse(tok *tokenize.Token) []*Node {
	token = tok
	return program()
}
