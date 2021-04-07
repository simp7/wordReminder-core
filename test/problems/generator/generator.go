package generator

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/test/problem"
	"github.com/simp7/wordReminder-core/test/problems"
	"github.com/simp7/wordReminder-core/voca"
)

type generator struct {
	makeProblem func(voca.Word) test.Problem
}

func Meaning() *generator {
	gen := new(generator)
	gen.makeProblem = problem.Meaning
	return gen
}

func Spelling() *generator {
	gen := new(generator)
	gen.makeProblem = problem.Spelling
	return gen
}

func (g *generator) doEach(w voca.Word) test.Problem {
	return g.makeProblem(w)
}

func (g *generator) Generate(words voca.WordSet) test.Problems {

	if words.Len() == 0 {
		return nil
	}

	result := problems.Array()

	for i := 0; i < words.Len(); i += 1 {
		result.Add(g.doEach(words.Get(i)))
	}

	return result

}
