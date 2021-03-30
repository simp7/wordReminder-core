package spelling

import (
	"github.com/simp7/wordReminder-core/voca"
)

type spelling struct {
	Data string
}

func New(spell string) voca.Spelling {
	return spelling{Data: spell}
}

func (s spelling) IsRight(input string) bool {
	return s.Data == input
}

func (s spelling) String() string {
	return s.Data
}
