package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
	"strconv"
)

func IsBuiltInFunc(ident string) bool {
	builtIns := []string{
		"strlen",
		"len",
		"append",
		"print",
		"itos",
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
	fmt.Printf(objs[0].S)
	return nil, nil
}

func ItoS(mem *Memory, args []*analyze.ExprLevel) (*Object, error) {
	objs, err := args2Objs(mem, args)
	if err != nil {
		return nil, err
	}
	return NewStringObject(strconv.Itoa(objs[0].I)), nil
}
