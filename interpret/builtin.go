package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func IsBuiltInFunc(ident string) bool {
	builtIns := []string{
		"strlen",
		"len",
	}
	for _, b := range builtIns {
		if b == ident {
			return true
		}
	}
	return false
}

func ExecBuiltIn(ident string, mem *Memory, args []*analyze.ExprLevel) (*Object, error) {
	switch ident {
	case "strlen":
		return StrLen(mem, args)
	case "len":
		return Len(mem, args)
	}
	return nil, fmt.Errorf("builtin function, %s is undefined", ident)
}

func StrLen(mem *Memory, args []*analyze.ExprLevel) (*Object, error) {
	objs, err := args2Objs(mem, args)
	if err != nil {
		return nil, err
	}
	return NewIntObject(len(objs[0].S)), nil
}

func Len(mem *Memory, args []*analyze.ExprLevel) (*Object, error) {
	objs, err := args2Objs(mem, args)
	if err != nil {
		return nil, err
	}
	return NewIntObject(len(objs[0].L)), nil
}
