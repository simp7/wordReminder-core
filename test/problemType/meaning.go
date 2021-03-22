package problemType

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca"
)

type meaning struct {
}

func Meaning() test.Type {
	m := new(meaning)
	return m
}

func (m meaning) GetQuestion(word voca.Word) voca.Unit {
	return word.Spell
}

func (m meaning) GetAnswer(word voca.Word) voca.Unit {
	return word.Mean
}
