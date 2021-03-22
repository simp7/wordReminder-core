package input

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca"
)

type choice struct {
	choiceAmount int
}

func Choice(word voca.Word, choices int) test.Input {
	c := new(choice)
	return c
}

func (c *choice) Show() {

}
