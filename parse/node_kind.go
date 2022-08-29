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

	Not    // !
	Assign // =
	And    // &&
	Or     // ||
	Eq     // ==
	Ne     // !=
	Lt     // <
	Le     // <=
	Gt     // >
	Ge     // >=
	Add    // +
	Sub    // -
	Mul    // *
	Div    // /
	Mod    // %

	Ident
	Call
	Float
	Int
	String
	Raw
	Array
	Dict
	KV
	True
	False
	Null

	Args
	Params

	//Comment
	//White
	//Newline
)
