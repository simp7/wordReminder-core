package test

import (
	"github.com/simp7/wordReminder-core/voca"
)

type Problem struct {
	word     voca.Word
	Question voca.Unit
	answer   voca.Unit
}

func Meaning(word voca.Word) Problem {
	p := Problem{word, word.Spell(), word.Mean()}
	return p
}

func Spelling(word voca.Word) Problem {
	p := Problem{word, word.Mean(), word.Spell()}
	return p
}

func (p Problem) IsCorrect(userAnswer string) bool {
	return p.answer.IsRight(userAnswer)
}
