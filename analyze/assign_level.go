package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type AssignLevel struct {
	Kind       AssignLevelKind
	AndOrLevel *AndOrLevel
	VarDecl    *VarDecl
	Assignment *Assignment
}

func (a *AssignLevel) GetType() (*DataType, error) {
	switch a.Kind {
	case AAndOrLevel:
		return a.AndOrLevel.GetType()
	}
	return nil, fmt.Errorf("assignLv type error")
}

func NewAssignLevel(node *parse.Node) (*AssignLevel, error) {
	switch node.Kind {
	case parse.VarDecl:
		return newAssignLevelVarDecl(node)
	case parse.Assign, parse.ShortVarDecl:
		return newAssignLevelAssignment(node)
	}
	a, err := NewAndOrLevel(node)
	if err != nil {
		return nil, err
	}
	return &AssignLevel{Kind: AAndOrLevel, AndOrLevel: a}, nil
}

func newAssignLevelVarDecl(node *parse.Node) (*AssignLevel, error) {
	decl, err := NewVarDecl(node)
	if err != nil {
		return nil, err
	}

	return &AssignLevel{Kind: AVarDecl, VarDecl: decl}, nil
}

func newAssignLevelAssignment(node *parse.Node) (*AssignLevel, error) {
	as, err := NewAssignment(node)
	if err != nil {
		return nil, err
	}
	return &AssignLevel{Kind: AAssign, Assignment: as}, nil
}
