# logic-evaluator-go
Simple logic evaluator in Go
Evaluates the given logic in string, and checks if the logic is satisfiable and if the logic is a tautology.
No algorithm for SAT solving used, this program directly checks all 2^n patterns (n: number of variables).

## Example Usage
```go
package main

import (
    "fmt"
    logic "github.com/motoki317/logic-evaluator-go"
)

func main() {
    text := "a -> not not a"

	interpreter, _ := logic.NewInterpreter(text)

	solution, _ := interpreter.CheckSatisfiable()

	if solution != nil {
		fmt.Println("This logic is satisfiable.")
		fmt.Println("Possible solution: " + solution.String())
	} else {
		fmt.Println("This logic is NOT satisfiable.")
	}

	counterExample, _ := interpreter.CheckTautology()

	if counterExample == nil {
		fmt.Println("This logic is a tautology.")
	} else {
		fmt.Println("This logic is NOT a tautology.")
		fmt.Println("Counter example: " + counterExample.String())
	}
}
```

### Allowed Operators
You can use these operators to construct a sentence. Other characters are not supported and will be considered variables.
- NOT `¬`, `~`, `!`, `not`
- AND `∧`, `/\`, `&`, `and`
- OR `∨`, `\/`, `|`, `or`
- IMPLY `⇒`, `→`, `->`, `=>`
- EQUIVALENT `⇔`, `↔`, `<->`, `<=>`

### Example Inputs
- `α⇒¬¬α`
- `((p⇒q)∧(q⇒r)∧(r⇒p))⇔((p⇔q)∧(q⇔r))`
- `((α⇒β)∧α)⇒β`
- `(¬α∧¬β∧¬γ)∨(α∧¬β∧γ)∨(α∧β∧γ)`
- `(P -> Q) -> (~Q -> ~P)`

### Example Output
```
Input: (P -> Q) -> (~Q -> ~P)
Replaced text: (P⇒Q)⇒(¬Q⇒¬P)
Evaluated: ((P⇒Q)⇒((¬Q)⇒(¬P)))
-----
This logic is satisfiable.
Possible solution: {P=false, Q=false}
-----
This logic is a tautology.
```
