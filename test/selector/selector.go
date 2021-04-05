package selector

import (
	"github.com/simp7/wordReminder-core/voca"
	"math/rand"
	"time"
)

type selector struct {
	seed  int64
	words voca.WordSet
}

func New(words voca.WordSet) *selector {

	g := new(selector)

	g.seed = time.Now().Unix()
	g.words = words

	return g

}

func (s *selector) Select(amount int) voca.WordSet {

	idx := amount
	if amount > s.words.Len() {
		idx = s.words.Len()
	}

	rand.Seed(s.seed)
	rand.Shuffle(s.words.Len(), s.words.Swap)
	s.words.Trim(idx)

	return s.words

}
