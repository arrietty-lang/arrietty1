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

func (o *Object) AssignWithIndex(index *Object, value *Object) error {
	if o.Kind == ODict {
		o.D[index.S] = value
		return nil
	}
	if o.Kind == OList {
		if len(o.L) <= index.I {
			return fmt.Errorf("index out of range")
		}
		o.L[index.I] = value
		return nil
	}
	return fmt.Errorf("%s can't assgin value throw index", o.Kind.String())
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
func NewListObject(items []*Object) *Object {
	return &Object{Kind: OList, L: items}
}
func NewDictObject(kvs map[string]*Object) *Object {
	return &Object{Kind: ODict, D: kvs}
}
func NewBoolObject(b bool) *Object {
	if b {
		return NewTrueObject()
	}
	return NewFalseObject()
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

func ConvertDictToObject(mem *Memory, dictLv *analyze.DictLevel) (*Object, error) {
	d := &Object{Kind: ODict}
	d.D = map[string]*Object{}
	for _, kvLv := range dictLv.KVs {
		value, err := evalUnary(mem, kvLv.Value)
		if err != nil {
			return nil, err
		}
		d.D[kvLv.Key] = value
	}
	return d, nil
}
func ConvertListToObject(mem *Memory, listLv *analyze.ListLevel) (*Object, error) {
	l := &Object{Kind: OList}
	for _, unaryLv := range listLv.Items {
		item, err := evalUnary(mem, unaryLv)
		if err != nil {
			return nil, err
		}
		l.L = append(l.L, item)
	}
	return l, nil
}
