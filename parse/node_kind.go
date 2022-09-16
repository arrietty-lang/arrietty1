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

	Import

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
	Any

	Args
	Params
	Param

	Access
	Parenthesis

	Comment
	//White
	//Newline
)

var nodeKinds = [...]string{
	Block:        "Block",
	Return:       "Return",
	If:           "If",
	IfElse:       "IfElse",
	While:        "While",
	For:          "For",
	Not:          "Not",
	Plus:         "Plus",
	Minus:        "Minus",
	And:          "And",
	Or:           "Or",
	Eq:           "Eq",
	Ne:           "Ne",
	Lt:           "Lt",
	Le:           "Le",
	Gt:           "Gt",
	Ge:           "Ge",
	Add:          "Add",
	Sub:          "Sub",
	Mul:          "Mul",
	Div:          "Div",
	Mod:          "Mod",
	FuncDef:      "FuncDef",
	VarDecl:      "VarDecl",
	ShortVarDecl: "ShortVarDecl",
	Assign:       "Assign",
	Ident:        "Ident",
	Call:         "Call",
	Float:        "Float",
	Int:          "Int",
	String:       "String",
	RawString:    "RawString",
	List:         "List",
	Dict:         "Dict",
	KV:           "KV",
	Bool:         "Bool",
	True:         "True",
	False:        "False",
	Void:         "Void",
	Null:         "Null",
	Any:          "Any",
	Args:         "Args",
	Params:       "Params",
	Param:        "Param",
	Access:       "Access",
	Parenthesis:  "Parenthesis",
	Comment:      "Comment",
}

func (n NodeKind) String() string {
	return nodeKinds[n]
}
