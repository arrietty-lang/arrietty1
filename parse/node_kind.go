package parse

type NodeKind int

const (
	_ NodeKind = iota

	Block
	Return
	If
	IfElse
	While
	For

	Not // !
	Plus
	Minus

	And // &&
	Or  // ||
	Eq  // ==
	Ne  // !=
	Lt  // <
	Le  // <=
	Gt  // >
	Ge  // >=
	Add // +
	Sub // -
	Mul // *
	Div // /
	Mod // %

	FuncDef
	VarDecl
	ShortVarDecl
	Assign // =

	Ident
	Call
	Float
	Int
	String
	RawString
	List
	Dict
	KV
	Bool
	True
	False
	Void
	Null

	Args
	Params
	Param

	Access
	Parenthesis

	//Comment
	//White
	//Newline
)
