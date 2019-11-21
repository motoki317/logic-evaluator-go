package sentence

import "github.com/motoki317/logic-evaluator-go/logic/base"

type ValueSentence struct {
	name  string
	value *base.Bool
}

func (s ValueSentence) Value() (bool, error) {
	return s.value.Get(), nil
}

func (s ValueSentence) String() string {
	return s.name
}

func NewValueSentence(name string, value *base.Bool) Sentence {
	return ValueSentence{
		name:  name,
		value: value,
	}
}
