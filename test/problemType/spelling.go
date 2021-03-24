package problemType

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca"
)

type spelling struct {
}

func Spelling() test.Type {
	s := new(spelling)
	return s
}

func (s spelling) GetQuestion(word voca.Word) voca.Unit {
	return word.Mean()
}

func (s spelling) GetAnswer(word voca.Word) voca.Unit {
	return word.Spell()
}
