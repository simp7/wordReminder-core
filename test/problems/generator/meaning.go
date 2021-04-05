package generator

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/test/problem"
	"github.com/simp7/wordReminder-core/test/problems"
	"github.com/simp7/wordReminder-core/voca"
)

type meaningGen struct {
}

func Meaning() *meaningGen {
	return new(meaningGen)
}

func (m *meaningGen) doEach(w voca.Word) test.Problem {
	return problem.Meaning(w)
}

func (m *meaningGen) Generate(words voca.WordSet) test.Problems {

	if words.Len() == 0 {
		return nil
	}

	result := problems.Array()

	for i := 0; i < words.Len(); i += 1 {
		result.Add(m.doEach(words.Get(i)))
	}

	return result

}
