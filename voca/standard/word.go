package standard

import (
	"github.com/simp7/wordReminder-core/voca"
)

type word struct {
	Spelling voca.Spelling
	Meaning  voca.Meanings
}

func Word(spell string, mean ...voca.Meaning) voca.Word {
	return word{Spelling: Spelling(spell), Meaning: Meanings(mean...)}
}

func (w word) Mean() voca.Meanings {
	return w.Meaning
}

func (w word) Spell() voca.Spelling {
	return w.Spelling
}
