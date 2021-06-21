package problems

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca"
)

func New(words voca.WordSet, amount int, gen func(voca.WordSet) *generator, function func(...test.Problem) test.Problems) test.Problems {

	sel := Selector(words)
	selected := sel.Select(amount)
	problems := gen(selected).Generate()

	return function(problems...)

}

func Meaning(words voca.WordSet, amount int, structure func(...test.Problem) test.Problems) test.Problems {
	return New(words, amount, meaning, structure)
}

func Spelling(words voca.WordSet, amount int, structure func(...test.Problem) test.Problems) test.Problems {
	return New(words, amount, spelling, structure)
}
