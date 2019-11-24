package logic_evaluator_go

import (
	"errors"
	"github.com/motoki317/logic-evaluator-go/base"
	"github.com/motoki317/logic-evaluator-go/consts"
	"github.com/motoki317/logic-evaluator-go/operator"
	"regexp"
	"strings"
)

func isOperatorChar(char rune) bool {
	for _, op := range operator.OrderedOperators() {
		if char == rune(op) {
			return true
		}
	}
	return false
}

func isVariableChar(char rune) bool {
	if char == '(' || char == ')' {
		return false
	}
	return !isOperatorChar(char)
}

func isVariableName(text string) bool {
	for _, ch := range []rune(text) {
		if !isVariableChar(ch) {
			return false
		}
	}
	return true
}

func isConstantChar(char rune) bool {
	for _, c := range consts.Constants() {
		if char == rune(c) {
			return true
		}
	}
	return false
}

func getOperator(ch rune) (operator.Operator, error) {
	for _, op := range operator.OrderedOperators() {
		if ch == rune(op) {
			return op, nil
		}
	}
	return 'e', errors.New("internal error: unknown operator: " + string(ch))
}

func replaceTexts(text string) (string, error) {
	text = strings.ToLower(text)

	for _, op := range operator.ReverseOrderedOperators() {
		texts, err := op.ReplaceableTexts()
		if err != nil {
			return "", err
		}
		for _, replaceable := range texts {
			text = strings.ReplaceAll(text, replaceable, string(op))
		}
	}

	for _, constant := range consts.Constants() {
		texts, err := constant.ReplaceableTexts()
		if err != nil {
			return "", err
		}
		for _, replaceable := range texts {
			text = strings.ReplaceAll(text, replaceable, string(constant))
		}
	}

	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, " ", "")
	return text, nil
}

func hasValidParentheses(text string) bool {
	openings := 0
	closings := 0
	for _, ch := range []rune(text) {
		if ch == '(' {
			openings++
		}
		if ch == ')' {
			closings++
		}
	}
	return openings == closings
}

func getVariables(text string) map[string]*base.Bool {
	nonVariableChars := "\\(\\)"
	for _, op := range operator.OrderedOperators() {
		nonVariableChars += string(op)
	}
	for _, c := range consts.Constants() {
		nonVariableChars += string(c)
	}

	r := regexp.MustCompile("([^" + nonVariableChars + "]+)")
	vars := r.FindAllString(text, -1)

	variablesMap := make(map[string]*base.Bool)
	for _, v := range vars {
		variablesMap[v] = &base.Bool{}
	}

	return variablesMap
}

/*
Get the index of closing parenthesis corresponding to the given opening index.
Returns -1 if not found.
*/
func getClosingParenthesisIndex(text string, opening int) int {
	openingDepth := 0
	depth := 0
	chars := []rune(text)
	for i, ch := range chars {
		if ch == ')' {
			depth--
		}

		if i == opening {
			openingDepth = depth
		}

		if ch == '(' {
			depth++
		}
	}

	depth = 0
	for i, ch := range chars {
		if ch == ')' {
			depth--
			if opening < i && depth == openingDepth {
				return i
			}
		}

		if ch == '(' {
			depth++
		}
	}

	return -1
}

func getOperatorIndices(text string) (map[operator.Operator]int, error) {
	chars := []rune(text)
	operatorIndices := make(map[operator.Operator]int)
	parDepth := 0
	for i, ch := range chars {
		// Parentheses
		if ch == '(' {
			parDepth++
			continue
		} else if ch == ')' {
			parDepth--
			continue
		}

		// Do not evaluate characters inside parentheses immediately
		if parDepth != 0 {
			continue
		}

		// Check first operator indices
		if isOperatorChar(ch) {
			op, err := getOperator(ch)
			if err != nil {
				return nil, err
			}
			if _, ok := operatorIndices[op]; !ok {
				operatorIndices[op] = i
			}
		}
	}
	return operatorIndices, nil
}
