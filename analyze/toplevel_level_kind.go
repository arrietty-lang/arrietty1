package analyze

type ToplevelKind int

const (
	_ ToplevelKind = iota
	TPReturn
	TPIf
	TPIfElse
	TPWhile
	TPFor

	TPNot
	TPPlus
	TPMinus

	TPAnd
	TPOr

	TPEq
	TPNe
	TPLt
	TPLe
	TPGt
	TPGe

	TPAdd
	TPSub
	TPMul
	TPDiv
	TPMod

	TPFuncDef
	TPParams
	TPParam

	TPVarDecl
	TPAssign

	TPIdent
	TPCall
	TPArgs

	TPFloat
	TPInt
	TPString // raw string...?
	TPList
	TPDict
	TPKv
	TPBool
	TPTrue
	TPFalse
	TPVoid
	TPNull

	TPComment

	TPAccess      // []
	TPParenthesis // ()
)

var toplevelKinds = [...]string{
	TPReturn:      "Return",
	TPIf:          "If",
	TPIfElse:      "IfElse",
	TPWhile:       "While",
	TPFor:         "For",
	TPNot:         "Not",
	TPPlus:        "Plus",
	TPMinus:       "Minus",
	TPAnd:         "And",
	TPOr:          "Or",
	TPEq:          "Eq",
	TPNe:          "Ne",
	TPLt:          "Lt",
	TPLe:          "Le",
	TPGt:          "Gt",
	TPGe:          "Ge",
	TPAdd:         "Add",
	TPSub:         "Sub",
	TPMul:         "Mul",
	TPDiv:         "Div",
	TPMod:         "Mod",
	TPFuncDef:     "FuncDef",
	TPParams:      "Params",
	TPParam:       "Param",
	TPVarDecl:     "VarDecl",
	TPAssign:      "AssignTo",
	TPIdent:       "Ident",
	TPCall:        "Call",
	TPArgs:        "Args",
	TPFloat:       "Float",
	TPInt:         "Int",
	TPString:      "String",
	TPList:        "List",
	TPDict:        "Dict",
	TPKv:          "Kv",
	TPBool:        "Bool",
	TPTrue:        "True",
	TPFalse:       "False",
	TPVoid:        "Void",
	TPNull:        "Null",
	TPComment:     "Comment",
	TPAccess:      "Access",
	TPParenthesis: "Parenthesis",
}

func (t ToplevelKind) String() string {
	return toplevelKinds[t]
}
