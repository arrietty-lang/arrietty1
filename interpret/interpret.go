package interpret

import (
	"github.com/x0y14/arrietty/analyze"
)

var FileMem *Memory

func Interpret(script map[string]*analyze.TopLevel) (*Object, error) {
	FileMem = NewMemory(nil, script)
	mainFunc, err := FileMem.GetFunc("main")
	if err != nil {
		return nil, err
	}
	v, _, err := ExecFunction(FileMem, mainFunc, nil)
	return v, err
}
