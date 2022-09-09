package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

type Object struct {
	Kind ObjectKind
	S    string
	F    float64
	I    int
	D    map[string]*Object
	L    []*Object
}

func NewIntObject(i int) *Object {
	return &Object{Kind: OInt, I: i}
}
func NewFloatObject(f float64) *Object {
	return &Object{Kind: OFloat, F: f}
}
func NewStringObject(s string) *Object {
	return &Object{Kind: OString, S: s}
}
func NewTrueObject() *Object {
	return &Object{Kind: OTrue}
}
func NewFalseObject() *Object {
	return &Object{Kind: OFalse}
}
func NewNullObject() *Object {
	return &Object{Kind: ONull}
}

func ConvertAtomToObject(atom *analyze.Atom) (*Object, error) {
	switch atom.Kind {
	case analyze.AFloat:
		return &Object{Kind: OFloat, F: atom.F}, nil
	case analyze.AInt:
		return &Object{Kind: OInt, I: atom.I}, nil
	case analyze.AString:
		return &Object{Kind: OString, S: atom.S}, nil
	case analyze.ATrue:
		return &Object{Kind: OTrue}, nil
	case analyze.AFalse:
		return &Object{Kind: OFalse}, nil
	case analyze.ANull:
		return &Object{Kind: ONull}, nil
	}
	return nil, fmt.Errorf("unimplemented: %s", atom.Kind.String())
}

func ConvertDictToObject() {
	// todo
}
func ConvertListToObject() {
	// todo
}
