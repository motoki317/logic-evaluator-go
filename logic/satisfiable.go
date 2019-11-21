package logic

import (
	"github.com/motoki317/logic-evaluator-go/logic/sentence"
)

/*
Checks if this logic is satisfiable.
Runs through all 2^n (n: number of variables) combinations and directly checks them.
If this is satisfiable, then returns the variables map on one of those instances.
If not, then returns null.
*/
func (i *Interpreter) CheckSatisfiable() (Variables, error) {
	variables := i.makeOrderedVariables()
	result, err := isSatisfiable(i.sentence, variables, 0)
	if err != nil {
		return nil, err
	}

	if result {
		// Replicate the result
		ret := make(map[string]bool)
		for _, v := range variables {
			ret[v.name] = v.value.Get()
		}
		return ret, nil
	}
	return nil, nil
}

/*
Checks if the given logic is satisfiable using depth-first-search algorithm.
Stops the search as soon as one satisfiable condition is found.
If then, example that satisfies the logic is stored in the given "variables" variable.
*/
func isSatisfiable(sentence *sentence.Sentence, variables []*variable, depth int) (bool, error) {
	if depth == len(variables) {
		return (*sentence).Value()
	}

	variable := variables[depth]
	variable.value.Set(false)
	result, err := isSatisfiable(sentence, variables, depth+1)
	if err != nil {
		return false, err
	}
	if result {
		return true, nil
	}

	variable.value.Set(true)
	return isSatisfiable(sentence, variables, depth+1)
}
