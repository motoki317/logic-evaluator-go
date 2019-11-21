package logic_evaluator_go

import "github.com/motoki317/logic-evaluator-go/sentence"

/*
Checks if this logic evaluates to true on any conditions of the variables.
Runs through all 2^n (n: number of variables) combinations and directly checks them.
If this is NOT a tautology, then returns a counter example.
If this is a tautology, returns null.
*/
func (i *Interpreter) CheckTautology() (Variables, error) {
	// Ordered variables
	variables := i.makeOrderedVariables()
	result, err := isTautology(i.sentence, variables, 0)
	if err != nil {
		return nil, err
	}

	if !result {
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
Checks if the given logic is a tautology, by checking for all combination of variables.
Stops the search as soon as one counter example is found.
If then, the counter example is stored in the given "variables" variable.
*/
func isTautology(sentence *sentence.Sentence, variables []*variable, depth int) (bool, error) {
	if depth == len(variables) {
		return (*sentence).Value()
	}

	variable := variables[depth]
	variable.value.Set(false)
	result, err := isTautology(sentence, variables, depth+1)
	if err != nil {
		return false, err
	}
	if !result {
		return false, nil
	}

	variable.value.Set(true)
	return isTautology(sentence, variables, depth+1)
}
