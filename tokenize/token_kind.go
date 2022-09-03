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
