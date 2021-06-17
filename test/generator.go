package test

import (
	"github.com/simp7/wordReminder-core/voca"
)

type Generator interface {
	Generate(set voca.WordSet) Problems
}
