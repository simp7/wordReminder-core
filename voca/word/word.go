package word

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/unit"
	"github.com/simp7/wordReminder-core/voca/unit/meanings"
	"github.com/simp7/wordReminder-core/voca/unit/spelling"
)

type word struct {
	Spelling unit.Spelling
	Meaning  unit.Meanings
}

func New(spell string, mean ...unit.Meaning) voca.Word {
	return word{Spelling: spelling.New(spell), Meaning: meanings.New(mean...)}
}

func (w word) Mean() voca.Unit {
	return w.Meaning
}

func (w word) Spell() voca.Unit {
	return w.Spelling
}
