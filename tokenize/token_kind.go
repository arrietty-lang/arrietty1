package tokenize

type TokenKind int

const (
	_ TokenKind = iota
	Eof

	Comment
	White
	NewLine

	Ident // 5
	Float
	Int
	String
	RawString
	True // 10
	False
	Null

	symbolsBegin
	Lrb   // (
	Rrb   // ) 15
	Lsb   // [
	Rsb   // ]
	Lcb   // {
	Rcb   // }
	Dot   // . 20
	Comma // ,
	Colon // :
	Semi  // ;
	symbolsEnd

	operatorsBegin // 25
	Add            // +
	Sub            // -
	Mul            // *
	Div            // /
	Mod            // % 30

	Eq // ==
	Ne // !=
	Gt // >
	Lt // <
	Ge // >= 35
	Le // <=

	Assign      // =
	AddAssign   // +=
	SubAssign   // -=
	MulAssign   // *= 40
	DivAssign   // /=
	ModAssign   // %=
	ColonAssign // :=

	And // &&
	Or  // ||
	Not // ! 45
	operatorsEnd
)

var tokenKinds = [...]string{
	Eof:         "Eof",
	Comment:     "Comment",
	White:       "White",
	NewLine:     "NewLine",
	Ident:       "Ident",
	Float:       "Float",
	Int:         "Int",
	String:      "String",
	RawString:   "RawString",
	True:        "True",
	False:       "False",
	Null:        "Null",
	Lrb:         "Lrb",
	Rrb:         "Rrb",
	Lsb:         "Lsb",
	Rsb:         "Rsb",
	Lcb:         "Lcb",
	Rcb:         "Rcb",
	Dot:         "Dot",
	Comma:       "Comma",
	Colon:       "Colon",
	Semi:        "Semi",
	Add:         "Add",
	Sub:         "Sub",
	Mul:         "Mul",
	Div:         "Div",
	Mod:         "Mod",
	Eq:          "Eq",
	Ne:          "Ne",
	Gt:          "Gt",
	Lt:          "Lt",
	Ge:          "Ge",
	Le:          "Le",
	Assign:      "Assign",
	AddAssign:   "AddAssign",
	SubAssign:   "SubAssign",
	MulAssign:   "MulAssign",
	DivAssign:   "DivAssign",
	ModAssign:   "ModAssign",
	ColonAssign: "ColonAssign",
	And:         "And",
	Or:          "Or",
	Not:         "Not",
}

func (t TokenKind) String() string {
	return tokenKinds[t]
}
