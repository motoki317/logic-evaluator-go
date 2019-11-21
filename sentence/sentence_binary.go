package sentence

import (
	"github.com/motoki317/logic-evaluator-go/operator"
)

type BinarySentence struct {
	a  *Sentence
	op operator.Operator
	b  *Sentence
}

func (s BinarySentence) Value() (bool, error) {
	first, err := (*s.a).Value()
	if err != nil {
		return false, err
	}
	second, err := (*s.b).Value()
	if err != nil {
		return false, err
	}
	return s.op.Eval(first, second)
}

func (s BinarySentence) String() string {
	return "(" + (*s.a).String() + string(s.op) + (*s.b).String() + ")"
}

func NewBinarySentence(a *Sentence, op operator.Operator, b *Sentence) Sentence {
	return BinarySentence{
		a:  a,
		op: op,
		b:  b,
	}
}
