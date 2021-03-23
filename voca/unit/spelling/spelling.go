package spelling

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/unit"
)

type spelling struct {
	Data string
}

func New(spell string) unit.Spelling {
	return spelling{Data: spell}
}

func (s spelling) IsEqual(u voca.Unit) bool {
	return s == u
}

func (s spelling) Format() string {
	return s.Data
}
