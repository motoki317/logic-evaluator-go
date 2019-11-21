package logic_evaluator_go

import (
	"errors"
	"github.com/motoki317/logic-evaluator-go/base"
	"github.com/motoki317/logic-evaluator-go/operator"
	"github.com/motoki317/logic-evaluator-go/sentence"
	"regexp"
)

type Interpreter struct {
	text      string
	sentence  *sentence.Sentence
	variables map[string]*base.Bool
}

func (i *Interpreter) Text() string {
	return i.text
}

func (i *Interpreter) EvaluatedText() string {
	return (*i.sentence).String()
}

func (i *Interpreter) Value() (bool, error) {
	return (*i.sentence).Value()
}

func NewInterpreter(text string) (*Interpreter, error) {
	if !hasValidParentheses(text) {
		return nil, errors.New("number of opening/closing parentheses do not match")
	}

	replacedText, err := replaceTexts(text)
	if err != nil {
		return nil, err
	}

	interpreter := &Interpreter{
		text: replacedText,
	}

	interpreter.variables = getVariables(interpreter.text)
	s, err := interpretSentence(interpreter.text, &interpreter.variables)
	if err != nil {
		return nil, err
	}
	interpreter.sentence = s

	return interpreter, nil
}

func getVariables(text string) map[string]*base.Bool {
	nonVariableChars := "\\(\\)"
	for _, v := range operator.OrderedOperators() {
		nonVariableChars += string(v)
	}

	r := regexp.MustCompile("([^" + nonVariableChars + "]+)")
	vars := r.FindAllString(text, -1)

	variablesMap := make(map[string]*base.Bool)
	for _, v := range vars {
		variablesMap[v] = &base.Bool{}
	}

	return variablesMap
}

func interpretSentence(input string, variables *map[string]*base.Bool) (*sentence.Sentence, error) {
	// Check if given input is variable
	if isVariableName(input) {
		s := sentence.NewValueSentence(input, (*variables)[input])
		return &s, nil
	}

	// Check if given input is closed with parentheses
	chars := []rune(input)
	if chars[0] == '(' && chars[len(chars)-1] == ')' && getClosingParenthesisIndex(input, 0) == len(chars)-1 {
		return interpretSentence(string(chars[1:len(chars)-1]), variables)
	}

	operatorIndices, err := getOperatorIndices(input)
	if err != nil {
		return nil, err
	}

	// Operators
	for _, op := range operator.ReverseOrderedOperators() {
		if _, ok := operatorIndices[op]; !ok {
			continue
		}

		switch op {
		case operator.Not:
			// Only evaluate NOT operator when located at the beginning of input string
			if i := operatorIndices[op]; i != 0 {
				continue
			}

			internal, err := interpretSentence(string(chars[1:]), variables)
			if err != nil {
				return nil, err
			}
			s := sentence.NewUnarySentence(internal, op)
			return &s, nil
		case operator.And:
			fallthrough
		case operator.Or:
			fallthrough
		case operator.Imply:
			fallthrough
		case operator.Equiv:
			location := operatorIndices[op]
			firstSentence, err := interpretSentence(string(chars[0:location]), variables)
			if err != nil {
				return nil, err
			}
			secondSentence, err := interpretSentence(string(chars[location+1:]), variables)
			if err != nil {
				return nil, err
			}

			s := sentence.NewBinarySentence(firstSentence, op, secondSentence)
			return &s, nil
		}
	}

	return nil, errors.New("internal error on interpreting a sentence: " + input)
}
