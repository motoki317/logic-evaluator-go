package sentence

import "github.com/motoki317/logic-evaluator-go/consts"

type ConstantSentence struct {
	value consts.Constant
}

func (s ConstantSentence) Value() (bool, error) {
	return s.value.Eval()
}

func (s ConstantSentence) String() string {
	return "(" + s.value.String() + ")"
}

func NewConstantSentence(c consts.Constant) Sentence {
	return ConstantSentence{value: c}
}
