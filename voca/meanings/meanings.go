package meanings

import (
	"wordReminder-core/voca"
)

func New(mean...voca.Meaning) voca.Meanings {
	return voca.Meanings{Data: mean}
}