package word

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/meanings"
	"github.com/simp7/wordReminder-core/voca/spelling"
)

func New(spell string, mean ...voca.Meaning) voca.Word {
	return voca.Word{Spell: spelling.New(spell), Mean: meanings.New(mean...)}
}
