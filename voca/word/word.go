package word

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/meanings"
	"github.com/simp7/wordReminder-core/voca/spelling"
)

type word struct {
	Spelling voca.Spelling
	Meaning  voca.Meanings
}

func New(spell string, mean ...voca.Meaning) voca.Word {
	return word{Spelling: spelling.New(spell), Meaning: meanings.New(mean...)}
}

func (w word) Mean() voca.Meanings {
	return w.Meaning
}

func (w word) Spell() voca.Spelling {
	return w.Spelling
}
