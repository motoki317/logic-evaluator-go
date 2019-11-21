package logic

import (
	"github.com/motoki317/logic-evaluator-go/logic/base"
	"strconv"
	"strings"
)

type Variables map[string]bool

type variable struct {
	name  string
	value *base.Bool
}

func (v Variables) String() string {
	vars := make([]string, 0, len(v))
	for name, b := range v {
		vars = append(vars, name+"="+strconv.FormatBool(b))
	}
	return "(" + strings.Join(vars, ", ") + ")"
}

func (i *Interpreter) makeOrderedVariables() []*variable {
	variables := make([]*variable, 0, len(i.variables))
	for name, value := range i.variables {
		variables = append(variables, &variable{
			name:  name,
			value: value,
		})
	}
	return variables
}
