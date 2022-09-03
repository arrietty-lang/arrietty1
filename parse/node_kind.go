package parse

type NodeKind int

const (
	_ NodeKind = iota

	Function
	Block
	Return
	If
	IfElse
	While
	For

	Not // !
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
	Raw
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

	//Comment
	//White
	//Newline
)
