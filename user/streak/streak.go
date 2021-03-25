package streak

import (
	"errors"
	"github.com/simp7/wordReminder-core/user"
	"strconv"
)

type streak int

var (
	NegativeStreakErr = errors.New("streak should not be negative")
)

func (s *streak) Increment() {
	*s += 1
}

func (s *streak) String() string {
	data := int(*s)
	return strconv.Itoa(data)
}

func New(preset int) (user.Streak, error) {
	s := new(streak)
	if preset < 0 {
		return nil, NegativeStreakErr
	}
	*s = streak(preset)
	return s, nil
}
