package problem

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca"
)

type meaningProblem struct {
	Word voca.Word
}

func Meaning(word voca.Word) test.Problem {
	p := new(meaningProblem)
	p.Word = word
	return p
}

func (p *meaningProblem) Question() voca.Unit {
	return p.Word.Spell()
}

func (p *meaningProblem) IsCorrect(userAnswer string) bool {
	return p.Word.Mean().IsRight(userAnswer)
}
