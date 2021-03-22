package problem

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca"
)

type problem struct {
	word       voca.Word
	input      test.Input
	answerType test.AnswerType
}

func New(word voca.Word, input test.Input, answerType test.AnswerType) test.Problem {
	p := new(problem)
	p.word = word
	p.input = input
	p.answerType = answerType
	return p
}

func (p *problem) Question() string {
	return p.answerType.GetQuestion(p.word)
}

func (p *problem) IsCorrect(userAnswer string) bool {
	return p.answerType.GetAnswer(p.word) == userAnswer
}
