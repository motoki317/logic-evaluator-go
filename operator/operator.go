package operator

import (
	"errors"
)

type Operator rune

const (
	Not   Operator = '¬'
	And   Operator = '∧'
	Or    Operator = '∨'
	Imply Operator = '⇒'
	Equiv Operator = '⇔'
)

func OrderedOperators() []Operator {
	return []Operator{Not, And, Or, Imply, Equiv}
}

func ReverseOrderedOperators() []Operator {
	s := OrderedOperators()
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (op Operator) ReplaceableTexts() ([]string, error) {
	switch op {
	case Not:
		return []string{"~", "!", "not"}, nil
	case And:
		return []string{"/\\", "&", "and"}, nil
	case Or:
		return []string{"\\/", "|", "or"}, nil
	case Imply:
		return []string{"→", "->", "=>", "imply", "implies"}, nil
	case Equiv:
		return []string{"↔", "<->", "<=>", "=", "equals", "equal", "is"}, nil
	}
	return []string{}, errors.New("unknown operator: " + string(op))
}

func (op Operator) Eval(a, b bool) (bool, error) {
	switch op {
	case Not:
		return !a, nil
	case And:
		return a && b, nil
	case Or:
		return a || b, nil
	case Imply:
		return !a || b, nil
	case Equiv:
		return a == b, nil
	}
	return false, errors.New("unknown operator: " + string(op))
}
