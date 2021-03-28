package problem

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca"
)

type meaningProblem struct {
	Word  voca.Word
	Input test.Input
}

func Meaning(word voca.Word, input test.Input) test.Problem {
	p := new(meaningProblem)
	p.Word = word
	p.Input = input
	return p
}

func (p *meaningProblem) Question() voca.Unit {
	return p.Word.Spell()
}

func (p *meaningProblem) IsCorrect(userAnswer string) bool {
	return p.Word.Mean().IsRight(userAnswer)
}
