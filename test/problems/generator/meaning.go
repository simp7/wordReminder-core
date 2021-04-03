package generator

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/test/problem"
	"github.com/simp7/wordReminder-core/test/problems"
	"github.com/simp7/wordReminder-core/voca"
)

type meaningGen struct {
	words voca.WordSet
}

func Meaning(set voca.WordSet) *meaningGen {
	m := new(meaningGen)
	m.words = set
	return m
}

func (m *meaningGen) doEach(w voca.Word) test.Problem {
	return problem.Meaning(w)
}

func (m *meaningGen) Do() test.Problems {

	if m.words.Len() == 0 {
		return nil
	}

	result := problems.Array()

	for i := 0; i < m.words.Len(); i += 1 {
		result.Add(m.doEach(m.words.Get(i)))
	}

	return result

}
