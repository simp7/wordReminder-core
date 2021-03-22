package spelling

import "github.com/simp7/wordReminder-core/voca"

type spelling string

func New(spell string) voca.Spelling {
	return voca.Spelling{Data: spell}
}
