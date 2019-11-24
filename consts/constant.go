package consts

import "errors"

type Constant rune

const (
	True  = '1'
	False = '0'
)

func Constants() []Constant {
	return []Constant{True, False}
}

func (c Constant) ReplaceableTexts() ([]string, error) {
	switch c {
	case True:
		return []string{"true"}, nil
	case False:
		return []string{"false"}, nil
	}
	return []string{}, errors.New("unknown constant: " + string(c))
}

func (c Constant) Eval() (bool, error) {
	switch c {
	case True:
		return true, nil
	case False:
		return false, nil
	}
	return false, errors.New("unknown constant: " + string(c))
}

func (c Constant) String() string {
	switch c {
	case True:
		return "true"
	case False:
		return "false"
	}
	return string(c)
}
