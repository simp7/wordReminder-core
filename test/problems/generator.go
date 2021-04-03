package problems

import "github.com/simp7/wordReminder-core/test"

type Generator interface {
	Do() test.Problems
}
