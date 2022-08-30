package tokenize

import "log"

type Token struct {
	Kind TokenKind
	Pos  *Position

	Str      string
	NumFloat float64
	NumInt   int

	Next *Token
}

func NewToken(kind TokenKind, pos *Position, str string, nf float64, ni int) *Token {
	return &Token{
		Kind:     kind,
		Pos:      pos,
		Str:      str,
		NumFloat: nf,
		NumInt:   ni,
		Next:     nil,
	}
}

func NewOpSymbol(cur *Token, pos *Position, str string) *Token {
	var tok *Token
	switch str {
	case "(":
		tok = NewToken(Lrb, pos, str, 0, 0)
	case ")":
		tok = NewToken(Rrb, pos, str, 0, 0)
	case "[":
		tok = NewToken(Lsb, pos, str, 0, 0)
	case "]":
		tok = NewToken(Rsb, pos, str, 0, 0)
	case "{":
		tok = NewToken(Lcb, pos, str, 0, 0)
	case "}":
		tok = NewToken(Rcb, pos, str, 0, 0)
	case ".":
		tok = NewToken(Dot, pos, str, 0, 0)
	case ",":
		tok = NewToken(Comma, pos, str, 0, 0)
	case ":":
		tok = NewToken(Colon, pos, str, 0, 0)
	case ";":
		tok = NewToken(Semi, pos, str, 0, 0)

	case "+":
		tok = NewToken(Add, pos, str, 0, 0)
	case "-":
		tok = NewToken(Sub, pos, str, 0, 0)
	case "*":
		tok = NewToken(Mul, pos, str, 0, 0)
	case "/":
		tok = NewToken(Div, pos, str, 0, 0)
	case "%":
		tok = NewToken(Mod, pos, str, 0, 0)

	case "==":
		tok = NewToken(Eq, pos, str, 0, 0)
	case "!=":
		tok = NewToken(Ne, pos, str, 0, 0)
	case ">":
		tok = NewToken(Gt, pos, str, 0, 0)
	case "<":
		tok = NewToken(Lt, pos, str, 0, 0)
	case ">=":
		tok = NewToken(Ge, pos, str, 0, 0)
	case "<=":
		tok = NewToken(Le, pos, str, 0, 0)

	case "=":
		tok = NewToken(Assign, pos, str, 0, 0)
	case "+=":
		tok = NewToken(AddAssign, pos, str, 0, 0)
	case "-=":
		tok = NewToken(SubAssign, pos, str, 0, 0)
	case "*=":
		tok = NewToken(MulAssign, pos, str, 0, 0)
	case "/=":
		tok = NewToken(DivAssign, pos, str, 0, 0)
	case "%=":
		tok = NewToken(ModAssign, pos, str, 0, 0)

	case "&&":
		tok = NewToken(And, pos, str, 0, 0)
	case "||":
		tok = NewToken(Or, pos, str, 0, 0)
	case "!":
		tok = NewToken(Not, pos, str, 0, 0)
	default:
		log.Fatalf("unsupported operator/symbol: %s", str)
	}
	cur.Next = tok
	return tok
}

func NewIdent(cur *Token, pos *Position, str string) *Token {
	var tok *Token
	switch str {
	case "for":
		tok = NewToken(KWFor, pos, str, 0, 0)
	case "while":
		tok = NewToken(KWWhile, pos, str, 0, 0)
	case "if":
		tok = NewToken(KWIf, pos, str, 0, 0)
	case "else":
		tok = NewToken(KWElse, pos, str, 0, 0)
	case "return":
		tok = NewToken(KWReturn, pos, str, 0, 0)

	// data types
	case "dict":
		tok = NewToken(KWDict, pos, str, 0, 0)
	case "float":
		tok = NewToken(KWFloat, pos, str, 0, 0)
	case "int":
		tok = NewToken(KWInt, pos, str, 0, 0)
	case "string":
		tok = NewToken(KWString, pos, str, 0, 0)
	case "bool":
		tok = NewToken(KWBool, pos, str, 0, 0)
	case "void":
		tok = NewToken(KWVoid, pos, str, 0, 0)

	case "true":
		tok = NewToken(True, pos, str, 0, 0)
	case "false":
		tok = NewToken(False, pos, str, 0, 0)
	case "null":
		tok = NewToken(Null, pos, str, 0, 0)

	default:
		tok = NewToken(Ident, pos, str, 0, 0)
	}
	cur.Next = tok
	return tok
}

func NewEof(cur *Token, pos *Position) *Token {
	tok := NewToken(Eof, pos, "", 0, 0)
	cur.Next = tok
	return tok
}

func NewFloat(cur *Token, pos *Position, n float64) *Token {
	tok := NewToken(Float, pos, "", n, 0)
	cur.Next = tok
	return tok
}

func NewInt(cur *Token, pos *Position, n int) *Token {
	tok := NewToken(Int, pos, "", 0, n)
	cur.Next = tok
	return tok
}

func NewString(cur *Token, pos *Position, str string, isRaw bool) *Token {
	var tok *Token
	if isRaw {
		tok = NewToken(Raw, pos, str, 0, 0)
	} else {
		tok = NewToken(String, pos, str, 0, 0)
	}
	cur.Next = tok
	return tok
}

func NewComment(cur *Token, pos *Position, str string) *Token {
	tok := NewToken(Comment, pos, str, 0, 0)
	cur.Next = tok
	return tok
}

func NewWhite(cur *Token, pos *Position, str string) *Token {
	tok := NewToken(White, pos, str, 0, 0)
	cur.Next = tok
	return tok
}

func NewNL(cur *Token, pos *Position, str string) *Token {
	tok := NewToken(NewLine, pos, str, 0, 0)
	cur.Next = tok
	return tok
}
