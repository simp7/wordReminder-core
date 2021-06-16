package creator

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/test/problems"
	"github.com/simp7/wordReminder-core/test/problems/generator"
	"github.com/simp7/wordReminder-core/voca"
)

type creator struct {
	words     voca.WordSet
	generator problems.Generator
	selector  test.Selector
}

func Meaning(words voca.WordSet) *creator {
	c := new(creator)
	c.generator = generator.Meaning()
	return c
}

func (c *creator) Create(amount int) test.Problems {
	return c.generator.Generate(c.selector.Select(amount))
}
