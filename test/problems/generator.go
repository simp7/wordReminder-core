package problems

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca"
)

type Generator interface {
	Generate(set voca.WordSet) test.Problems
}
