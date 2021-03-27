package input

import (
	"github.com/simp7/wordReminder-core/test"
)

type choice struct {
	choiceAmount int
}

func Choice(choices int) test.Input {
	c := new(choice)
	c.choiceAmount = choices
	return c
}
