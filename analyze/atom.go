package analyze

import "github.com/x0y14/arrietty/parse"

type Atom struct {
	Kind AtomKind
	S    string
	I    int
	F    float64
}

func (a *Atom) GetType() (*DataType, error) {
	switch a.Kind {
	case AFloat:
		return &DataType{Type: TFloat}, nil
	case AInt:
		return &DataType{Type: TInt}, nil
	case AString:
		return &DataType{Type: TString}, nil
	case ATrue, AFalse:
		return &DataType{Type: TBool}, nil
	case ANull:
		return &DataType{Type: TVoid}, nil
	}
	return nil, NewUndefinedErr("atom?")
}

func NewAtom(node *parse.Node) (*Atom, error) {
	switch node.Kind {
	case parse.Float:
		return &Atom{Kind: AFloat, F: node.F}, nil
	case parse.Int:
		return &Atom{Kind: AInt, I: node.I}, nil
	case parse.String:
		return &Atom{Kind: AString, S: node.S}, nil
	case parse.True:
		return &Atom{Kind: ATrue}, nil
	case parse.False:
		return &Atom{Kind: AFalse}, nil
	case parse.Null:
		return &Atom{Kind: ANull}, nil
	}
	return nil, NewUnexpectNodeErr(node)
}
