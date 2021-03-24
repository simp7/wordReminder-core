package wordSet

import "github.com/simp7/wordReminder-core/voca"

type set struct {
	ID    string
	Words []voca.Word
}

func New(id string, words ...voca.Word) voca.WordSet {
	s := new(set)
	s.ID = id
	s.Words = words
	return s
}

func (s *set) Len() int {
	return len(s.Words)
}

func (s *set) Get(idx int) voca.Word {
	return s.Words[idx]
}

func (s *set) Less(i, j int) bool {
	w1 := s.Words[i]
	w2 := s.Words[j]
	return w1.Spell().Format() < w2.Spell().Format()
}

func (s *set) Swap(i, j int) {
	s.Words[i], s.Words[j] = s.Words[j], s.Words[i]
}
