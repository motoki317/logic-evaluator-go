package logic_evaluator_go

import (
	"github.com/motoki317/logic-evaluator-go/sentence"
)

/*
Checks if this logic is satisfiable.
Runs through all 2^n (n: number of variables) combinations and directly checks them.
If this is satisfiable, then returns true and the variables map on one of those instances.
If not, then returns false.
*/
func (i *Interpreter) IsSatisfiable() (res bool, solution Variables, err error) {
	variables := i.makeOrderedVariables()
	res, err = dfsSatisfiable(i.sentence, variables, 0)
	if err != nil {
		return
	}

	if res {
		// Replicate the result
		solution = make(map[string]bool)
		for _, v := range variables {
			solution[v.name] = v.value.Get()
		}
		return
	}
	return
}

/*
Checks if the given logic is satisfiable using depth-first-search algorithm.
Stops the search as soon as one satisfiable condition is found.
If then, example that satisfies the logic is stored in the given "variables" variable.
*/
func dfsSatisfiable(sentence *sentence.Sentence, variables []*variable, depth int) (bool, error) {
	if depth == len(variables) {
		return (*sentence).Value()
	}

	variable := variables[depth]
	variable.value.Set(false)
	result, err := dfsSatisfiable(sentence, variables, depth+1)
	if err != nil {
		return false, err
	}
	if result {
		return true, nil
	}

	variable.value.Set(true)
	return dfsSatisfiable(sentence, variables, depth+1)
}
