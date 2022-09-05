package analyze

type SemanticNodeKind int

const (
	_ SemanticNodeKind = iota
	SReturn
	SIf
	SIfElse
	SWhile
	SFor

	SNot
	SPlus
	SMinus

	SAnd
	SOr

	SEq
	SNe
	SLt
	SLe
	SGt
	SGe

	SAdd
	SSub
	SMul
	SDiv
	SMod

	SFuncDef
	SParams
	SParam

	SVarDecl
	SAssign

	SIdent
	SCall
	SArgs

	SFloat
	SInt
	SString // raw string...?
	SList
	SDict
	SKV
	SBool
	STrue
	SFalse
	SVoid
	SNull

	SAccess      // []
	SParenthesis // ()
)
