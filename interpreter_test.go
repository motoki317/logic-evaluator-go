package logic_evaluator_go_test

import (
	logic "github.com/motoki317/logic-evaluator-go"
	"testing"
)

type testCase struct {
	input       string
	satisfiable bool
	tautology   bool
}

func prepareTestCases() []testCase {
	return []testCase{
		{
			input:       "a",
			satisfiable: true,
			tautology:   false,
		},
		{
			input:       "not a",
			satisfiable: true,
			tautology:   false,
		},
		{
			input:       "a and (not a)",
			satisfiable: false,
			tautology:   false,
		},
		{
			input:       "a and not a",
			satisfiable: false,
			tautology:   false,
		},
		{
			input:       "a or not a",
			satisfiable: true,
			tautology:   true,
		},
		{
			input:       "a or a",
			satisfiable: true,
			tautology:   false,
		},
		{
			input:       "a implies a",
			satisfiable: true,
			tautology:   true,
		},
		{
			input:       "(p implies q) implies (not q implies not p)",
			satisfiable: true,
			tautology:   true,
		},
		{
			input:       "a equals a",
			satisfiable: true,
			tautology:   true,
		},
		{
			input:       "a equals not a",
			satisfiable: false,
			tautology:   false,
		},
		{
			input:       "((p implies q) and (q implies r) and (r implies p)) equals ((p equals q) and (q equals r))",
			satisfiable: true,
			tautology:   true,
		},
	}
}

func TestInterpreter_IsSatisfiable(t *testing.T) {
	cases := prepareTestCases()
	for _, v := range cases {
		i, err := logic.NewInterpreter(v.input)
		if err != nil {
			t.Errorf("error on initializing interpreter: %s", err)
		}

		isSatisfiable, solution, err := i.IsSatisfiable()
		if err != nil {
			t.Errorf("error on satisfiablity check %s", err)
		}

		if isSatisfiable != v.satisfiable {
			t.Errorf("expected %v, got %v, for sentence %s", v.satisfiable, isSatisfiable, v.input)
		}

		if isSatisfiable && solution == nil {
			t.Errorf("expected solution not to be nil, got nil, for sentence %s", v.input)
		} else if !isSatisfiable && solution != nil {
			t.Errorf("expected solution to be nil, got %s, for sentence %s", solution.String(), v.input)
		}
	}
}

func TestInterpreter_IsTautology(t *testing.T) {
	cases := prepareTestCases()
	for _, v := range cases {
		i, err := logic.NewInterpreter(v.input)
		if err != nil {
			t.Errorf("error on initializing interpreter: %s", err)
		}

		isTautology, counterExample, err := i.IsTautology()
		if err != nil {
			t.Errorf("error on tautology check %s", err)
		}

		if isTautology != v.tautology {
			t.Errorf("want %v, got %v, for sentence %s", v.satisfiable, isTautology, v.input)
		}

		if isTautology && counterExample != nil {
			t.Errorf("expected counter example to be nil, got %s, for sentence %s", counterExample.String(), v.input)
		} else if !isTautology && counterExample == nil {
			t.Errorf("expected counter example not to be nil, got nil, for sentence %s", v.input)
		}
	}
}
