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
	True
	False
	Null

	Args
	Params

	Access

	//Comment
	//White
	//Newline
)
