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
	Raw
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
	Dot   // .
	Comma // ,
	Colon // :
	Semi  // ;
	symbolsEnd

	operatorsBegin
	Add // +
	Sub // -
	Mul // *
	Div // /
	Mod // %

	Eq // ==
	Ne // !=
	Gt // >
	Lt // <
	Ge // >=
	Le // <=

	Assign    // =
	AddAssign // +=
	SubAssign // -=
	MulAssign // *=
	DivAssign // /=
	ModAssign // %=

	And // &&
	Or  // ||
	Not // !
	operatorsEnd

	keywordBegin // 予約後
	KWFor
	KWWhile
	KWIf
	KWElse
	KWReturn

	KWDict
	KWFloat
	KWInt
	KWString
	KWBool
	KWVoid
	keywordEnd
)
