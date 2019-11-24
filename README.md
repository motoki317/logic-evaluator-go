# logic-evaluator-go
This simple logic evaluator evaluates the given logic in string, and checks if the logic is satisfiable and if the logic is a tautology.
No advanced algorithms used, this program directly checks all 2^n patterns (n: number of variables).

## Usage

### Operators
These strings are considered as operators.
- NOT `¬`, `~`, `!`, `not`
- AND `∧`, `/\ `, `&`, `and`
- OR `∨`, `\/`, `|`, `or`
- IMPLY `⇒`, `→`, `->`, `=>`
- EQUIVALENT `⇔`, `↔`, `<->`, `<=>`

### Constants
These strings are considered as constants.
- TRUE `true`, `1`
- FALSE `false`, `0`

### Variables
Any string other than above is regarded as a variable.

## Examples

### Code
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

### Example Inputs
- `α⇒¬¬α`
- `((p⇒q)∧(q⇒r)∧(r⇒p))⇔((p⇔q)∧(q⇔r))`
- `((α⇒β)∧α)⇒β`
- `(¬α∧¬β∧¬γ)∨(α∧¬β∧γ)∨(α∧β∧γ)`
- `(P -> Q) -> (~Q -> ~P)`
- `true AND α -> α`

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
