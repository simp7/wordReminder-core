package meaning

import (
	"github.com/simp7/wordReminder-core/voca"
)

func New(class voca.Class, mean string) voca.Meaning {
	return voca.Meaning{Class: class, Mean: mean}
}
