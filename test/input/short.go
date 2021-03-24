package input

import "github.com/simp7/wordReminder-core/test"

type short struct {
}

func Short() test.Input {
	s := new(short)
	return s
}
