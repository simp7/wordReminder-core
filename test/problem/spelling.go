package problem

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca"
)

type spellingProblem struct {
	Word  voca.Word
	Input test.Input
}

func Spelling(word voca.Word, input test.Input) test.Problem {
	s := new(spellingProblem)
	s.Word = word
	s.Input = input
	return s
}

func (s spellingProblem) Question() voca.Unit {
	return s.Word.Mean()
}

func (s spellingProblem) IsCorrect(userAnswer string) bool {
	return s.Word.Spell().IsRight(userAnswer)
}
