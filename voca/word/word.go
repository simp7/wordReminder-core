package word

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/meanings"
)

func New(spelling string, mean ...voca.Meaning) voca.Word {
	return voca.Word{Spelling: spelling, Mean: meanings.New(mean...)}
}
