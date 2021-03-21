package test

import "wordReminder-core/voca"

type AnswerType interface {
	GetQuestion(voca.Word) string
	GetAnswer(voca.Word) string
}