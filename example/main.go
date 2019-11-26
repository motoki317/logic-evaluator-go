/*
Evaluates the given logic in string, and checks if the logic is satisfiable and if the logic is a tautology.
No algorithm for SAT solving used, this program directly checks all 2^n patterns (n: number of variables).

Some example inputs:
α⇒¬¬α
((p⇒q)∧(q⇒r)∧(r⇒p))⇔((p⇔q)∧(q⇔r))
((α⇒β)∧α)⇒β
(¬α∧¬β∧¬γ)∨(α∧¬β∧γ)∨(α∧β∧γ)
*/

package main

import (
	"bufio"
	"fmt"
	"github.com/motoki317/logic-evaluator-go"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	text, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	text = strings.ReplaceAll(text, "\n", "")

	interpreter, err := logic_evaluator_go.NewInterpreter(text)
	if err != nil {
		panic(err)
	}

	if interpreter.Text() != text {
		fmt.Println("Replaced text: " + interpreter.Text())
	}
	fmt.Println("Evaluated: " + interpreter.EvaluatedText())

	fmt.Println("-----")

	isSatisfiable, solution, err := interpreter.IsSatisfiable()
	if err != nil {
		panic(err)
	}

	if isSatisfiable {
		fmt.Println("This logic is satisfiable.")
		fmt.Println("Possible solution: " + solution.String())
	} else {
		fmt.Println("This logic is NOT satisfiable.")
	}

	fmt.Println("-----")

	isTautology, counterExample, err := interpreter.IsTautology()
	if err != nil {
		panic(err)
	}

	if isTautology {
		fmt.Println("This logic is a tautology.")
	} else {
		fmt.Println("This logic is NOT a tautology.")
		fmt.Println("Counter example: " + counterExample.String())
	}
}
