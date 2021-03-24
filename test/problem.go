package test

import "github.com/simp7/wordReminder-core/voca"

type Problem interface {
	Question() voca.Unit
	IsCorrect(userAnswer string) bool
}
