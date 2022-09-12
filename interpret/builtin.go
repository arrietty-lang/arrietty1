package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
	"strconv"
	"strings"
)

func IsBuiltInFunc(ident string) bool {
	builtIns := []string{
		"strlen",
		"len",
		"append",
		"print",
		"itos",
		"split",
		"keys",
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
	case "append":
		return Append(mem, args)
	case "print":
		return Print(mem, args)
	case "itos":
		return ItoS(mem, args)
	case "split":
		return Split(mem, args)
	case "keys":
		return Keys(mem, args)
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

func Append(mem *Memory, args []*analyze.ExprLevel) (*Object, error) {
	objs, err := args2Objs(mem, args)
	if err != nil {
		return nil, err
	}
	objs[0].L = append(objs[0].L, objs[1])
	return nil, nil
}

func Print(mem *Memory, args []*analyze.ExprLevel) (*Object, error) {
	objs, err := args2Objs(mem, args)
	if err != nil {
		return nil, err
	}
	fmt.Print(objs[0].S)
	return nil, nil
}

func ItoS(mem *Memory, args []*analyze.ExprLevel) (*Object, error) {
	objs, err := args2Objs(mem, args)
	if err != nil {
		return nil, err
	}
	return NewStringObject(strconv.Itoa(objs[0].I)), nil
}

func Split(mem *Memory, args []*analyze.ExprLevel) (*Object, error) {
	objs, err := args2Objs(mem, args)
	if err != nil {
		return nil, err
	}

	var splits []*Object
	for _, s := range strings.Split(objs[0].S, objs[1].S) {
		splits = append(splits, NewStringObject(s))
	}
	return NewListObject(splits), nil
}

func Keys(mem *Memory, args []*analyze.ExprLevel) (*Object, error) {
	objs, err := args2Objs(mem, args)
	if err != nil {
		return nil, err
	}
	var keys []*Object
	for k := range objs[0].D {
		keys = append(keys, NewStringObject(k))
	}

	return NewListObject(keys), nil
}
