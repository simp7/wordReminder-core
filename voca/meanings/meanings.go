package meanings

import (
	"github.com/simp7/wordReminder-core/voca"
)

func New(mean ...voca.Meaning) voca.Meanings {
	return voca.Meanings{Data: mean}
}
