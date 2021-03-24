package streak

import "github.com/simp7/wordReminder-core/user"

type streak struct {
}

func New() user.Streak {
	s := new(streak)
	return s
}
