package problem

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca"
)

type problem struct {
	word     voca.Word
	question voca.Unit
	answer   voca.Unit
}

func Meaning(word voca.Word) test.Problem {
	p := new(problem)
	p.word = word
	p.question = p.word.Spell()
	p.answer = p.word.Mean()
	return p
}

func Spelling(word voca.Word) test.Problem {
	p := new(problem)
	p.word = word
	p.question = p.word.Mean()
	p.answer = p.word.Spell()
	return p
}

func (p *problem) Question() voca.Unit {
	return p.question
}

func (p *problem) IsCorrect(userAnswer string) bool {
	return p.answer.IsRight(userAnswer)
}
