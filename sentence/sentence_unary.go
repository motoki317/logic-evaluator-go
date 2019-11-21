package sentence

import (
	"github.com/motoki317/logic-evaluator-go/operator"
)

type UnarySentence struct {
	a  *Sentence
	op operator.Operator
}

func (s UnarySentence) Value() (bool, error) {
	value, err := (*s.a).Value()
	if err != nil {
		return false, err
	}
	return s.op.Eval(value, true)
}

func (s UnarySentence) String() string {
	return "(" + string(s.op) + (*s.a).String() + ")"
}

func NewUnarySentence(a *Sentence, op operator.Operator) Sentence {
	return UnarySentence{
		a:  a,
		op: op,
	}
}
