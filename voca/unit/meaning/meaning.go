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

func (m meaning) IsEqual(u voca.Unit) bool {
	return m == u
}

func (m meaning) Format() string {
	return m.Mean
}
