package sentence

import "fmt"

type Sentence interface {
	fmt.Stringer
	Value() (bool, error)
}
