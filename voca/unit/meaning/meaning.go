package meaning

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/unit"
)

type meaning struct {
	Class voca.Class
	Mean  string
}

func New(class voca.Class, mean string) unit.Meaning {
	return meaning{Class: class, Mean: mean}
}

func (m meaning) IsRight(input string) bool {
	return m.Mean == input
}

func (m meaning) Format() string {
	return m.Mean
}
