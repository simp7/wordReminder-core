package problems

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca"
)

func Meaning(words voca.WordSet, amount int, function func(...test.Problem) test.Problems) test.Problems {

	pSelector := Selector(words)
	selected := pSelector.Select(amount)

	problems := make([]test.Problem, selected.Len())

	for i := 0; i < selected.Len(); i++ {
		problems[i] = test.Meaning(selected.Get(i))
	}

	return function(problems...)

}
