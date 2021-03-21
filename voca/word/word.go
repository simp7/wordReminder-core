package word

import (
	"wordReminder-core/voca"
	"wordReminder-core/voca/meanings"
)

func New(spelling string, mean ...voca.Meaning) voca.Word {
	return voca.Word{Spelling: spelling, Mean: meanings.New(mean...)}
}