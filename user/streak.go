package user

type Streak uint

func (s *Streak) Increment() {
	*s++
}

func (s *Streak) Decrement() {
	if *s > 0 {
		*s--
	}
}
