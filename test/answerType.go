package test

import "github.com/simp7/wordReminder-core/voca"

type AnswerType interface {
	GetQuestion(voca.Word) string
	GetAnswer(voca.Word) string
}
