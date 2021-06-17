package standard

import "github.com/simp7/wordReminder-core/voca"

type set struct {
	ID    string
	Words []voca.Word
}

func Words(id string, words ...voca.Word) voca.WordSet {
	s := new(set)
	s.ID = id
	s.Words = words
	return s
}

func (s *set) Len() int {
	return len(s.Words)
}

func (s *set) Get(idx int) voca.Word {
	if s.Len() <= idx {
		return nil
	}
	return s.Words[idx]
}

func (s *set) Less(i, j int) bool {
	w1 := s.Get(i)
	w2 := s.Get(j)
	if w1 == nil || w2 == nil {
		return false
	}
	return w1.Spell().String() < w2.Spell().String()
}

func (s *set) Swap(i, j int) {
	if s.Get(i) == nil || s.Get(j) == nil {
		return
	}
	s.Words[i], s.Words[j] = s.Get(j), s.Get(i)
}

func (s *set) Trim(to int) {
	s.Words = s.Words[:to]
}
