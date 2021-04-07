package streak

import (
	"errors"
	"github.com/simp7/wordReminder-core/user"
)

type streak int

var (
	NegativeStreakErr = errors.New("streak should not be negative")
)

func (s *streak) Increment() {
	*s += 1
}

func New(preset int) (user.Streak, error) {
	s := new(streak)
	if preset < 0 {
		return nil, NegativeStreakErr
	}
	*s = streak(preset)
	return s, nil
}
