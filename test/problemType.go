package test

import "github.com/simp7/wordReminder-core/voca"

type Type interface {
	GetQuestion(voca.Word) voca.Unit
	GetAnswer(voca.Word) voca.Unit
}
