package test

import "github.com/simp7/wordReminder-core/voca"

type Selector interface {
	Select(amount int) voca.WordSet
}
