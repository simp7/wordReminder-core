package wordSet

import "wordReminder-core/voca"

type set struct {
	ID string
	Words []voca.Word
}

func New(id string, words... voca.Word) voca.WordSet {
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
	return s.Words[i].Spelling < s.Words[j].Spelling
}

func (s *set) Swap(i, j int) {
	s.Words[i], s.Words[j] = s.Words[j], s.Words[i]
}