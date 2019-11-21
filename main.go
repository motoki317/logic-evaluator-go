/*
Evaluates the given logic in string, and checks if the logic is satisfiable and if the logic is a tautology.
No algorithm for SAT solving used, this program directly checks all 2^n patterns (n: number of variables).

Some example inputs:
α⇒¬¬α
((p⇒q)∧(q⇒r)∧(r⇒p))⇔((p⇔q)∧(q⇔r))
((α⇒β)∧α)⇒β
(¬α∧¬β∧¬γ)∨(α∧¬β∧γ)∨(α∧β∧γ)
*/

package logic_evaluator_go

import (
	"bufio"
	"fmt"
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

	interpreter, err := NewInterpreter(text)
	if err != nil {
		panic(err)
	}

	if interpreter.Text() != text {
		fmt.Println("Replaced text: " + interpreter.Text())
	}
	fmt.Println("Evaluated: " + interpreter.EvaluatedText())

	fmt.Println("-----")

	solution, err := interpreter.CheckSatisfiable()
	if err != nil {
		panic(err)
	}

	if solution != nil {
		fmt.Println("This logic is satisfiable.")
		fmt.Println("Possible solution: " + solution.String())
	} else {
		fmt.Println("This logic is NOT satisfiable.")
	}

	fmt.Println("-----")

	counterExample, err := interpreter.CheckTautology()
	if err != nil {
		panic(err)
	}

	if counterExample == nil {
		fmt.Println("This logic is a tautology.")
	} else {
		fmt.Println("This logic is NOT a tautology.")
		fmt.Println("Counter example: " + counterExample.String())
	}
}
