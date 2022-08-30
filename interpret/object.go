package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type Object struct {
	Kind ObjectKind

	*Fn
	*Literal
}

func (o *Object) isLiteral() bool {
	return o.Kind == ObjLiteral
}

func (o *Object) String() string {
	switch o.Kind {
	case ObjFn:
		return "Function"
	case ObjLiteral:
		switch o.Literal.Kind {
		case Ident:
			return o.Str
		case Float:
			return fmt.Sprintf("%v", o.NumFloat)
		case Int:
			return fmt.Sprintf("%v", o.NumInt)
		case String, Raw:
			return fmt.Sprintf("%s", o.Str)
		case True:
			return fmt.Sprintf("true")
		case False:
			return fmt.Sprintf("false")
		case Null:
			return fmt.Sprintf("null")
		}
	}
	return "UNKNOWN"
}

func NewFn(params []*parse.Node, body *parse.Node) *Object {
	var objParams []*Object
	for _, param := range params {
		objParams = append(objParams, NewIdent(param.Str))
	}

	return &Object{
		Kind: ObjFn,
		Fn: &Fn{
			Params: objParams,
			Body:   body,
			Local:  NewStorage(),
		},
		Literal: nil,
	}
}

func NewIdent(s string) *Object {
	return &Object{
		Kind: ObjLiteral,
		Fn:   nil,
		Literal: &Literal{
			Kind: Ident,
			Str:  s,
		},
	}
}

func NewFloat(n float64) *Object {
	return &Object{
		Kind: ObjLiteral,
		Fn:   nil,
		Literal: &Literal{
			Kind:     Float,
			NumFloat: n,
		},
	}
}

func NewInt(n int) *Object {
	return &Object{
		Kind: ObjLiteral,
		Fn:   nil,
		Literal: &Literal{
			Kind:   Int,
			NumInt: n,
		},
	}
}

func NewString(s string) *Object {
	return &Object{
		Kind: ObjLiteral,
		Fn:   nil,
		Literal: &Literal{
			Kind: String,
			Str:  s,
		},
	}
}

func NewRaw(s string) *Object {
	return &Object{
		Kind: ObjLiteral,
		Fn:   nil,
		Literal: &Literal{
			Kind: Raw,
			Str:  s,
		},
	}
}

func NewArray(items []*parse.Node) (*Object, error) {
	var objs []*Object
	for _, item := range items {
		o, err := eval(globalStorage, item)
		if err != nil {
			return nil, err
		}
		objs = append(objs, o)
	}

	return &Object{
		Kind: ObjLiteral,
		Fn:   nil,
		Literal: &Literal{
			Kind:  Array,
			Items: objs,
		},
	}, nil
}

func NewDict(kv []*parse.Node) (*Object, error) {
	var m map[string]*Object

	for _, k := range kv {
		valueObj, err := eval(globalStorage, k.Rhs)
		if err != nil {
			return nil, err
		}
		m[k.Lhs.Str] = valueObj
	}

	return &Object{
		Kind: ObjLiteral,
		Fn:   nil,
		Literal: &Literal{
			Kind: Dict,
			KVS:  m,
		},
	}, nil
}

func NewTrue() *Object {
	return &Object{
		Kind: ObjLiteral,
		Fn:   nil,
		Literal: &Literal{
			Kind: True,
		},
	}
}

func NewFalse() *Object {
	return &Object{
		Kind: ObjLiteral,
		Fn:   nil,
		Literal: &Literal{
			Kind: False,
		},
	}
}

func NewNull() *Object {
	return &Object{
		Kind: ObjLiteral,
		Fn:   nil,
		Literal: &Literal{
			Kind: Null,
		},
	}
}
