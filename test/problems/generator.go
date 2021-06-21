package problems

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca"
)

type generator struct {
	words   voca.WordSet
	genFunc func(word voca.Word) test.Problem
}

func newGenerator(words voca.WordSet) *generator {

	g := new(generator)
	g.words = words

	return g

}

func meaning(words voca.WordSet) *generator {

	g := newGenerator(words)
	g.genFunc = test.Meaning

	return g

}

func spelling(words voca.WordSet) *generator {

	g := newGenerator(words)
	g.genFunc = test.Spelling

	return g

}

func (g *generator) Generate() []test.Problem {

	problems := make([]test.Problem, g.words.Len())

	for i := 0; i < g.words.Len(); i++ {
		problems[i] = g.genFunc(g.words.Get(i))
	}

	return problems

}
